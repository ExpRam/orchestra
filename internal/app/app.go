package app

import (
	"context"
	"log/slog"
	"os"

	"github.com/expram/orchestra/internal/manifest/processor"
	"github.com/spf13/cobra"

	"github.com/expram/orchestra/internal/cli"
	"github.com/expram/orchestra/internal/config"
	"github.com/expram/orchestra/internal/config/loader"
	"github.com/expram/orchestra/internal/config/source/env"
	"github.com/expram/orchestra/internal/config/source/flag"
	"github.com/expram/orchestra/internal/manifest/collector"
	"github.com/expram/orchestra/internal/manifest/dto"
	"github.com/expram/orchestra/internal/manifest/kind/ord"
	ordprocessor "github.com/expram/orchestra/internal/manifest/kind/ord/processor"
	ordv1 "github.com/expram/orchestra/internal/manifest/kind/ord/v1"
	"github.com/expram/orchestra/internal/manifest/reader"
	"github.com/expram/orchestra/internal/manifest/registry"
	"github.com/expram/orchestra/internal/manifest/renderer"
	"github.com/expram/orchestra/internal/manifest/resolver"
	"github.com/expram/orchestra/internal/manifest/yaml"
	"github.com/expram/orchestra/internal/mapstructure"
	"github.com/expram/orchestra/internal/operation"
	"github.com/expram/orchestra/internal/pongo2"
	"github.com/expram/orchestra/internal/provider"
	"github.com/expram/orchestra/internal/workspace/filesystem"
	"github.com/expram/orchestra/internal/workspace/preparer"
)

const (
	envFilePath = ".env"
	envPrefix   = "ORCHESTRA_"
)

const exitSuccess = 0

type App struct {
	root    *cobra.Command
	runtime Runtime
	static  Static
}

type Static struct {
	opUseCase operation.CallInfrastructureOperationUseCase
}

type Runtime struct {
	cfg provider.Provider[config.Config]
}

func NewApp() App {
	logLevel := new(slog.LevelVar)
	slog.SetDefault(newLogger(logLevel))

	static := newStatic()

	cfg := config.Defaults()
	root := cli.NewRootCommand(cfg)

	root.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		loaded, err := loadConfig(cmd)
		if err != nil {
			slog.Warn("load configuration, falling back to defaults", cli.ErrorAttr(err))

			loaded = config.Defaults()
		}

		cfg = loaded

		logLevel.Set(logLevelFor(cfg))
		slog.Debug("debug mode enabled")

		return nil
	}

	runtime := newRuntime(provider.NewProvider(func() config.Config { return cfg }))

	root.AddCommand(
		cli.NewPlanCommand(static.opUseCase),
		cli.NewApplyCommand(static.opUseCase),
		cli.NewDestroyCommand(static.opUseCase),
	)

	return App{root: root, runtime: runtime, static: static}
}

func (a App) Run(ctx context.Context) int {
	cmd, err := a.root.ExecuteContextC(ctx)
	if err != nil {
		return cli.HandleExecutionError(cmd, err)
	}

	return exitSuccess
}

func loadConfig(cmd *cobra.Command) (config.Config, error) {
	return loader.NewLoader(
		env.NewFile(envFilePath, envPrefix),
		env.NewOS(envPrefix),
		flag.NewFlag(cmd.Flags()),
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

func newRuntime(cfg provider.Provider[config.Config]) Runtime {
	return Runtime{cfg: cfg}
}

func newStatic() Static {
	workspaceFactory := filesystem.NewFileSystemWorkspaceFactory()
	operationPreparer := preparer.NewPreparer(workspaceFactory)

	operationCollector := collector.NewCollector()

	templateRenderer := pongo2.NewPongo2TemplateRenderer()
	operationRenderer := renderer.NewRenderer(templateRenderer)

	manifestValidator := dto.NewKRMUnresolvedManifestValidator()
	manifestParser := yaml.NewYamlUnresolvedManifestParser(manifestValidator)
	operationReader := reader.NewReader(manifestParser)

	specDecoder := mapstructure.NewMapstructureSpecDecoder(resolver.SpecTag)
	specRegistry := registry.NewRegistry(
		resolver.RegisterSpec[ordv1.Spec](ordv1.Type),
	)
	operationResolver := resolver.NewResolver(specDecoder, specRegistry)

	processorRegistry := registry.NewRegistry(
		processor.Register[ord.Spec](ord.Kind, ordprocessor.NewOrdProcessor()),
	)

	operationProcessor := processor.NewProcessor(processorRegistry)

	opUseCase := operation.NewCallInfrastructureOperationUseCase(
		operationPreparer,
		operationCollector,
		operationRenderer,
		operationReader,
		operationResolver,
		operationProcessor,
	)

	return Static{opUseCase}
}
