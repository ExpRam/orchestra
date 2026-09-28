package gojsonschema

import (
	"errors"
	"fmt"

	gojsonschemav1 "github.com/xeipuuv/gojsonschema"

	"github.com/expram/orchestra/internal/errscope"
)

func translate(resultErrors []gojsonschemav1.ResultError) error {
	var problems errscope.Problems

	for _, resultError := range resultErrors {
		problems.Add(describe(resultError))
	}

	return problems.Err()
}

func describe(resultError gojsonschemav1.ResultError) error {
	if resultError.Field() == gojsonschemav1.STRING_ROOT_SCHEMA_PROPERTY {
		return errors.New(resultError.Description())
	}

	return fmt.Errorf("%q: %s", resultError.Field(), resultError.Description())
}
