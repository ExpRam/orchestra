package renderer

import (
	"path/filepath"

	"github.com/expram/orchestra/internal/errscope"
	"github.com/expram/orchestra/internal/workspace"
)

type Renderer struct {
	template TemplateRenderer
}

func NewRenderer(template TemplateRenderer) Renderer {
	return Renderer{template: template}
}

func (r Renderer) Render(ws workspace.Workspace, files []string, libraries []workspace.Directory) ([]string, error) {
	var (
		rendered []string
		problems errscope.Problems
	)

	for _, file := range files {
		target, err := r.render(ws, file, libraries)
		if err != nil {
			problems.Add(err)

			continue
		}

		rendered = append(rendered, target)
	}

	if err := problems.Err(); err != nil {
		return nil, err
	}

	return rendered, nil
}

func (r Renderer) render(ws workspace.Workspace, file string, libraries []workspace.Directory) (string, error) {
	data, err := ws.ReadFile(file)
	if err != nil {
		return "", err
	}

	output, err := r.template.Render(ws, libraries, data)
	if err != nil {
		return "", errscope.In(file, Error{Problems: []error{err}})
	}

	target := filepath.Join(workspace.RENDERED, file)
	if err := ws.WriteFile(target, output); err != nil {
		return "", err
	}

	return target, nil
}
