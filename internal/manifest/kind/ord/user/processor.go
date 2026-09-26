package user

import (
	"fmt"

	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/manifest/registry"
	"github.com/expram/orchestra/internal/operation"
	"github.com/expram/orchestra/internal/workspace"
)

type Processor struct {
	registry *registry.Registry[manifest.Type, manifest.Manifest]
}

func NewProcessor(registry *registry.Registry[manifest.Type, manifest.Manifest]) *Processor {
	return &Processor{registry: registry}
}

func (p *Processor) Process(op operation.InfrastructureOperation, unresolvedUser []manifest.UnresolvedManifest, ws workspace.Workspace) error {
	manifests := make(map[manifest.Type]manifest.UnresolvedManifest)
	registered := p.registry.Get()
	for _, unresolvedManifest := range unresolvedUser {
		utype := manifest.Type{
			Kind:       unresolvedManifest.Kind,
			APIVersion: unresolvedManifest.APIVersion,
		}

		_, ok := registered[utype]
		if !ok {
			return fmt.Errorf("unknown manifest type with kind %s, apiVersion %s and name %s", utype.Kind, utype.APIVersion, unresolvedManifest.Name)
		}

		manifests[utype] = unresolvedManifest
	}

	for k, v := range manifests {
		fmt.Println(k, v)
	}

	return nil
}
