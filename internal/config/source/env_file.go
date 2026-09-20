package source

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/expram/orchestra/internal/config"
	"github.com/joho/godotenv"
)

type EnvFile struct {
	path string
}

var _ config.Source = (*EnvFile)(nil)

func NewEnvFile(path string) *EnvFile {
	return &EnvFile{path: path}
}

func (s *EnvFile) Load() (config.Values, error) {
	env, err := godotenv.Read(s.path)

	if errors.Is(err, fs.ErrNotExist) {
		return config.Values{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}

	return env, nil
}
