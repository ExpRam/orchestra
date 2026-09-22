package reader

import (
	"fmt"

	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/workspace"
)

type Reader struct {
	parser UnresolvedManifestParser
}

func New(parser UnresolvedManifestParser) Reader {
	return Reader{parser: parser}
}

func (r Reader) Read(ws workspace.Workspace, files []string) ([]manifest.UnresolvedManifest, error) {
	var manifests []manifest.UnresolvedManifest

	for _, file := range files {
		data, err := ws.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("read %q: %w", file, err)
		}

		parsed, err := r.parser.Parse(file, data)
		if err != nil {
			return nil, err
		}

		manifests = append(manifests, parsed...)
	}

	return manifests, nil
}
