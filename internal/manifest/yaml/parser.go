package yaml

import (
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/validation"
	yamldoc "github.com/expram/orchestra/internal/yaml"
)

type document struct {
	APIVersion string         `yaml:"apiVersion"`
	Kind       string         `yaml:"kind"`
	Metadata   map[string]any `yaml:"metadata"`
	Spec       map[string]any `yaml:"spec"`
	Unknown    map[string]any `yaml:",inline"`
}

type YamlUnresolvedManifestParser struct {
	validator UnresolvedManifestValidator
}

var _ UnresolvedManifestParser = YamlUnresolvedManifestParser{}

func NewYamlUnresolvedManifestParser(validator UnresolvedManifestValidator) YamlUnresolvedManifestParser {
	return YamlUnresolvedManifestParser{validator: validator}
}

func (p YamlUnresolvedManifestParser) Parse(data []byte) ([]manifest.UnresolvedManifest, error) {
	documents, err := yamldoc.Split(data)
	if err != nil {
		return nil, validation.Failed(yamldoc.Errors(err)...).Err(nil)
	}

	manifests := make([]manifest.UnresolvedManifest, 0, len(documents))

	for _, doc := range documents {
		unresolved, err := p.parse(doc)
		if err != nil {
			return nil, err
		}

		manifests = append(manifests, unresolved)
	}

	return manifests, nil
}

func (p YamlUnresolvedManifestParser) parse(doc yamldoc.Document) (manifest.UnresolvedManifest, error) {
	source := newSource(doc.Index, doc.Line)

	if !doc.Mapping() {
		return manifest.UnresolvedManifest{}, validation.Failed(errNotAnObject).Err(source)
	}

	var decoded document
	if err := doc.Decode(&decoded); err != nil {
		return manifest.UnresolvedManifest{}, validation.Failed(yamldoc.Errors(err)...).Err(source)
	}

	if result := p.validator.Validate(decoded); !result.Valid() {
		return manifest.UnresolvedManifest{}, result.Err(source)
	}

	return manifest.NewUnresolvedManifest(
		source,
		decoded.Kind,
		decoded.APIVersion,
		decoded.Metadata,
		decoded.Spec,
	)
}
