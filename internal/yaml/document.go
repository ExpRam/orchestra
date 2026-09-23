package yaml

import (
	"bytes"
	"errors"
	"io"

	yamlv3 "go.yaml.in/yaml/v3"
)

const nullTag = "!!null"

type Document struct {
	Index int
	Line  int

	node *yamlv3.Node
}

func (d Document) Location() string {
	return location(d.Line, d.Index)
}

func (d Document) Decode(value any) error {
	if err := d.node.Decode(value); err != nil {
		return translate(err, d.Index)
	}

	return nil
}

func Split(data []byte) ([]Document, error) {
	var documents []Document

	decoder := yamlv3.NewDecoder(bytes.NewReader(data))

	for index := 1; ; index++ {
		var node yamlv3.Node

		if err := decoder.Decode(&node); err != nil {
			if errors.Is(err, io.EOF) {
				return documents, nil
			}

			return nil, translate(err, index)
		}

		content := content(&node)
		if content == nil {
			continue
		}

		documents = append(documents, Document{Index: index, Line: content.Line, node: content})
	}
}

func content(node *yamlv3.Node) *yamlv3.Node {
	if node.Kind != yamlv3.DocumentNode || len(node.Content) != 1 {
		return nil
	}

	content := node.Content[0]
	if content.Tag == nullTag {
		return nil
	}

	return content
}
