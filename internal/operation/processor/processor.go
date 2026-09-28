package processor

import (
	"fmt"

	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/operation"
	"github.com/expram/orchestra/internal/registry"
	"github.com/expram/orchestra/internal/workspace"
)

type Processor struct {
	handlers *registry.Registry[manifest.Kind, KindHandler]
}

// typedHandler Art of Kludge-Oriented Programming.
type typedHandler[S any] struct {
	handler SpecHandler[S]
}

func Register[S any](kind manifest.Kind, handler SpecHandler[S]) registry.Registration[manifest.Kind, KindHandler] {
	return registry.Register[manifest.Kind, KindHandler](kind, typedHandler[S]{handler: handler})
}

func NewProcessor(handlers *registry.Registry[manifest.Kind, KindHandler]) Processor {
	return Processor{handlers: handlers}
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
	handler, ok := p.handlers.Lookup(m.Type.Kind)
	if !ok {
		return fmt.Errorf("no handler for kind %q", m.Type.Kind)
	}

	return handler.Handle(HandleInput{
		Manifest:  m,
		Operation: op,
		Catalog:   catalog,
		Workspace: ws,
	})
}

func (w typedHandler[S]) Handle(input HandleInput) error {
	spec, ok := input.Manifest.Spec.(S)
	if !ok {
		var excepted S

		return fmt.Errorf("cannot handle spec %T, excepted %T", input.Manifest.Spec, excepted)
	}

	return w.handler.Handle(input, spec)
}
