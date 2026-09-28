package processor

import (
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/operation"
	"github.com/expram/orchestra/internal/workspace"
)

type HandleInput struct {
	Manifest  manifest.Manifest
	Operation operation.InfrastructureOperation
	Catalog   manifest.Catalog
	Workspace workspace.Workspace
}

type KindHandler interface {
	Handle(input HandleInput) error
}

type SpecHandler[S any] interface {
	Handle(input HandleInput, spec S) error
}
