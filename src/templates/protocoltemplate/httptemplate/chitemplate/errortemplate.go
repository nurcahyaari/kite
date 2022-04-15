package chitemplate

import (
	_ "embed"

	"github.com/nurcahyaari/kite/src/templates"
)

//go:embed err.tmpl
var errorTemplateLoc string

type ErrorTemplate interface {
	Render() (string, error)
}

type ErrorTemplateImpl struct {
	*templates.TemplateNewImpl
}

func NewErrorTemplate() ErrorTemplate {
	template := templates.NewTemplateNewImpl("error", errorTemplateLoc)
	return &ErrorTemplateImpl{
		TemplateNewImpl: template,
	}
}

func (s *ErrorTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, nil)
}
