package manifest

import (
	"errors"
	"fmt"
	"strings"

	"github.com/expram/orchestra/internal/errscope"
)

var (
	ErrEmptyKind       = errors.New("kind must not be empty")
	ErrEmptyAPIVersion = errors.New("api version must not be empty")
	ErrEmptyName       = errors.New("name must not be empty")
	ErrMissingSpec     = errors.New("spec must be specified")
)

type Kind string

func NewKind(value string) (Kind, error) {
	if strings.TrimSpace(value) == "" {
		return "", ErrEmptyKind
	}

	return Kind(value), nil
}

func (k Kind) String() string {
	return string(k)
}

type APIVersion string

func NewAPIVersion(value string) (APIVersion, error) {
	if strings.TrimSpace(value) == "" {
		return "", ErrEmptyAPIVersion
	}

	group, version, qualified := strings.Cut(value, "/")
	if qualified && (group == "" || version == "" || strings.Contains(version, "/")) {
		return "", fmt.Errorf("api version %q must be <version> or <group>/<version>", value)
	}

	return APIVersion(value), nil
}

func (v APIVersion) String() string {
	return string(v)
}

type Name string

func NewName(value string) (Name, error) {
	if strings.TrimSpace(value) == "" {
		return "", ErrEmptyName
	}

	return Name(value), nil
}

func (n Name) String() string {
	return string(n)
}

type UnresolvedManifest struct {
	Kind       Kind
	APIVersion APIVersion
	Name       Name
	Metadata   Metadata
	Spec       Spec
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
	spec Spec,
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
		Spec:       spec,
	}, nil
}
