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

const libraryPrefix = "yaml: "

var linePattern = regexp.MustCompile(`^line (\d+): (.*)$`)

func translate(err error, index int) error {
	var typeError *yamlv3.TypeError
	if !errors.As(err, &typeError) {
		return problem(err.Error(), index)
	}

	var problems errscope.Problems
	for _, message := range typeError.Errors {
		problems.Add(problem(message, index))
	}

	return problems.Err()
}

func problem(message string, index int) error {
	message = strings.TrimPrefix(message, libraryPrefix)

	match := linePattern.FindStringSubmatch(message)
	if match == nil {
		return errors.New(message)
	}

	line, err := strconv.Atoi(match[1])
	if err != nil {
		return errors.New(message)
	}

	return errscope.In(location(line, index), errors.New(match[2]))
}

func location(line, index int) string {
	if index > 1 {
		return fmt.Sprintf("line %d (document %d)", line, index)
	}

	return fmt.Sprintf("line %d", line)
}
