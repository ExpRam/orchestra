package contract

import "github.com/expram/orchestra/internal/validation"

type Parser[T any] interface {
	Parse(name string, data []byte) ([]T, error)
}

type Validator[T any] interface {
	Validate(value T) validation.Result
}
