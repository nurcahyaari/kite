package chitemplate

import (
	_ "embed"

	"github.com/nurcahyaari/kite/internal/templates"
)

//go:embed internalhttprouter.tmpl
var internalHttpRouterTemplateLoc string

type InternalHttpRouterTemplate interface {
	Render() (string, error)
}

type InternalHttpRouterTemplateImpl struct {
	*templates.TemplateNewImpl
}

func NewInternalHttpRouterTemplate() InternalHttpTemplate {
	template := templates.NewTemplateNewImpl("router", internalHttpRouterTemplateLoc)
	return &InternalHttpTemplateImpl{
		TemplateNewImpl: template,
	}
}

func (s *InternalHttpRouterTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, nil)
}
