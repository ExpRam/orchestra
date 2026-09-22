package validation

import (
	"fmt"
	"strings"
)

const (
	sourceSeparator  = ": "
	problemSeparator = "; "
)

type Result struct {
	Problems []error
}

func Failed(problems ...error) Result {
	return Result{Problems: problems}
}

func (r Result) Valid() bool {
	return len(r.Problems) == 0
}

func (r Result) Err(source fmt.Stringer) error {
	if r.Valid() {
		return nil
	}

	return Error{Source: source, Problems: r.Problems}
}

type Error struct {
	Source   fmt.Stringer
	Problems []error
}

func (e Error) Error() string {
	messages := make([]string, 0, len(e.Problems))
	for _, problem := range e.Problems {
		messages = append(messages, problem.Error())
	}

	joined := strings.Join(messages, problemSeparator)
	if e.Source == nil {
		return joined
	}

	return e.Source.String() + sourceSeparator + joined
}

func (e Error) Unwrap() []error {
	return e.Problems
}
