package cli

import (
	"log/slog"
	"strings"

	"github.com/expram/orchestra/internal/errscope"
)

const (
	errorKey         = "error"
	scopeSeparator   = ": "
	problemSeparator = "; "
	groupOpen        = "["
	groupClose       = "]"
)

type unwrappable interface {
	Unwrap() []error
}

func ErrorAttr(err error) slog.Attr {
	return slog.String(errorKey, renderError(err, false))
}

func renderError(err error, listed bool) string {
	switch e := err.(type) {
	case errscope.Error:
		return e.Scope + scopeSeparator + renderError(e.Cause, listed)
	case unwrappable:
		return renderAggregate(e.Unwrap(), listed)
	default:
		return err.Error()
	}
}

func renderAggregate(problems []error, listed bool) string {
	if len(problems) == 1 {
		return renderError(problems[0], listed)
	}

	messages := make([]string, 0, len(problems))
	for _, problem := range problems {
		messages = append(messages, renderError(problem, true))
	}

	joined := strings.Join(messages, problemSeparator)
	if listed {
		return groupOpen + joined + groupClose
	}

	return joined
}
