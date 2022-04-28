package loggertemplate

import (
	_ "embed"

	"github.com/nurcahyaari/kite/internal/templates"
)

//go:embed logger.tmpl
var loggerTemplateLoc string

type LoggerTemplateData struct {
}

type LoggerTemplate interface {
	Render() (string, error)
}

type LoggerTemplateImpl struct {
	*templates.TemplateNewImpl
}

func NewLoggerTemplate() LoggerTemplate {
	template := templates.NewTemplateNewImpl("logger", loggerTemplateLoc)
	return &LoggerTemplateImpl{
		TemplateNewImpl: template,
	}
}

func (s *LoggerTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, nil)
}
