package preparer

import "github.com/expram/orchestra/internal/workspace"

type WorkspaceFactory interface {
	Create(keep bool) (workspace.Workspace, error)
}
