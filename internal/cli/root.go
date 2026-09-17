package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "orchestra",
	Short:         "Orchestrator for GitOps. Describe anything with k8s-like manifests",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() int {
	cmd, err := rootCmd.ExecuteC()
	if err != nil {
		return handleExecutionError(cmd, err)
	}

	return 0
}
