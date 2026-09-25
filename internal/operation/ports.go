package operation

import (
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/workspace"
)

type WorkspacePreparer interface {
	Prepare(builtin, user string, keep bool) (workspace.Workspace, error)
}

type ManifestFileCollector interface {
	Collect(ws workspace.Workspace, directory string) ([]string, error)
}

type ManifestRenderer interface {
	Render(ws workspace.Workspace, files []string, libraries []workspace.Directory) ([]string, error)
}

type UnresolvedManifestReader interface {
	Read(ws workspace.Workspace, files []string) ([]manifest.UnresolvedManifest, error)
}

type ManifestResolver interface {
	Resolve(unresolved []manifest.UnresolvedManifest) (manifest.Catalog, error)
}

type ProcessInput struct {
	Manifest  manifest.Manifest
	Operation InfrastructureOperation
	Catalog   manifest.Catalog
	Workspace workspace.Workspace
}

type ManifestProcessor interface {
	Process(input ProcessInput) error
}

type ManifestProcessors interface {
	Process(op InfrastructureOperation, catalog manifest.Catalog, ws workspace.Workspace) error
}
