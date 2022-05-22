package gracefultemplate

import (
	_ "embed"

	"github.com/nurcahyaari/kite/internal/templates"
)

//go:embed graceful.tmpl
var gracefulTemplateLoc string

type GracefulTemplateData struct {
}

type GracefulTemplate interface {
	Render() (string, error)
}

type GracefulTemplateImpl struct {
	*templates.TemplateNewImpl
}

func NewGracefulTemplate() GracefulTemplate {
	template := templates.NewTemplateNewImpl("graceful", gracefulTemplateLoc)
	return &GracefulTemplateImpl{
		TemplateNewImpl: template,
	}
}

func (s *GracefulTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, nil)
}
