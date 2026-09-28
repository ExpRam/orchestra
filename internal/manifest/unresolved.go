package manifest

import (
	"errors"

	"github.com/expram/orchestra/internal/errscope"
)

var ErrMissingSpec = errors.New("spec must be specified")

type RawSpec = map[string]any

type UnresolvedManifest struct {
	Kind       Kind
	APIVersion APIVersion
	Name       Name
	Metadata   Metadata
	RawSpec    RawSpec
}

func (m UnresolvedManifest) Type() Type {
	return Type{Kind: m.Kind, APIVersion: m.APIVersion}
}

func (m UnresolvedManifest) Identity() Identity {
	return Identity{Kind: m.Type().Kind, Name: m.Name}
}

func NewUnresolvedManifest(
	kind, apiVersion, name string,
	metadata Metadata,
	spec RawSpec,
) (UnresolvedManifest, error) {
	var problems errscope.Problems

	manifestKind, err := NewKind(kind)
	problems.Add(err)

	manifestAPIVersion, err := NewAPIVersion(apiVersion)
	problems.Add(err)

	manifestName, err := NewName(name)
	problems.Add(err)

	if spec == nil {
		problems.Add(ErrMissingSpec)
	}

	if len(problems) > 0 {
		return UnresolvedManifest{}, Error{Problems: problems}
	}

	return UnresolvedManifest{
		Kind:       manifestKind,
		APIVersion: manifestAPIVersion,
		Name:       manifestName,
		Metadata:   metadata,
		RawSpec:    spec,
	}, nil
}
