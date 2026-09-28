package processor

import (
	"fmt"

	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/kind/ord"
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/operation/processor"
	"github.com/expram/orchestra/internal/registry"
)

type OrdProcessor struct {
	registry *registry.Registry[manifest.Type, ord.Spec]
}

func NewOrdProcessor(registry *registry.Registry[manifest.Type, ord.Spec]) OrdProcessor {
	return OrdProcessor{registry: registry}
}

func (p OrdProcessor) Process(input processor.ProcessInput, spec ord.Spec) error {
	identity := manifest.Type{
		Kind:       manifest.Kind(input.Manifest.Name),
		APIVersion: manifest.APIVersion(spec.ApiVersion),
	}

	if err := p.registry.Add(registry.Register(identity, spec)); err != nil {
		return manifest.Error{Problems: errscope.Problems{fmt.Errorf(
			"kind %q in api version %q %w", identity.Kind, identity.APIVersion, err,
		)}}
	}

	return nil
}
