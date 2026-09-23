package cli

import (
	"github.com/spf13/cobra"

	"github.com/expram/orchestra/internal/config"
	"github.com/expram/orchestra/internal/config/source/flag"
)

func NewRootCommand(defaults config.Config) *cobra.Command {
	root := &cobra.Command{
		Use:           "orchestra",
		Short:         "Orchestrator for GitOps. Describe anything with k8s-like manifests",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.PersistentFlags().Bool(
		flag.Name(config.PathDebug),
		defaults.Debug,
		"enable debug mode",
	)

	return root
}
