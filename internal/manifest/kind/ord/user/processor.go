package user

import (
	"fmt"

	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/manifest/registry"
	"github.com/expram/orchestra/internal/operation"
	"github.com/expram/orchestra/internal/validation"
	"github.com/expram/orchestra/internal/workspace"
)

type Processor struct {
	registry *registry.Registry[manifest.Type, manifest.Manifest]
}

func NewProcessor(registry *registry.Registry[manifest.Type, manifest.Manifest]) Processor {
	return Processor{registry: registry}
}

func (p Processor) Process(op operation.InfrastructureOperation, unresolvedUser []manifest.UnresolvedManifest, ws workspace.Workspace) error {
	var problems errscope.Problems

	manifests := make(map[manifest.Type]manifest.UnresolvedManifest)
	for _, unresolvedManifest := range unresolvedUser {
		utype := unresolvedManifest.Type()

		if _, ok := p.registry.Lookup(utype); !ok {
			problems.Add(errscope.In(unresolvedManifest.Identity().String(), validation.Error{Problems: []error{fmt.Errorf(
				"unsupported kind %q in api version %q", utype.Kind, utype.APIVersion,
			)}}))

			continue
		}

		manifests[utype] = unresolvedManifest
	}

	return problems.Err()
}
