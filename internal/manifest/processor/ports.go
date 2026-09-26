package processor

import (
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/operation"
	"github.com/expram/orchestra/internal/workspace"
)

type ProcessInput struct {
	Manifest  manifest.Manifest
	Operation operation.InfrastructureOperation
	Catalog   manifest.Catalog
	Workspace workspace.Workspace
}

type KindProcessor interface {
	Process(input ProcessInput) error
}

type ManifestProcessor[S any] interface {
	Process(input ProcessInput, spec S) error
}
