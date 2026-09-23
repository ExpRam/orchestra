package yaml

import (
	"bytes"
	"errors"
	"io"

	yamlv3 "go.yaml.in/yaml/v3"

	"github.com/expram/orchestra/internal/errscope"
)

const nullTag = "!!null"

var (
	errNotMapping    = errors.New("document must be a mapping")
	errNonStringKeys = errors.New("mapping keys must be strings")
)

type Document struct {
	index int
	line  int
	node  *yamlv3.Node
}

func (d Document) Location() string {
	return location(d.line, d.index)
}

func (d Document) Object() (map[string]any, error) {
	if d.node.Kind != yamlv3.MappingNode {
		return nil, errscope.In(d.Location(), errNotMapping)
	}

	var tree any
	if err := d.node.Decode(&tree); err != nil {
		return nil, translate(err, d.index, d.Location())
	}

	object, ok := tree.(map[string]any)
	if !ok {
		return nil, errscope.In(d.Location(), errNonStringKeys)
	}

	return object, nil
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

			return nil, translate(err, index, "")
		}

		content := content(&node)
		if content == nil {
			continue
		}

		documents = append(documents, Document{index: index, line: content.Line, node: content})
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
