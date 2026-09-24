package ord

import (
	"errors"
	"strings"

	"github.com/expram/orchestra/internal/manifest"
)

const Kind manifest.Kind = "OrchestraResourceDefinition"

var (
	ErrMissingSchema       = errors.New("schema must be specified")
	ErrEmptyAnnotationName = errors.New("annotation name must not be empty")
	ErrEmptyAnnotationType = errors.New("annotation type must not be empty")
	ErrEmptyCardinality    = errors.New("cardinality must not be empty")
)

type Cardinality string

func NewCardinality(value string) (Cardinality, error) {
	if strings.TrimSpace(value) == "" {
		return "", ErrEmptyCardinality
	}

	return Cardinality(value), nil
}

type Execution struct {
	Cardinality Cardinality
}

func NewExecution(cardinality Cardinality) Execution {
	return Execution{Cardinality: cardinality}
}

type Annotation struct {
	Name      string
	Type      string
	Settings  map[string]any
	Execution Execution
}

func NewAnnotation(name string, aType string, settings map[string]any, execution Execution) (Annotation, error) {
	if strings.TrimSpace(name) == "" {
		return Annotation{}, ErrEmptyAnnotationName
	}

	if strings.TrimSpace(aType) == "" {
		return Annotation{}, ErrEmptyAnnotationType
	}

	return Annotation{
		Name:      name,
		Type:      aType,
		Settings:  settings,
		Execution: execution,
	}, nil
}

type Spec struct {
	Schema      map[string]any
	Annotations []Annotation
}

func NewSpec(schema map[string]any, annotations []Annotation) (Spec, error) {
	if schema == nil {
		return Spec{}, ErrMissingSchema
	}

	return Spec{
		Schema:      schema,
		Annotations: annotations,
	}, nil
}
