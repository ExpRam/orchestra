package config

import (
	"errors"
	"fmt"
	"io/fs"
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

	if errors.Is(err, fs.ErrNotExist) {
		return Values{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("cannot read: %w", err)
	}

	values := make(Values, len(env))

	maps.Copy(values, env)

	return values, nil
}
