package yaml

import (
	"github.com/expram/orchestra/internal/contract"
	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/manifest/dto"
	"github.com/expram/orchestra/internal/validation"
	yamldoc "github.com/expram/orchestra/internal/yaml"
)

type YamlUnresolvedManifestParser struct {
	validator dto.UnresolvedManifestValidator
}

var _ contract.Parser[manifest.UnresolvedManifest] = YamlUnresolvedManifestParser{}

func NewYamlUnresolvedManifestParser(validator dto.UnresolvedManifestValidator) YamlUnresolvedManifestParser {
	return YamlUnresolvedManifestParser{validator: validator}
}

func (p YamlUnresolvedManifestParser) Parse(data []byte) ([]manifest.UnresolvedManifest, error) {
	documents, err := yamldoc.Split(data)
	if err != nil {
		return nil, validation.Error{Problems: []error{err}}
	}

	var problems errscope.Problems

	manifests := make([]manifest.UnresolvedManifest, 0, len(documents))

	for _, doc := range documents {
		unresolved, err := p.parse(doc)
		if err != nil {
			problems.Add(err)

			continue
		}

		manifests = append(manifests, unresolved)
	}

	if err := problems.Err(); err != nil {
		return nil, err
	}

	return manifests, nil
}

func (p YamlUnresolvedManifestParser) parse(doc yamldoc.Document) (manifest.UnresolvedManifest, error) {
	object, err := doc.Object()
	if err != nil {
		return manifest.UnresolvedManifest{}, validation.Error{Problems: []error{err}}
	}

	valid, err := p.validator.Validate(object)
	if err != nil {
		return manifest.UnresolvedManifest{}, errscope.In(doc.Location(), validation.Error{Problems: []error{err}})
	}

	unresolved, err := manifest.NewUnresolvedManifest(valid.Kind, valid.APIVersion, valid.Name, valid.Metadata, valid.Spec)
	if err != nil {
		return manifest.UnresolvedManifest{}, errscope.In(doc.Location(), err)
	}

	return unresolved, nil
}
