package processor

import (
	"fmt"

	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/manifest/kind/ord"
	"github.com/expram/orchestra/internal/operation"
)

type OrdProcessor struct{}

func NewOrdProcessor() OrdProcessor {
	return OrdProcessor{}
}

func (p OrdProcessor) Process(input operation.ProcessInput) error {
	spec := input.Manifest.Spec.(ord.Spec)

	identity := manifest.Type{
		Kind:       manifest.Kind(input.Manifest.Name),
		APIVersion: manifest.APIVersion(spec.ApiVersion),
	}

	fmt.Println(identity)
	return nil
}
