package renderer

import "github.com/expram/orchestra/internal/workspace"

type TemplateRenderer interface {
	Render(ws workspace.Workspace, libraries []workspace.Directory, data []byte) ([]byte, error)
}
