package operation

import "github.com/expram/orchestra/internal/workspace"

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
	wsUseCase workspace.PrepareWorkspaceUseCase
}

func (c *callInfrastructureOperationUseCase) Process(op InfrastructureOperation, wsSettings WorkspaceSettings) error {
	ws, err := c.wsUseCase.Prepare(wsSettings.builtinDirectory, wsSettings.userDirectory)
	if err != nil {
		return err
	}

	_, _ = op, ws

	return nil
}

func NewCallInfrastructureOperationUseCase(
	wsUseCase workspace.PrepareWorkspaceUseCase,
) CallInfrastructureOperationUseCase {
	return &callInfrastructureOperationUseCase{wsUseCase: wsUseCase}
}
