package processor

import (
	"fmt"

	"github.com/expram/orchestra/internal/operation"
)

type OrdProcessor struct{}

func NewOrdProcessor() OrdProcessor {
	return OrdProcessor{}
}

func (p OrdProcessor) Process(input operation.ProcessInput) error {
	fmt.Println("Hello From ORD Processor")
	return nil
}
