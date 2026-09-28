package handler

import (
	"fmt"

	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/kind/ord"
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/operation/processor"
	"github.com/expram/orchestra/internal/registry"
)

type OrdHandler struct {
	registry *registry.Registry[manifest.Type, ord.Spec]
}

func NewOrdHandler(registry *registry.Registry[manifest.Type, ord.Spec]) OrdHandler {
	return OrdHandler{registry: registry}
}

func (h OrdHandler) Handle(input processor.HandleInput, spec ord.Spec) error {
	identity := manifest.Type{
		Kind:       manifest.Kind(input.Manifest.Name),
		APIVersion: manifest.APIVersion(spec.ApiVersion),
	}

	if err := h.registry.Add(registry.Register(identity, spec)); err != nil {
		return manifest.Error{Problems: errscope.Problems{fmt.Errorf(
			"kind %q in api version %q %w", identity.Kind, identity.APIVersion, err,
		)}}
	}

	return nil
}
