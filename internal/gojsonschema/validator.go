package gojsonschema

import (
	"bytes"
	"encoding/json"

	gojsonschemav1 "github.com/xeipuuv/gojsonschema"

	"github.com/expram/orchestra/internal/kind/ord/pipeline"
)

type SchemaValidator struct{}

var _ pipeline.SchemaValidator = SchemaValidator{}

func NewSchemaValidator() SchemaValidator {
	return SchemaValidator{}
}

func (v SchemaValidator) Validate(schema map[string]any, document map[string]any) error {
	jsonSchema, err := normalize(schema)
	if err != nil {
		return err
	}

	jsonDocument, err := normalize(document)
	if err != nil {
		return err
	}

	result, err := gojsonschemav1.Validate(
		gojsonschemav1.NewRawLoader(jsonSchema),
		gojsonschemav1.NewRawLoader(jsonDocument),
	)
	if err != nil {
		return err
	}

	return translate(result.Errors())
}

func normalize(value map[string]any) (any, error) {
	encoded, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(bytes.NewReader(encoded))
	decoder.UseNumber()

	var normalized any
	if err := decoder.Decode(&normalized); err != nil {
		return nil, err
	}

	return normalized, nil
}
