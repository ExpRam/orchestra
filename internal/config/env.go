package config

import (
	"os"
	"strings"
)

type Env struct{}

func NewEnv() *Env {
	return &Env{}
}

func (s *Env) Load() (Values, error) {
	env := make(Values)

	for _, item := range os.Environ() {
		key, value, ok := strings.Cut(item, "=")
		if !ok {
			continue
		}

		env[key] = value
	}

	return env, nil
}
