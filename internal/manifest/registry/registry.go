package registry

import (
	"fmt"

	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/validation"
)

const (
	SpecTag   = "manifest"
	specScope = "spec"
)

type Registration struct {
	typ     manifest.Type
	newSpec func() Migration
}

func Register[S any, P interface {
	*S
	Migration
}](typ manifest.Type) Registration {
	return Registration{
		typ:     typ,
		newSpec: func() Migration { return P(new(S)) },
	}
}

type Registry struct {
	decoder SpecDecoder
	specs   map[manifest.Type]func() Migration
}

func NewRegistry(decoder SpecDecoder, registrations ...Registration) Registry {
	specs := make(map[manifest.Type]func() Migration, len(registrations))

	for _, registration := range registrations {
		if _, ok := specs[registration.typ]; ok {
			panic(fmt.Sprintf(
				"manifest kind %q in api version %q is registered twice",
				registration.typ.Kind, registration.typ.APIVersion,
			))
		}

		specs[registration.typ] = registration.newSpec
	}

	return Registry{decoder: decoder, specs: specs}
}

func (r Registry) Resolve(unresolved []manifest.UnresolvedManifest) ([]manifest.Manifest, error) {
	var (
		resolved []manifest.Manifest
		problems errscope.Problems
	)

	for _, candidate := range unresolved {
		resolvedManifest, err := r.resolve(candidate)
		if err != nil {
			problems.Add(errscope.In(candidate.Identity().String(), err))

			continue
		}

		resolved = append(resolved, resolvedManifest)
	}

	if err := problems.Err(); err != nil {
		return nil, err
	}

	return resolved, nil
}

func (r Registry) resolve(candidate manifest.UnresolvedManifest) (manifest.Manifest, error) {
	newSpec, ok := r.specs[candidate.Type()]
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
