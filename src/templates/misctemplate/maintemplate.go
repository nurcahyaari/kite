package misctemplate

import (
	_ "embed"

	"github.com/nurcahyaari/kite/src/templates"
)

//go:embed main.tmpl
var mainTemplateLoc string

type MainTemplate interface {
	Render() (string, error)
}

type MainTemplateImpl struct {
	*templates.TemplateNewImpl
}

func NewMainTemplate() MainTemplate {
	template := templates.NewTemplateNewImpl("main", mainTemplateLoc)
	return &MainTemplateImpl{
		TemplateNewImpl: template,
	}
}

func (s *MainTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, nil)
}
