package configtemplate

import (
	_ "embed"
	"strings"

	"github.com/nurcahyaari/kite/templates"
)

//go:embed config.tmpl
var configTemplateLoc string

type ConfigTemplateData struct {
	DatabaseDialeg string
}

type ConfigTemplate interface {
	Render() (string, error)
}

type ConfigTemplateImpl struct {
	*templates.TemplateNewImpl
	Data ConfigTemplateData
}

func NewConfigTemplate(data ConfigTemplateData) ConfigTemplate {
	template := templates.NewTemplateNewImpl("config", configTemplateLoc)
	template.AddTemplateFunction("Title", strings.Title)
	template.AddTemplateFunction("ToUpper", strings.ToUpper)

	return &ConfigTemplateImpl{
		TemplateNewImpl: template,
		Data:            data,
	}
}

func (s *ConfigTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, s.Data)
}
