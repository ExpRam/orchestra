package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/expram/orchestra/internal/cli"
	"github.com/expram/orchestra/internal/config"
	"github.com/expram/orchestra/internal/config/source"
	"github.com/spf13/cobra"
)

const (
	envFilePath = ".env"
	envPrefix   = "ORCHESTRA_"
)

const exitSuccess = 0

type App struct {
	root     *cobra.Command
	loader   *config.Loader
	logLevel *slog.LevelVar
	runtime  *runtime
}

type runtime struct {
	cfg config.Config
}

func New() *App {
	root := cli.NewRootCmd()

	app := &App{
		root: root,
		loader: config.NewLoader(
			source.NewEnvFile(envFilePath),
			source.NewEnv(envPrefix),
			source.NewFlag(root.PersistentFlags()),
		),
		logLevel: new(slog.LevelVar),
	}

	slog.SetDefault(newLogger(app.logLevel))

	root.PersistentPreRunE = func(*cobra.Command, []string) error {
		return app.initialize()
	}

	root.AddCommand(
		cli.NewPlanCommand(),
		cli.NewApplyCommand(),
		cli.NewDestroyCommand(),
	)

	return app
}

func (a *App) Run(ctx context.Context) int {
	cmd, err := a.root.ExecuteContextC(ctx)
	if err != nil {
		return cli.HandleExecutionError(cmd, err)
	}

	return exitSuccess
}

func (a *App) initialize() error {
	cfg, err := a.loadConfig()
	if err != nil {
		return err
	}

	a.logLevel.Set(logLevelFor(cfg))
	slog.Debug("debug mode enabled")

	a.runtime = newRuntime(cfg)

	return nil
}

func (a *App) loadConfig() (config.Config, error) {
	values, err := a.loader.Load()
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
	return &runtime{cfg: cfg}
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
