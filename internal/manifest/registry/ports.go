package registry

type Migration interface {
	Validate() error
	Convert() (any, error)
}

type SpecDecoder interface {
	Decode(tree map[string]any, target any) error
}
