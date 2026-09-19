package config

import (
	"fmt"
	"strings"
)

type Values = map[string]string

type Config struct {
	Debug bool
}

func NewConfig() *Config {
	return &Config{Debug: false}
}

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
			return nil, fmt.Errorf("load config source: %w", err)
		}

		merge(result, values)
	}

	return result, nil
}

func merge(dst, src Values) {
	for k, v := range src {
		dst[strings.ToUpper(k)] = v
	}
}
