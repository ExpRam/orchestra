package pongo2

import (
	"errors"
	"fmt"

	pongo2v6 "github.com/flosch/pongo2/v6"

	"github.com/expram/orchestra/internal/errscope"
)

const stringTemplate = "<string>"

func translate(err error, l *loader) error {
	var templateErr *pongo2v6.Error
	if !errors.As(err, &templateErr) || templateErr.OrigError == nil {
		return err
	}

	name := templateErr.Filename
	imported := name != "" && name != stringTemplate

	if failure, ok := l.failed[name]; imported && ok {
		return at(templateErr, errscope.In(importScope(name), failure))
	}

	cause := at(templateErr, translate(templateErr.OrigError, l))
	if imported {
		return errscope.In(importScope(l.resolve(name)), cause)
	}

	return cause
}

func at(templateErr *pongo2v6.Error, cause error) error {
	if templateErr.Line <= 0 {
		return cause
	}

	return errscope.In(fmt.Sprintf("line %d, column %d", templateErr.Line, templateErr.Column), cause)
}

func importScope(file string) string {
	return fmt.Sprintf("import %q", file)
}
