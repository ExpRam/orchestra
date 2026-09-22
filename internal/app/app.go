package app

import (
	"context"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/expram/orchestra/internal/cli"
	"github.com/expram/orchestra/internal/config"
	"github.com/expram/orchestra/internal/config/loader"
	"github.com/expram/orchestra/internal/config/source/env"
	"github.com/expram/orchestra/internal/config/source/flag"
	"github.com/expram/orchestra/internal/operation"
	"github.com/expram/orchestra/internal/provider"
	"github.com/expram/orchestra/internal/workspace"
)

const (
	envFilePath = ".env"
	envPrefix   = "ORCHESTRA_"
)

const exitSuccess = 0

type App struct {
	root    *cobra.Command
	runtime *Runtime
	static *Static
}

type Static struct {
	opUseCase operation.CallInfrastructureOperationUseCase
}

type Runtime struct {
	cfg provider.Provider[config.Config]
}

func New() *App {
	logLevel := new(slog.LevelVar)
	slog.SetDefault(newLogger(logLevel))

	static := newStatic()

	cfg := config.Defaults()
	root := cli.NewRootCmd(cfg)

	root.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		loaded, err := loadConfig(cmd)
		if err != nil {
			slog.Warn("load configuration, falling back to defaults", "error", err)

			loaded = config.Defaults()
		}

		cfg = loaded

		logLevel.Set(logLevelFor(cfg))
		slog.Debug("debug mode enabled")

		return nil
	}

	runtime := newRuntime(provider.New(func() config.Config { return cfg }))

	root.AddCommand(
		cli.NewPlanCommand(static.opUseCase),
		cli.NewApplyCommand(static.opUseCase),
		cli.NewDestroyCommand(static.opUseCase),
	)

	return &App{root: root, runtime: runtime, static: static}
}

func (a *App) Run(ctx context.Context) int {
	cmd, err := a.root.ExecuteContextC(ctx)
	if err != nil {
		return cli.HandleExecutionError(cmd, err)
	}

	return exitSuccess
}

func loadConfig(cmd *cobra.Command) (config.Config, error) {
	return loader.New(
		env.NewFile(envFilePath, envPrefix),
		env.NewOS(envPrefix),
		flag.New(cmd.Flags()),
	).Load()
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

func newRuntime(cfg provider.Provider[config.Config]) *Runtime {
	return &Runtime{cfg: cfg}
}

func newStatic() *Static {
	wsUseCase := workspace.NewPrepareWorkspaceUseCase()

	opUseCase := operation.NewCallInfrastructureOperationUseCase(
		wsUseCase,
	)

	return &Static{opUseCase}
}
