package misctemplate

import (
	_ "embed"

	"github.com/nurcahyaari/kite/internal/templates"
)

//go:embed gitignore.tmpl
var gitignoreTemplateLoc string

type GitignoreTemplate interface {
	Render() (string, error)
}

type GitignoreTemplateImpl struct {
	*templates.TemplateNewImpl
}

func NewGitignoreTemplate() GitignoreTemplate {
	template := templates.NewTemplateNewImpl("", gitignoreTemplateLoc)
	return &GitignoreTemplateImpl{
		TemplateNewImpl: template,
	}
}

func (s *GitignoreTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, nil)
}
