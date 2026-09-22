package yaml

import (
	"fmt"

	"github.com/expram/orchestra/internal/manifest"
)

type source struct {
	name     string
	document int
	line     int
}

var _ manifest.Source = source{}

func newSource(name string, document, line int) source {
	return source{name: name, document: document, line: line}
}

func stream(name string) source {
	return source{name: name}
}

func (s source) String() string {
	switch {
	case s.line == 0:
		return s.name
	case s.document > 1:
		return fmt.Sprintf("%s:%d (document %d)", s.name, s.line, s.document)
	default:
		return fmt.Sprintf("%s:%d", s.name, s.line)
	}
}
