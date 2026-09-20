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

func NewApplyCommand() *cobra.Command {
	c := &applyCommand{}

	return newInfrastructureCommand(
		"apply <directories...>",
		"Apply infrastructure",
		&c.baseInfrastructureCommand,
		func() error {
			return nil
		},
	)
}
