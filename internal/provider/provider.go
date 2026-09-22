package provider

type Provider[T any] interface {
	Get() T
}

type Func[T any] func() T

func (f Func[T]) Get() T {
	return f()
}

func New[T any](get func() T) Provider[T] {
	return Func[T](get)
}
