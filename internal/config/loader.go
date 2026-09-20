package config

import "fmt"

type Source interface {
	Load() (Values, error)
}

type Loader struct {
	sources []Source
}

func NewLoader(sources ...Source) *Loader {
	return &Loader{sources: sources}
}

func (l *Loader) Load() (Values, error) {
	result := make(Values)

	for _, source := range l.sources {
		values, err := source.Load()
		if err != nil {
			return nil, fmt.Errorf("load config source %T: %w", source, err)
		}

		merge(result, values)
	}

	return result, nil
}

func merge(dst, src Values) {
	for key, value := range src {
		dst[normalize(key)] = value
	}
}
