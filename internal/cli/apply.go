package cli

import (
	"github.com/expram/orchestra/internal/operation"
	"github.com/spf13/cobra"
)

type applyCommand struct {
	baseInfrastructureCommand
}

var _ InfrastructureCommand = applyCommand{}

func (applyCommand) Operation() operation.InfrastructureOperation {
	return operation.InfrastructureOperation{
		Mode: operation.ModeApply,
		Type: operation.TypeProvision,
	}
}

func NewApplyCommand(useCase operation.CallInfrastructureOperationUseCase) *cobra.Command {
	c := &applyCommand{}

	return newInfrastructureCommand(
		"apply <builtin-directory> [user-manifests-directory]",
		"Apply infrastructure",
		&c.baseInfrastructureCommand,
		c,
		useCase,
	)
}
