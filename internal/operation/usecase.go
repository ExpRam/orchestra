package operation

import (
	"errors"
	"fmt"

	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/manifest"
	"github.com/expram/orchestra/internal/workspace"
)

type WorkspaceSettings struct {
	builtinDirectory string
	userDirectory    string
	keep             bool
}

func NewWorkspaceSettings(builtinDirectory, userDirectory string, keep bool) WorkspaceSettings {
	return WorkspaceSettings{
		builtinDirectory: builtinDirectory,
		userDirectory:    userDirectory,
		keep:             keep,
	}
}

type CallInfrastructureOperationUseCase interface {
	Process(op InfrastructureOperation, wsSettings WorkspaceSettings) error
}

type callInfrastructureOperationUseCase struct {
	preparer   WorkspacePreparer
	collector  ManifestFileCollector
	renderer   ManifestRenderer
	reader     UnresolvedManifestReader
	resolver   ManifestResolver
	processors ManifestProcessors
}

func (c callInfrastructureOperationUseCase) Process(op InfrastructureOperation, wsSettings WorkspaceSettings) (err error) {
	ws, err := c.preparer.Prepare(wsSettings.builtinDirectory, wsSettings.userDirectory, wsSettings.keep)
	if err != nil {
		return errscope.In("prepare workspace", err)
	}
	defer func() {
		err = errors.Join(err, ws.Close())
	}()

	builtinFiles, err := c.collector.Collect(ws, workspace.BUILTIN)
	if err != nil {
		return errscope.In(fmt.Sprintf("collect %s manifest files", workspace.BUILTIN), err)
	}

	renderedFiles, err := c.renderer.Render(ws, builtinFiles, []workspace.Directory{workspace.BUILTIN})
	if err != nil {
		return errscope.In(fmt.Sprintf("render %s manifests", workspace.BUILTIN), err)
	}

	unresolvedManifests, err := c.reader.Read(ws, renderedFiles)
	if err != nil {
		return errscope.In(fmt.Sprintf("read %s manifests", workspace.BUILTIN), err)
	}

	manifests, err := c.resolver.Resolve(unresolvedManifests)
	if err != nil {
		return errscope.In(fmt.Sprintf("resolve %s manifests", workspace.BUILTIN), err)
	}

	if err := manifest.EnsureUnique(manifests); err != nil {
		return errscope.In(fmt.Sprintf("check %s manifests", workspace.BUILTIN), err)
	}

	if err := c.processors.Process(op, manifests, ws); err != nil {
		return errscope.In(fmt.Sprintf("process %s manifests", workspace.BUILTIN), err)
	}

	return nil
}

func NewCallInfrastructureOperationUseCase(
	preparer WorkspacePreparer,
	collector ManifestFileCollector,
	renderer ManifestRenderer,
	reader UnresolvedManifestReader,
	resolver ManifestResolver,
	processors ManifestProcessors,
) CallInfrastructureOperationUseCase {
	return callInfrastructureOperationUseCase{
		preparer:   preparer,
		collector:  collector,
		renderer:   renderer,
		reader:     reader,
		resolver:   resolver,
		processors: processors,
	}
}
