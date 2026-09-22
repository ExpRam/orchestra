package yaml

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/expram/orchestra/internal/validation"
)

const fieldName = "name"

var (
	errNotAnObject       = errors.New("manifest must be an object")
	errMissingKind       = errors.New(`"kind" is required`)
	errMissingAPIVersion = errors.New(`"apiVersion" is required`)
	errMissingMetadata   = errors.New(`"metadata" is required`)
	errMissingName       = errors.New(`"metadata.name" is required`)
	errInvalidName       = errors.New(`"metadata.name" must be a non-empty string`)
	errMissingSpec       = errors.New(`"spec" is required`)
)

type YamlUnresolvedManifestValidator struct{}

var _ UnresolvedManifestValidator = YamlUnresolvedManifestValidator{}

func NewYamlUnresolvedManifestValidator() YamlUnresolvedManifestValidator {
	return YamlUnresolvedManifestValidator{}
}

func (v YamlUnresolvedManifestValidator) Validate(decoded document) validation.Result {
	var problems []error

	if strings.TrimSpace(decoded.Kind) == "" {
		problems = append(problems, errMissingKind)
	}

	if strings.TrimSpace(decoded.APIVersion) == "" {
		problems = append(problems, errMissingAPIVersion)
	}

	problems = append(problems, validateMetadata(decoded.Metadata)...)

	if decoded.Spec == nil {
		problems = append(problems, errMissingSpec)
	}

	for _, field := range slices.Sorted(maps.Keys(decoded.Unknown)) {
		problems = append(problems, fmt.Errorf("unknown field %q", field))
	}

	return validation.Failed(problems...)
}

func validateMetadata(metadata map[string]any) []error {
	if metadata == nil {
		return []error{errMissingMetadata}
	}

	value, ok := metadata[fieldName]
	if !ok {
		return []error{errMissingName}
	}

	name, ok := value.(string)
	if !ok || strings.TrimSpace(name) == "" {
		return []error{errInvalidName}
	}

	return nil
}
