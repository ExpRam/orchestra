package cli

import (
	"github.com/expram/orchestra/internal/config"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "orchestra",
		Short:         "Orchestrator for GitOps. Describe anything with k8s-like manifests",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.PersistentFlags().Bool(
		config.KeyDebug,
		false,
		"enable debug mode",
	)

	return root
}
