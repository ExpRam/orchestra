package yaml

import (
	"fmt"
	"strconv"

	"github.com/expram/orchestra/internal/manifest"
)

type source struct {
	document int
	line     int
}

var _ manifest.Source = source{}

func newSource(document, line int) source {
	return source{document: document, line: line}
}

func (s source) String() string {
	if s.document > 1 {
		return fmt.Sprintf("%d (document %d)", s.line, s.document)
	}

	return strconv.Itoa(s.line)
}
