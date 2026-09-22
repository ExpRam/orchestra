package manifest

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyKind       = errors.New("kind must not be empty")
	ErrEmptyAPIVersion = errors.New("api version must not be empty")
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

type UnresolvedManifest struct {
	Source     Source
	Kind       Kind
	APIVersion APIVersion
	Metadata   Metadata
	Spec       Spec
}

func (m UnresolvedManifest) WithSource(source Source) UnresolvedManifest {
	m.Source = source

	return m
}

func NewUnresolvedManifest(
	source Source,
	kind, apiVersion string,
	metadata Metadata,
	spec Spec,
) (UnresolvedManifest, error) {
	var problems []error

	manifestKind, err := NewKind(kind)
	if err != nil {
		problems = append(problems, err)
	}

	manifestAPIVersion, err := NewAPIVersion(apiVersion)
	if err != nil {
		problems = append(problems, err)
	}

	if len(problems) > 0 {
		return UnresolvedManifest{}, Error{Source: source, Problems: problems}
	}

	return UnresolvedManifest{
		Source:     source,
		Kind:       manifestKind,
		APIVersion: manifestAPIVersion,
		Metadata:   metadata,
		Spec:       spec,
	}, nil
}
