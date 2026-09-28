package processor

type KindHandler interface {
	Handle(input HandleInput) error
}

type SpecHandler[S any] interface {
	Handle(input HandleInput, spec S) error
}
