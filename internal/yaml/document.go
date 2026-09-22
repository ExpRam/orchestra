package yaml

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	yamlv3 "go.yaml.in/yaml/v3"
)

const nullTag = "!!null"

type Document struct {
	Index int
	Line  int

	node *yamlv3.Node
}

func (d Document) Mapping() bool {
	return d.node.Kind == yamlv3.MappingNode
}

func (d Document) Decode(value any) error {
	return d.node.Decode(value)
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

			return nil, fmt.Errorf("document %d: %w", index, err)
		}

		content := content(&node)
		if content == nil {
			continue
		}

		documents = append(documents, Document{Index: index, Line: content.Line, node: content})
	}
}

func Errors(err error) []error {
	var typeError *yamlv3.TypeError
	if !errors.As(err, &typeError) {
		return []error{err}
	}

	problems := make([]error, 0, len(typeError.Errors))
	for _, message := range typeError.Errors {
		problems = append(problems, errors.New(message))
	}

	return problems
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
