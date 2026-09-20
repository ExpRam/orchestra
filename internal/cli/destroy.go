package cli

import (
	"github.com/expram/orchestra/internal/operation"
	"github.com/spf13/cobra"
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

func NewDestroyCommand() *cobra.Command {
	c := &destroyCommand{}

	return newInfrastructureCommand(
		"destroy <directories...>",
		"Destroy infrastructure",
		&c.baseInfrastructureCommand,
		func() error {
			return nil
		},
	)
}
