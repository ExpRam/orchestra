package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/expram/orchestra/internal/cli"
	"github.com/expram/orchestra/internal/config"
	"github.com/expram/orchestra/internal/config/source"
	"github.com/expram/orchestra/internal/workspace"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	envFilePath = ".env"
	envPrefix   = "ORCHESTRA_"
)

const exitSuccess = 0

type App struct {
	root    *cobra.Command
}

type runtime struct {
	cfg config.Config

	prepareWorkspace workspace.PrepareWorkspaceUseCase
}

func New() *App {
	root := cli.NewRootCmd()

	logLevel := new(slog.LevelVar)
	slog.SetDefault(newLogger(logLevel))

	cfg, err := loadConfig(root.PersistentFlags())
	if err != nil {
		slog.Warn("load configuration, falling back to defaults", "error", err)
	}

	logLevel.Set(logLevelFor(cfg))
	slog.Debug("debug mode enabled")

	// runtime := newRuntime(cfg)

	root.AddCommand(
		cli.NewPlanCommand(),
		cli.NewApplyCommand(),
		cli.NewDestroyCommand(),
	)

	return &App{
		root:    root,
	}
}

func (a *App) Run(ctx context.Context) int {
	cmd, err := a.root.ExecuteContextC(ctx)
	if err != nil {
		return cli.HandleExecutionError(cmd, err)
	}

	return exitSuccess
}

func loadConfig(flags *pflag.FlagSet) (config.Config, error) {
	loader := config.NewLoader(
		source.NewEnvFile(envFilePath),
		source.NewEnv(envPrefix),
		source.NewFlag(flags, os.Args[1:]),
	)

	values, err := loader.Load()
	if err != nil {
		return config.Config{}, fmt.Errorf("load configuration: %w", err)
	}

	cfg, err := config.ToConfig(values)
	if err != nil {
		return config.Config{}, fmt.Errorf("decode configuration: %w", err)
	}

	return cfg, nil
}

func newRuntime(cfg config.Config) *runtime {
	prepareWorkspace := workspace.NewPrepareWorkspaceUseCase()
	return &runtime{cfg: cfg, prepareWorkspace: prepareWorkspace}
}

func logLevelFor(cfg config.Config) slog.Level {
	if cfg.Debug {
		return slog.LevelDebug
	}

	return slog.LevelInfo
}

func newLogger(level slog.Leveler) *slog.Logger {
	return slog.New(
		slog.NewTextHandler(
			os.Stderr,
			&slog.HandlerOptions{
				Level: level,
			},
		),
	)
}
