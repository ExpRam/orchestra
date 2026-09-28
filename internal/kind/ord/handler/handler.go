package handler

import (
	"fmt"

	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/kind/ord"
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/operation/processor"
	"github.com/expram/orchestra/internal/registry"
)

type Handler struct {
	registry *registry.Registry[manifest.Type, ord.Spec]
}

var _ processor.SpecHandler[ord.Spec] = Handler{}

func NewHandler(registry *registry.Registry[manifest.Type, ord.Spec]) Handler {
	return Handler{registry: registry}
}

func (h Handler) Handle(input processor.HandleInput, spec ord.Spec) error {
	identity := manifest.Type{
		Kind:       manifest.Kind(input.Manifest.Name),
		APIVersion: manifest.APIVersion(spec.APIVersion),
	}

	if err := h.registry.Add(registry.Register(identity, spec)); err != nil {
		return manifest.Error{Problems: errscope.Problems{fmt.Errorf(
			"kind %q in api version %q %w", identity.Kind, identity.APIVersion, err,
		)}}
	}

	return nil
}
