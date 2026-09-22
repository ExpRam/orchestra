package manifest

import "strings"

const (
	sourceSeparator  = ": "
	problemSeparator = "; "
)

type Error struct {
	Source   Source
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
