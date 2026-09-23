package reader

import (
	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/workspace"
)

type Reader struct {
	parser UnresolvedManifestParser
}

func NewReader(parser UnresolvedManifestParser) Reader {
	return Reader{parser: parser}
}

func (r Reader) Read(ws workspace.Workspace, files []string) ([]manifest.UnresolvedManifest, error) {
	var (
		manifests []manifest.UnresolvedManifest
		problems  errscope.Problems
	)

	for _, file := range files {
		data, err := ws.ReadFile(file)
		if err != nil {
			problems.Add(err)

			continue
		}

		parsed, err := r.parser.Parse(data)
		if err != nil {
			problems.Add(errscope.In(file, err))

			continue
		}

		manifests = append(manifests, parsed...)
	}

	if err := problems.Err(); err != nil {
		return nil, err
	}

	return manifests, nil
}
