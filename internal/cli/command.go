package cli

import (
	"github.com/expram/orchestra/internal/operation"
	"github.com/spf13/cobra"
)

type InfrastructureCommand interface {
	Operation() operation.InfrastructureOperation
}

type baseInfrastructureCommand struct {
	builtInDirectories []string
	filterTokens       []string
}

func newInfrastructureCommand(
	use string,
	short string,
	c *baseInfrastructureCommand,
	run func() error,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c.builtInDirectories = args
			return run()
		},
	}

	c.configure(cmd)

	return cmd
}

func (c *baseInfrastructureCommand) configure(cmd *cobra.Command) {
	cmd.Args = cobra.MinimumNArgs(1)

	cmd.Flags().StringSliceVarP(
		&c.filterTokens,
		"filter",
		"f",
		nil,
		"Filter user manifests by metadata values",
	)
}

func (c *baseInfrastructureCommand) bindArgs(args []string) {
	c.builtInDirectories = args
}
