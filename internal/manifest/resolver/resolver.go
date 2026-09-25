package resolver

import (
	"fmt"

	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/manifest/registry"
	"github.com/expram/orchestra/internal/validation"
)

const (
	SpecTag   = "manifest"
	specScope = "spec"
)

type SpecFactory func() Migration

func RegisterSpec[S any, P interface {
	*S
	Migration
}](typ manifest.Type) registry.Registration[manifest.Type, SpecFactory] {
	return registry.Register(typ, SpecFactory(func() Migration { return P(new(S)) }))
}

type Resolver struct {
	decoder SpecDecoder
	specs   registry.Registry[manifest.Type, SpecFactory]
}

func NewResolver(decoder SpecDecoder, specs registry.Registry[manifest.Type, SpecFactory]) Resolver {
	return Resolver{decoder: decoder, specs: specs}
}

func (r Resolver) Resolve(unresolved []manifest.UnresolvedManifest) (manifest.Catalog, error) {
	var (
		catalog  manifest.Catalog
		problems errscope.Problems
	)

	for _, candidate := range unresolved {
		resolved, err := r.resolve(candidate)
		if err != nil {
			problems.Add(errscope.In(candidate.Identity().String(), err))

			continue
		}

		catalog = append(catalog, resolved)
	}

	if err := problems.Err(); err != nil {
		return nil, err
	}

	return catalog, nil
}

func (r Resolver) resolve(candidate manifest.UnresolvedManifest) (manifest.Manifest, error) {
	newSpec, ok := r.specs.Lookup(candidate.Type())
	if !ok {
		return manifest.Manifest{}, validation.Error{Problems: []error{fmt.Errorf(
			"unsupported kind %q in api version %q", candidate.Kind, candidate.APIVersion,
		)}}
	}

	spec := newSpec()

	if err := r.decoder.Decode(candidate.Spec, spec); err != nil {
		return manifest.Manifest{}, errscope.In(specScope, validation.Error{Problems: []error{err}})
	}

	if err := spec.Validate(); err != nil {
		return manifest.Manifest{}, errscope.In(specScope, validation.Error{Problems: []error{err}})
	}

	core, err := spec.Convert()
	if err != nil {
		return manifest.Manifest{}, errscope.In(specScope, err)
	}

	return manifest.Manifest{
		Type:     candidate.Type(),
		Name:     candidate.Name,
		Metadata: candidate.Metadata,
		Spec:     core,
	}, nil
}
