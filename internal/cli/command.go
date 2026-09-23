package cli

import (
	"github.com/expram/orchestra/internal/operation"
	"github.com/spf13/cobra"
)

type InfrastructureCommand interface {
	Operation() operation.InfrastructureOperation
}

type baseInfrastructureCommand struct {
	builtinDirectory string
	userDirectory    string
	filterTokens     []string
	keepWorkspace    bool
}

func newInfrastructureCommand(
	use string,
	short string,
	c *baseInfrastructureCommand,
	infraCmd InfrastructureCommand,
	useCase operation.CallInfrastructureOperationUseCase,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		RunE: func(cmd *cobra.Command, args []string) error {
			c.bindArgs(args)

			return useCase.Process(infraCmd.Operation(), c.workspaceSettings())
		},
	}

	c.configure(cmd)

	return cmd
}

func (c *baseInfrastructureCommand) configure(cmd *cobra.Command) {
	cmd.Args = cobra.RangeArgs(1, 2)

	cmd.Flags().StringSliceVarP(
		&c.filterTokens,
		"filter",
		"f",
		nil,
		"Filter user manifests by metadata values",
	)

	cmd.Flags().BoolVar(
		&c.keepWorkspace,
		"keep-workspace",
		false,
		"Keep the workspace directory after the command finishes",
	)
}

func (c *baseInfrastructureCommand) bindArgs(args []string) {
	c.builtinDirectory = args[0]

	if len(args) > 1 {
		c.userDirectory = args[1]
	}
}

func (c *baseInfrastructureCommand) workspaceSettings() operation.WorkspaceSettings {
	return operation.NewWorkspaceSettings(c.builtinDirectory, c.userDirectory, c.keepWorkspace)
}
