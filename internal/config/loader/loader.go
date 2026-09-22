package loader

import (
	"fmt"

	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/v2"

	"github.com/expram/orchestra/internal/config"
)

const Tag = "config"

type Loader struct {
	sources []config.Source
}

func New(sources ...config.Source) *Loader {
	return &Loader{sources: sources}
}

func (l *Loader) Load() (config.Config, error) {
	values := koanf.New(config.Delim)

	for _, source := range l.sources {
		layer, err := source.Read()
		if err != nil {
			return config.Config{}, fmt.Errorf("load config source %s: %w", source.Name(), err)
		}

		if err := values.Load(confmap.Provider(layer, config.Delim), nil); err != nil {
			return config.Config{}, fmt.Errorf("merge config source %s: %w", source.Name(), err)
		}
	}

	cfg := config.Defaults()
	if err := values.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{Tag: Tag}); err != nil {
		return config.Config{}, fmt.Errorf("decode configuration: %w", err)
	}

	return cfg, nil
}
