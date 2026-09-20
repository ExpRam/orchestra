package cli

import (
	"github.com/expram/orchestra/internal/operation"
	"github.com/spf13/cobra"
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

func NewPlanCommand() *cobra.Command {
	c := &planCommand{}

	cmd := newInfrastructureCommand(
		"plan <directories...>",
		"Show planned infrastructure changes",
		&c.baseInfrastructureCommand,
		func() error {
			return nil
		},
	)

	cmd.Flags().BoolVar(
		&c.destroy,
		"destroy",
		false,
		"Plan infrastructure destruction",
	)

	return cmd
}
