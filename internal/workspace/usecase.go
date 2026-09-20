package workspace

import (
	"errors"
	"fmt"
)

type PrepareWorkspaceUseCase interface {
	Prepare(builtin, user string) (Workspace, error)
}

type prepareWorkspaceUseCase struct{}

func NewPrepareWorkspaceUseCase() PrepareWorkspaceUseCase {
	return prepareWorkspaceUseCase{}
}

func (p prepareWorkspaceUseCase) Prepare(builtin, user string) (Workspace, error) {
	w, err := NewFileSystemWorkspace(true)
	if err != nil {
		return nil, fmt.Errorf("prepare workspace: %w", err)
	}

	if err := p.populate(w, BUILTIN, builtin); err != nil {
		return nil, closeAndWrap(w, err)
	}
	if err := p.populate(w, USER, user); err != nil {
		return nil, closeAndWrap(w, err)
	}

	return w, nil
}

func (p prepareWorkspaceUseCase) populate(w Workspace, dir Directory, src string) error {
	if err := w.Mkdir(dir); err != nil {
		return fmt.Errorf("create %s directory: %w", dir, err)
	}
	if err := w.Copy(src, dir); err != nil {
		return fmt.Errorf("populate %s directory: %w", dir, err)
	}
	return nil
}

func closeAndWrap(w *FileSystemWorkspace, err error) error {
	if closeErr := w.Close(); closeErr != nil {
		err = errors.Join(err, closeErr)
	}
	return fmt.Errorf("prepare workspace: %w", err)
}
