package cli

import (
	"github.com/spf13/cobra"

	"github.com/expram/orchestra/internal/operation"
)

type planCommand struct {
	baseInfrastructureCommand
	destroy bool
}

var _ InfrastructureCommand = (*planCommand)(nil)

func (c *planCommand) Operation() operation.InfrastructureOperation {
	opType := operation.TypeProvision

	if c.destroy {
		opType = operation.TypeDeprovision
	}

	return operation.InfrastructureOperation{
		Mode: operation.ModePlan,
		Type: opType,
	}
}

func NewPlanCommand(useCase operation.CallInfrastructureOperationUseCase) *cobra.Command {
	c := &planCommand{}

	cmd := newInfrastructureCommand(
		"plan <builtin-directory> [user-manifests-directory]",
		"Show planned infrastructure changes",
		&c.baseInfrastructureCommand,
		c,
		useCase,
	)

	cmd.Flags().BoolVar(
		&c.destroy,
		"destroy",
		false,
		"Plan infrastructure destruction",
	)

	return cmd
}
