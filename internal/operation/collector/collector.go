package collector

import (
	"path"
	"slices"
	"strings"

	"github.com/expram/orchestra/internal/workspace"
)

var manifestExtensions = []string{".yml", ".yaml"}

type Collector struct{}

func NewCollector() Collector {
	return Collector{}
}

func (c Collector) Collect(ws workspace.Workspace, directory string) ([]string, error) {
	var files []string

	collect := func(file string, size int64) error {
		if size == 0 || !isManifestFile(file) {
			return nil
		}

		files = append(files, file)

		return nil
	}

	if err := ws.Walk(directory, collect); err != nil {
		return nil, err
	}

	return files, nil
}

func isManifestFile(file string) bool {
	ext := strings.ToLower(path.Ext(file))

	return slices.Contains(manifestExtensions, ext)
}
