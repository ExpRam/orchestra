package dto

import (
	"errors"
	"fmt"
	"maps"
	"slices"

	"github.com/expram/orchestra/internal/errscope"
)

const pathSeparator = "."

type manifestKeys struct {
	Kind       string
	APIVersion string
	Metadata   string
	Name       string
	Spec       string
}

var keys = manifestKeys{
	Kind:       "kind",
	APIVersion: "apiVersion",
	Metadata:   "metadata",
	Name:       "name",
	Spec:       "spec",
}

var knownKeys = []string{keys.Kind, keys.APIVersion, keys.Metadata, keys.Spec}

var errNotAnObject = errors.New("manifest must be an object")

type KRMUnresolvedManifestValidator struct{}

var _ UnresolvedManifestValidator = KRMUnresolvedManifestValidator{}

func NewKRMUnresolvedManifestValidator() KRMUnresolvedManifestValidator {
	return KRMUnresolvedManifestValidator{}
}

func (v KRMUnresolvedManifestValidator) Validate(tree any) (UnresolvedManifest, error) {
	root, ok := tree.(map[string]any)
	if !ok {
		return UnresolvedManifest{}, errNotAnObject
	}

	var problems errscope.Problems

	kind := text(root[keys.Kind], keys.Kind, &problems)
	apiVersion := text(root[keys.APIVersion], keys.APIVersion, &problems)
	metadata := object(root[keys.Metadata], keys.Metadata, &problems)
	name := text(metadata[keys.Name], keys.Metadata+pathSeparator+keys.Name, &problems)
	spec := object(root[keys.Spec], keys.Spec, &problems)

	for _, key := range slices.Sorted(maps.Keys(root)) {
		if !slices.Contains(knownKeys, key) {
			problems.Add(fmt.Errorf("unknown field %q", key))
		}
	}

	return UnresolvedManifest{
		Kind:       kind,
		APIVersion: apiVersion,
		Name:       name,
		Metadata:   metadata,
		Spec:       spec,
	}, problems.Err()
}

func text(value any, path string, problems *errscope.Problems) string {
	if value == nil {
		return ""
	}

	s, ok := value.(string)
	if !ok {
		problems.Add(fmt.Errorf("%q must be a string", path))
	}

	return s
}

func object(value any, path string, problems *errscope.Problems) map[string]any {
	if value == nil {
		return nil
	}

	m, ok := value.(map[string]any)
	if !ok {
		problems.Add(fmt.Errorf("%q must be an object", path))
	}

	return m
}
