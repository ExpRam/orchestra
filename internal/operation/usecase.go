package operation

import (
	"github.com/expram/orchestra/internal/workspace"
)

type WorkspaceSettings struct {
	builtinDirectory string
	userDirectory    string
}

func NewWorkspaceSettings(builtinDirectory, userDirectory string) WorkspaceSettings {
	return WorkspaceSettings{
		builtinDirectory: builtinDirectory,
		userDirectory:    userDirectory,
	}
}

type CallInfrastructureOperationUseCase interface {
	Process(op InfrastructureOperation, wsSettings WorkspaceSettings) error
}

type callInfrastructureOperationUseCase struct {
	preparer  WorkspacePreparer
	collector ManifestFileCollector
	reader    UnresolvedManifestReader
}

func (c *callInfrastructureOperationUseCase) Process(op InfrastructureOperation, wsSettings WorkspaceSettings) error {
	ws, err := c.preparer.Prepare(wsSettings.builtinDirectory, wsSettings.userDirectory)
	if err != nil {
		return err
	}

	builtinFiles, err := c.collector.Collect(ws, workspace.BUILTIN)
	if err != nil {
		return err
	}

	unresolvedManifests, err := c.reader.Read(ws, builtinFiles)
	if err != nil {
		return err
	}

	_, _ = op, unresolvedManifests

	return nil
}

func NewCallInfrastructureOperationUseCase(
	preparer WorkspacePreparer,
	collector ManifestFileCollector,
	reader UnresolvedManifestReader,
) CallInfrastructureOperationUseCase {
	return &callInfrastructureOperationUseCase{
		preparer:  preparer,
		collector: collector,
		reader:    reader,
	}
}
