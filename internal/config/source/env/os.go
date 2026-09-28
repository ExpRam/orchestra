package env

import (
	kenv "github.com/knadh/koanf/providers/env/v2"

	"github.com/expram/orchestra/internal/config"
	"github.com/expram/orchestra/internal/config/loader"
)

type OS struct {
	env
}

var _ loader.Source = OS{}

func NewOS(prefix string) OS {
	return OS{env{prefix: prefix}}
}

func (s OS) Name() string {
	return "env"
}

func (s OS) Read() (map[string]any, error) {
	return kenv.Provider(config.Delim, kenv.Opt{
		Prefix: s.prefix,
		TransformFunc: func(name, value string) (string, any) {
			return s.path(name), value
		},
	}).Read()
}
