package chitemplate

import (
	_ "embed"

	"github.com/nurcahyaari/kite/templates"
)

//go:embed internalhttptemplate.tmpl
var internalHttpTemplateLoc string

type InternalHttpTemplate interface {
	Render() (string, error)
}

type InternalHttpTemplateImpl struct {
	*templates.TemplateNewImpl
}

func NewInternalHttpTemplate() InternalHttpTemplate {
	template := templates.NewTemplateNewImpl("http", internalHttpTemplateLoc)
	return &InternalHttpTemplateImpl{
		TemplateNewImpl: template,
	}
}

func (s *InternalHttpTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, nil)
}
