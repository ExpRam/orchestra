package loader

import (
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/v2"

	"github.com/expram/orchestra/internal/config"
	"github.com/expram/orchestra/internal/errscope"
)

const Tag = "config"

type Loader struct {
	sources []config.Source
}

func NewLoader(sources ...config.Source) Loader {
	return Loader{sources: sources}
}

func (l Loader) Load() (config.Config, error) {
	values := koanf.New(config.Delim)

	for _, source := range l.sources {
		layer, err := source.Read()
		if err != nil {
			return config.Config{}, errscope.In("load config source "+source.Name(), err)
		}

		if err := values.Load(confmap.Provider(layer, config.Delim), nil); err != nil {
			return config.Config{}, errscope.In("merge config source "+source.Name(), err)
		}
	}

	cfg := config.Defaults()
	if err := values.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{Tag: Tag}); err != nil {
		return config.Config{}, errscope.In("decode configuration", err)
	}

	return cfg, nil
}
