package chitemplate

import (
	_ "embed"

	"github.com/nurcahyaari/kite/internal/templates"
)

//go:embed middleware.tmpl
var middlewareTemplateLoc string

type MiddlewareTemplate interface {
	Render() (string, error)
}

type MiddlewareTemplateImpl struct {
	*templates.TemplateNewImpl
}

func NewMiddlewareTemplate() MiddlewareTemplate {
	template := templates.NewTemplateNewImpl("middleware", middlewareTemplateLoc)
	return &MiddlewareTemplateImpl{
		TemplateNewImpl: template,
	}
}

func (s *MiddlewareTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, nil)
}
