package processor

import (
	"fmt"

	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/manifest/registry"
	"github.com/expram/orchestra/internal/operation"
	"github.com/expram/orchestra/internal/workspace"
)

type Processor struct {
	processors registry.Registry[manifest.Kind, operation.ManifestProcessor]
}

func NewProcessor(processors registry.Registry[manifest.Kind, operation.ManifestProcessor]) Processor {
	return Processor{processors: processors}
}

func (p Processor) Process(op operation.InfrastructureOperation, catalog manifest.Catalog, ws workspace.Workspace) error {
	for _, m := range catalog {
		if err := p.process(op, m, catalog, ws); err != nil {
			return errscope.In(m.Identity().String(), err)
		}
	}

	return nil
}

func (p Processor) process(op operation.InfrastructureOperation, m manifest.Manifest, catalog manifest.Catalog, ws workspace.Workspace) error {
	processor, ok := p.processors.Lookup(m.Type.Kind)
	if !ok {
		return fmt.Errorf("no processor for kind %q", m.Type.Kind)
	}

	return processor.Process(operation.ProcessInput{
		Manifest:  m,
		Operation: op,
		Catalog:   catalog,
		Workspace: ws,
	})
}
