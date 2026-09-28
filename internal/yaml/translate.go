package yaml

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	yamlv3 "go.yaml.in/yaml/v3"

	"github.com/expram/orchestra/internal/errscope"
)

const (
	libraryPrefix       = "yaml: "
	invalidMapKeyPrefix = "invalid map key:"
)

var linePattern = regexp.MustCompile(`^line (\d+): (.*)$`)

var errInvalidMapKey = errors.New("map keys must be scalars")

func translate(err error, index int, fallback string) error {
	var typeError *yamlv3.TypeError
	if !errors.As(err, &typeError) {
		return problem(err.Error(), index, fallback)
	}

	var problems errscope.Problems
	for _, message := range typeError.Errors {
		problems.Add(problem(message, index, fallback))
	}

	return problems.Err()
}

func problem(message string, index int, fallback string) error {
	message = strings.TrimPrefix(message, libraryPrefix)

	match := linePattern.FindStringSubmatch(message)
	if match == nil {
		return at(fallback, describe(message))
	}

	line, err := strconv.Atoi(match[1])
	if err != nil {
		return at(fallback, describe(message))
	}

	return at(location(line, index), describe(match[2]))
}

func describe(message string) error {
	if strings.HasPrefix(message, invalidMapKeyPrefix) {
		return errInvalidMapKey
	}

	return errors.New(message)
}

func at(scope string, cause error) error {
	if scope == "" {
		return cause
	}

	return errscope.In(scope, cause)
}

func location(line, index int) string {
	if index > 1 {
		return fmt.Sprintf("line %d (document %d)", line, index)
	}

	return fmt.Sprintf("line %d", line)
}
