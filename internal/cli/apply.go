package cli

import (
	"github.com/expram/orchestra/internal/operation"
	"github.com/spf13/cobra"
)

type applyCommand struct {
	baseInfrastructureCommand
}

var _ applyCommand = applyCommand{}

func (applyCommand) Operation() operation.InfrastructureOperation {
	return operation.InfrastructureOperation{
		Mode: operation.ModeApply,
		Type: operation.TypeProvision,
	}
}

func newApplyCommand() *cobra.Command {
	c := &applyCommand{}

	return newInfrastructureCommand(
		"apply <directories...>",
		"Apply infrastructure",
		&c.baseInfrastructureCommand,
		func() error {
			// usecase.Execute(c)
			return nil
		},
	)
}

func init() {
	rootCmd.AddCommand(newApplyCommand())
}
