package filesystem

import (
	"errors"
	"fmt"

	"github.com/expram/orchestra/internal/workspace"
)

type FileSystemPreparer struct{}

func NewFileSystemPreparer() FileSystemPreparer {
	return FileSystemPreparer{}
}

func (p FileSystemPreparer) Prepare(builtin, user string) (workspace.Workspace, error) {
	w, err := NewFileSystemWorkspace(true)
	if err != nil {
		return nil, fmt.Errorf("prepare workspace: %w", err)
	}

	if err := p.populate(w, workspace.BUILTIN, builtin); err != nil {
		return nil, closeAndWrap(w, err)
	}
	if err := p.populate(w, workspace.USER, user); err != nil {
		return nil, closeAndWrap(w, err)
	}

	return w, nil
}

func (p FileSystemPreparer) populate(w workspace.Workspace, dir workspace.Directory, src string) error {
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
