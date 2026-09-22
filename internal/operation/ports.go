package operation

import (
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/workspace"
)

type WorkspacePreparer interface {
	Prepare(builtin, user string) (workspace.Workspace, error)
}

type ManifestFileCollector interface {
	Collect(ws workspace.Workspace, directory string) ([]string, error)
}

type UnresolvedManifestReader interface {
	Read(ws workspace.Workspace, files []string) ([]manifest.UnresolvedManifest, error)
}
