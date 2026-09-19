package config

import (
	"os"
	"strings"
)

type Env struct {
	prefix string
}

func NewEnv(prefix string) *Env {
	return &Env{prefix: prefix}
}

func (s *Env) Load() (Values, error) {
	env := make(Values)

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
