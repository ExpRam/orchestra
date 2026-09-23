package filesystem

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"

	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/errscope/fserror"
	"github.com/expram/orchestra/internal/workspace"
)

type FileSystemWorkspace struct {
	root        *os.Root
	deleteAfter bool
}

var _ workspace.Workspace = (*FileSystemWorkspace)(nil)

func NewFileSystemWorkspace(deleteAfter bool) (*FileSystemWorkspace, error) {
	dir, err := os.MkdirTemp("", "workspace")
	if err != nil {
		return nil, errscope.In("create workspace", fserror.Cause(err))
	}

	root, err := os.OpenRoot(dir)
	if err != nil {
		return nil, errscope.In("create workspace", errors.Join(fserror.Cause(err), fserror.Cause(os.RemoveAll(dir))))
	}

	return &FileSystemWorkspace{
		root:        root,
		deleteAfter: deleteAfter,
	}, nil
}

func (ws *FileSystemWorkspace) Close() error {
	err := fserror.Cause(ws.root.Close())
	if ws.deleteAfter {
		err = errors.Join(err, fserror.Cause(os.RemoveAll(ws.root.Name())))
	} else {
		slog.Info("workspace kept", "path", ws.root.Name())
	}
	if err != nil {
		return errscope.In("close workspace", err)
	}
	return nil
}

func (ws *FileSystemWorkspace) Walk(dir string, fn func(path string, size int64) error) error {
	walk := func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return nested(dir, path, err)
		}
		if entry.IsDir() {
			return nil
		}

		info, err := entry.Info()
		if err != nil {
			return nested(dir, path, err)
		}

		return fn(path, info.Size())
	}

	if err := fs.WalkDir(ws.root.FS(), dir, walk); err != nil {
		return errscope.In(fmt.Sprintf("walk %q", dir), err)
	}

	return nil
}

func (ws *FileSystemWorkspace) ReadFile(path string) ([]byte, error) {
	data, err := ws.root.ReadFile(path)
	if err != nil {
		return nil, errscope.In(fmt.Sprintf("read %q", path), fserror.Cause(err))
	}
	return data, nil
}

func (ws *FileSystemWorkspace) WriteFile(path string, data []byte) error {
	if err := ws.writeFile(path, data); err != nil {
		return errscope.In(fmt.Sprintf("write %q", path), err)
	}
	return nil
}

func (ws *FileSystemWorkspace) writeFile(path string, data []byte) error {
	path, err := localPath(path)
	if err != nil {
		return err
	}
	if err := ws.Mkdir(filepath.Dir(path)); err != nil {
		return err
	}
	return fserror.Cause(ws.root.WriteFile(path, data, 0o644))
}

func (ws *FileSystemWorkspace) Mkdir(path string) error {
	if err := ws.root.MkdirAll(path, 0o755); err != nil {
		return errscope.In(fmt.Sprintf("create directory %q", path), fserror.Cause(err))
	}
	return nil
}

func (ws *FileSystemWorkspace) Remove(path string) error {
	if err := ws.remove(path); err != nil {
		return errscope.In(fmt.Sprintf("remove %q", path), err)
	}
	return nil
}

func (ws *FileSystemWorkspace) remove(path string) error {
	path, err := localPath(path)
	if err != nil {
		return err
	}
	if path == "." {
		return errors.New("cannot remove workspace root")
	}
	return fserror.Cause(ws.root.RemoveAll(path))
}

func (ws *FileSystemWorkspace) Copy(src, dst string) error {
	if err := ws.copy(src, dst); err != nil {
		return errscope.In(fmt.Sprintf("copy %q to %q", src, dst), err)
	}
	return nil
}

func (ws *FileSystemWorkspace) copy(src, dst string) error {
	dst, err := localPath(dst)
	if err != nil {
		return err
	}
	if err := ws.rejectSymlinks(dst); err != nil {
		return err
	}

	info, err := os.Lstat(src)
	if err != nil {
		return fserror.Cause(err)
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return fmt.Errorf("symbolic links are not supported: %q", src)
	}

	if info.Mode().IsRegular() {
		return ws.copyFile(src, dst, info)
	}
	if !info.IsDir() {
		return fmt.Errorf("unsupported file type: %q", src)
	}

	return filepath.WalkDir(src, func(path string, entry fs.DirEntry, walkErr error) error {
		if err := ws.copyEntry(src, dst, path, entry, walkErr); err != nil {
			return nested(src, path, err)
		}
		return nil
	})
}

func (ws *FileSystemWorkspace) copyEntry(src, dst, path string, entry fs.DirEntry, walkErr error) error {
	if walkErr != nil {
		return walkErr
	}

	rel, err := filepath.Rel(src, path)
	if err != nil {
		return err
	}
	target := filepath.Join(dst, rel)

	entryInfo, err := entry.Info()
	if err != nil {
		return err
	}

	switch {
	case entryInfo.Mode()&os.ModeSymlink != 0:
		return fmt.Errorf("symbolic links are not supported: %q", path)
	case entryInfo.IsDir():
		return fserror.Cause(ws.root.MkdirAll(target, entryInfo.Mode().Perm()))
	case entryInfo.Mode().IsRegular():
		return ws.copyFile(path, target, entryInfo)
	default:
		return fmt.Errorf("unsupported file type: %q", path)
	}
}

func (ws *FileSystemWorkspace) copyFile(src, dst string, info fs.FileInfo) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return fserror.Cause(err)
	}
	defer in.Close()

	if err := ws.root.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return fserror.Cause(err)
	}

	out, err := ws.root.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode().Perm())
	if err != nil {
		return fserror.Cause(err)
	}
	defer func() {
		err = errors.Join(err, fserror.Cause(out.Close()))
	}()

	_, err = io.Copy(out, in)
	return fserror.Cause(err)
}

func localPath(path string) (string, error) {
	if !filepath.IsLocal(path) {
		return "", fmt.Errorf("invalid workspace path: %q", path)
	}
	return filepath.Clean(path), nil
}

func (ws *FileSystemWorkspace) rejectSymlinks(path string) error {
	for dir := path; dir != "."; dir = filepath.Dir(dir) {
		info, err := ws.root.Lstat(dir)
		if errors.Is(err, fs.ErrNotExist) {
			continue
		}
		if err != nil {
			return errscope.In(strconv.Quote(dir), fserror.Cause(err))
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("symbolic links are not supported: %q", dir)
		}
	}
	return nil
}

func nested(root, path string, err error) error {
	err = fserror.Cause(err)
	if path == root {
		return err
	}
	return errscope.In(strconv.Quote(path), err)
}
