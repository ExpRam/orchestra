package contract

type Parser[T any] interface {
	Parse(data []byte) ([]T, error)
}

type Validator[In, Out any] interface {
	Validate(value In) (Out, error)
}
