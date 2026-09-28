package pongo2

import (
	pongo2v6 "github.com/flosch/pongo2/v6"

	"github.com/expram/orchestra/internal/operation/renderer"
	"github.com/expram/orchestra/internal/workspace"
)

const setName = "orchestra"

var bannedTags = []string{"ssi", "now"}

type TemplateRenderer struct{}

var _ renderer.TemplateRenderer = TemplateRenderer{}

func NewTemplateRenderer() TemplateRenderer {
	pongo2v6.SetAutoescape(false)

	return TemplateRenderer{}
}

func (r TemplateRenderer) Render(ws workspace.Workspace, libraries []workspace.Directory, data []byte) ([]byte, error) {
	loader := newLoader(ws, libraries)

	set := pongo2v6.NewSet(setName, loader)
	for _, tag := range bannedTags {
		if err := set.BanTag(tag); err != nil {
			return nil, err
		}
	}

	template, err := set.FromBytes(data)
	if err != nil {
		return nil, translate(err, loader)
	}

	rendered, err := template.ExecuteBytes(nil)
	if err != nil {
		return nil, translate(err, loader)
	}

	return rendered, nil
}
