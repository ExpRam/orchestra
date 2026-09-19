package cli

import (
	"fmt"

	"github.com/expram/orchestra/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "orchestra",
	Short:         "Orchestrator for GitOps. Describe anything with k8s-like manifests",
	SilenceUsage:  true,
	SilenceErrors: true,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		configLoader := config.NewLoader(
			config.NewEnvFile(".env"),
			config.NewEnv(),
			config.NewCommand(cmd.Flags()),
		)

		values, err := configLoader.Load()
		if err != nil {
			return err
		}

		fmt.Println(values["DEBUG"])

		_ = values

		return nil
	},
}

func Execute() int {
	Init()
	cmd, err := rootCmd.ExecuteC()
	if err != nil {
		return handleExecutionError(cmd, err)
	}

	return 0
}

func Init() {
	rootCmd.PersistentFlags().Bool(
		"debug",
		false,
		"enable debug mode",
	)
}
