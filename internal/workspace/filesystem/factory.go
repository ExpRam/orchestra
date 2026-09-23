package filesystem

import "github.com/expram/orchestra/internal/workspace"

type FileSystemWorkspaceFactory struct{}

func NewFileSystemWorkspaceFactory() FileSystemWorkspaceFactory {
	return FileSystemWorkspaceFactory{}
}

func (f FileSystemWorkspaceFactory) Create(keep bool) (workspace.Workspace, error) {
	ws, err := NewFileSystemWorkspace(!keep)
	if err != nil {
		return nil, err
	}

	return ws, nil
}
