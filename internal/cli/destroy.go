package cli

import (
	"github.com/spf13/cobra"

	"github.com/expram/orchestra/internal/operation"
)

type destroyCommand struct {
	baseInfrastructureCommand
}

var _ InfrastructureCommand = destroyCommand{}

func (destroyCommand) Operation() operation.InfrastructureOperation {
	return operation.InfrastructureOperation{
		Mode: operation.ModeApply,
		Type: operation.TypeDeprovision,
	}
}

func NewDestroyCommand(useCase operation.CallInfrastructureOperationUseCase) *cobra.Command {
	c := &destroyCommand{}

	return newInfrastructureCommand(
		"destroy <builtin-directory> [user-manifests-directory]",
		"Destroy infrastructure",
		&c.baseInfrastructureCommand,
		c,
		useCase,
	)
}
