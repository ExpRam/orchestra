package config

import (
	"maps"

	"github.com/joho/godotenv"
)

type EnvFile struct {
	path string
}

func NewEnvFile(path string) *EnvFile {
	return &EnvFile{path: path}
}

func (s *EnvFile) Load() (Values, error) {
	env, err := godotenv.Read(s.path)
	if err != nil {
		return nil, err
	}

	values := make(Values, len(env))

	maps.Copy(values, env)

	return values, nil
}
