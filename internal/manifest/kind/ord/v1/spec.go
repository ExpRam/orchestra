package v1

import (
	"fmt"

	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/manifest/kind/ord"
)

var Type = manifest.Type{Kind: ord.Kind, APIVersion: "v1"}

type Execution struct {
	Cardinality string `manifest:"cardinality"`
}

type Annotation struct {
	Name      string         `manifest:"name"`
	Type      string         `manifest:"type"`
	Settings  map[string]any `manifest:"settings"`
	Execution Execution      `manifest:"execution"`
}

type Spec struct {
	Schema      map[string]any `manifest:"schema"`
	Annotations []Annotation   `manifest:"annotations"`
}

func (s Spec) Validate() error {
	return nil
}

func (s Spec) Convert() (any, error) {
	var (
		annotations []ord.Annotation
		problems    errscope.Problems
	)

	for index, source := range s.Annotations {
		annotation, err := source.convert()
		if err != nil {
			problems.Add(errscope.In(fmt.Sprintf("annotations[%d]", index), err))

			continue
		}

		annotations = append(annotations, annotation)
	}

	spec, err := ord.NewSpec(s.Schema, annotations)
	problems.Add(err)

	if len(problems) > 0 {
		return nil, manifest.Error{Problems: problems}
	}

	return spec, nil
}

func (a Annotation) convert() (ord.Annotation, error) {
	var problems errscope.Problems

	cardinality, cardinalityErr := ord.NewCardinality(a.Execution.Cardinality)

	annotation, err := ord.NewAnnotation(a.Name, a.Type, a.Settings, ord.NewExecution(cardinality))
	problems.Add(err)
	problems.Add(cardinalityErr)

	return annotation, problems.Err()
}
