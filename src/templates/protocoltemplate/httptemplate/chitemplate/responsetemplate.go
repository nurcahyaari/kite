package chitemplate

import (
	_ "embed"

	"github.com/nurcahyaari/kite/templates"
)

//go:embed response.tmpl
var responseTemplateLoc string

type ResponseTemplate interface {
	Render() (string, error)
}

type ResponseTemplateImpl struct {
	*templates.TemplateNewImpl
}

func NewResponseTemplate() ResponseTemplate {
	template := templates.NewTemplateNewImpl("Response", responseTemplateLoc)
	return &ResponseTemplateImpl{
		TemplateNewImpl: template,
	}
}

func (s *ResponseTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, nil)
}
