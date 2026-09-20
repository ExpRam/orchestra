package cli

import (
	"errors"
	"log/slog"

	"github.com/spf13/cobra"
)

type exitCoder interface {
	ExitCode() int
}

func HandleExecutionError(cmd *cobra.Command, err error) int {
	slog.Error(
		"command execution failed",
		"command", cmd.CommandPath(),
		"error", err,
	)

	var coder exitCoder
	if errors.As(err, &coder) {
		return coder.ExitCode()
	}

	return 1
}
