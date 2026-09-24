package mapstructure

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	mapstructurev2 "github.com/go-viper/mapstructure/v2"

	"github.com/expram/orchestra/internal/errscope"
)

const (
	invalidKeysPrefix = "has invalid keys: "
	keySeparator      = ", "
	pathSeparator     = "."
)

func translate(err error) error {
	var problems errscope.Problems
	collect(err, &problems)

	return problems.Err()
}

func collect(err error, problems *errscope.Problems) {
	var decodeErr *mapstructurev2.DecodeError
	if errors.As(err, &decodeErr) && err == error(decodeErr) {
		describe(decodeErr, problems)

		return
	}

	switch e := err.(type) {
	case interface{ Unwrap() []error }:
		for _, child := range e.Unwrap() {
			collect(child, problems)
		}
	case interface{ Unwrap() error }:
		collect(e.Unwrap(), problems)
	default:
		problems.Add(err)
	}
}

func describe(decodeErr *mapstructurev2.DecodeError, problems *errscope.Problems) {
	path := decodeErr.Name()
	cause := decodeErr.Unwrap()

	if keys, ok := strings.CutPrefix(cause.Error(), invalidKeysPrefix); ok {
		for key := range strings.SplitSeq(keys, keySeparator) {
			problems.Add(fmt.Errorf("unknown field %q", join(path, key)))
		}

		return
	}

	if unconvertible, ok := errors.AsType[*mapstructurev2.UnconvertibleTypeError](cause); ok {
		problems.Add(fmt.Errorf("%q must be %s", path, article(unconvertible.Expected.Kind())))

		return
	}

	problems.Add(fmt.Errorf("%q: %w", path, cause))
}

func join(path, key string) string {
	if path == "" {
		return key
	}

	return path + pathSeparator + key
}

func article(kind reflect.Kind) string {
	switch kind {
	case reflect.String:
		return "a string"
	case reflect.Bool:
		return "a boolean"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "an integer"
	case reflect.Float32, reflect.Float64:
		return "a number"
	case reflect.Slice, reflect.Array:
		return "a list"
	case reflect.Map, reflect.Struct:
		return "an object"
	default:
		return "a " + kind.String()
	}
}
