package source

import (
	"os"
	"strings"

	"github.com/expram/orchestra/internal/config"
)

type Env struct {
	prefix string
}

var _ config.Source = (*Env)(nil)

func NewEnv(prefix string) *Env {
	return &Env{prefix: prefix}
}

func (s *Env) Load() (config.Values, error) {
	env := make(config.Values)

	for _, item := range os.Environ() {
		key, value, ok := strings.Cut(item, "=")
		if !ok {
			continue
		}

		name, found := strings.CutPrefix(key, s.prefix)
		if found && name != "" {
			env[name] = value
		}
	}

	return env, nil
}
