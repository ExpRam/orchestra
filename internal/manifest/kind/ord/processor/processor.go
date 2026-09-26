package processor

import (
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/manifest/kind/ord"
	"github.com/expram/orchestra/internal/manifest/registry"
	"github.com/expram/orchestra/internal/operation"
)

type OrdProcessor struct {
	registry *registry.Registry[manifest.Type, manifest.Manifest]
}

func NewOrdProcessor(registry *registry.Registry[manifest.Type, manifest.Manifest]) OrdProcessor {
	return OrdProcessor{registry: registry}
}

func (p OrdProcessor) Process(input operation.ProcessInput, spec ord.Spec) error {
	identity := manifest.Type{
		Kind:       manifest.Kind(input.Manifest.Name),
		APIVersion: manifest.APIVersion(spec.ApiVersion),
	}

	p.registry.Add(registry.Register(identity, input.Manifest))
	return nil
}
