package filesystem

import (
	"github.com/expram/orchestra/internal/operation/preparer"
	"github.com/expram/orchestra/internal/workspace"
)

type WorkspaceFactory struct{}

var _ preparer.WorkspaceFactory = WorkspaceFactory{}

func NewWorkspaceFactory() WorkspaceFactory {
	return WorkspaceFactory{}
}

func (f WorkspaceFactory) Create(keep bool) (workspace.Workspace, error) {
	ws, err := NewWorkspace(!keep)
	if err != nil {
		return nil, err
	}

	return ws, nil
}
