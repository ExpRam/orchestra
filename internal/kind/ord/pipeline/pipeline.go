package pipeline

import (
	"fmt"

	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/kind/ord"
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/operation"
	"github.com/expram/orchestra/internal/registry"
	"github.com/expram/orchestra/internal/validation"
	"github.com/expram/orchestra/internal/workspace"
)

const specScope = "spec"

type OrdPipeline struct {
	registry  *registry.Registry[manifest.Type, ord.Spec]
	validator SchemaValidator
}

func NewOrdPipeline(registry *registry.Registry[manifest.Type, ord.Spec], validator SchemaValidator) OrdPipeline {
	return OrdPipeline{registry: registry, validator: validator}
}

func (p OrdPipeline) Process(op operation.InfrastructureOperation, unresolvedUser []manifest.UnresolvedManifest, ws workspace.Workspace) error {
	var problems errscope.Problems

	manifests := make(map[manifest.Type]manifest.UnresolvedManifest)
	for _, unresolvedManifest := range unresolvedUser {
		if err := p.validate(unresolvedManifest); err != nil {
			problems.Add(errscope.In(unresolvedManifest.Identity().String(), err))

			continue
		}

		manifests[unresolvedManifest.Type()] = unresolvedManifest
	}

	return problems.Err()
}

func (p OrdPipeline) validate(unresolvedManifest manifest.UnresolvedManifest) error {
	utype := unresolvedManifest.Type()

	spec, ok := p.registry.Lookup(utype)
	if !ok {
		return validation.Error{Problems: []error{fmt.Errorf(
			"unsupported kind %q in api version %q", utype.Kind, utype.APIVersion,
		)}}
	}

	if err := p.validator.Validate(spec.Schema, unresolvedManifest.RawSpec); err != nil {
		return errscope.In(specScope, validation.Error{Problems: []error{err}})
	}

	return nil
}
