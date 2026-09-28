package preparer

import (
	"errors"

	"github.com/expram/orchestra/internal/operation"
	"github.com/expram/orchestra/internal/workspace"
)

type Preparer struct {
	factory WorkspaceFactory
}

var _ operation.WorkspacePreparer = Preparer{}

func NewPreparer(factory WorkspaceFactory) Preparer {
	return Preparer{factory: factory}
}

func (p Preparer) Prepare(builtin, user string, keep bool) (workspace.Workspace, error) {
	w, err := p.factory.Create(keep)
	if err != nil {
		return nil, err
	}

	if err := p.populate(w, workspace.Builtin, builtin); err != nil {
		return nil, closeAndJoin(w, err)
	}

	if user != "" {
		if err := p.populate(w, workspace.User, user); err != nil {
			return nil, closeAndJoin(w, err)
		}
	}

	return w, nil
}

func (p Preparer) populate(w workspace.Workspace, dir workspace.Directory, src string) error {
	if err := w.Mkdir(dir); err != nil {
		return err
	}

	return w.Copy(src, dir)
}

func closeAndJoin(w workspace.Workspace, err error) error {
	return errors.Join(err, w.Close())
}
