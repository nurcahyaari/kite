package configtemplate

import (
	_ "embed"
	"strings"

	"github.com/nurcahyaari/kite/src/templates"
)

//go:embed env.tmpl
var envTemplateLoc string

type EnvTemplateData struct {
	DatabaseDialeg string
}

type EnvTemplate interface {
	Render() (string, error)
}

type EnvTemplateImpl struct {
	*templates.TemplateNewImpl
	Data EnvTemplateData
}

func NewEnvTemplate(data EnvTemplateData) EnvTemplate {
	template := templates.NewTemplateNewImpl("", envTemplateLoc)
	template.AddTemplateFunction("Title", strings.Title)
	template.AddTemplateFunction("ToUpper", strings.ToUpper)

	return &EnvTemplateImpl{
		TemplateNewImpl: template,
		Data:            data,
	}
}

func (s *EnvTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, s.Data)
}
