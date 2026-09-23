package env

import (
	"errors"
	"io/fs"
	"os"

	"github.com/knadh/koanf/parsers/dotenv"

	"github.com/expram/orchestra/internal/config"
	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/errscope/fserror"
)

type File struct {
	env
	file string
}

var _ config.Source = File{}

func NewFile(file, prefix string) File {
	return File{env{prefix: prefix}, file}
}

func (s File) Name() string {
	return "file:" + s.file
}

func (s File) Read() (map[string]any, error) {
	raw, err := os.ReadFile(s.file)

	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}

	if err != nil {
		return nil, errscope.In("read", fserror.Cause(err))
	}

	return dotenv.ParserEnv(s.prefix, config.Delim, s.path).Unmarshal(raw)
}
