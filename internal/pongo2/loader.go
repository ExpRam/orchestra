package pongo2

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/expram/orchestra/internal/workspace"
)

const librarySeparator = ", "

var errOutsideLibrary = errors.New("path must be relative and stay inside a library")

type loader struct {
	ws        workspace.Workspace
	libraries []workspace.Directory
	resolved  map[string]string
	failed    map[string]error
}

func newLoader(ws workspace.Workspace, libraries []workspace.Directory) *loader {
	return &loader{
		ws:        ws,
		libraries: libraries,
		resolved:  make(map[string]string),
		failed:    make(map[string]error),
	}
}

func (l *loader) Abs(_, name string) string {
	return name
}

func (l *loader) Get(name string) (io.Reader, error) {
	data, err := l.load(name)
	if err != nil {
		l.failed[name] = err

		return nil, err
	}

	return bytes.NewReader(data), nil
}

func (l *loader) load(name string) ([]byte, error) {
	if !filepath.IsLocal(name) {
		return nil, errOutsideLibrary
	}

	for _, library := range l.libraries {
		file := filepath.Join(library, name)

		data, err := l.ws.ReadFile(file)
		if errors.Is(err, fs.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, err
		}

		l.resolved[name] = file

		return data, nil
	}

	return nil, fmt.Errorf("not found in %s", strings.Join(l.libraries, librarySeparator))
}

func (l *loader) resolve(name string) string {
	if file, ok := l.resolved[name]; ok {
		return file
	}

	return name
}
