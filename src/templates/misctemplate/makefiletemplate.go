package misctemplate

import (
	_ "embed"

	"github.com/nurcahyaari/kite/src/templates"
)

//go:embed makefile.tmpl
var makefileTemplateLoc string

type MakefileTemplate interface {
	Render() (string, error)
}

type MakefileTemplateImpl struct {
	*templates.TemplateNewImpl
}

func NewMakefileTemplate() MakefileTemplate {
	template := templates.NewTemplateNewImpl("", makefileTemplateLoc)
	return &MakefileTemplateImpl{
		TemplateNewImpl: template,
	}
}

func (s *MakefileTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, nil)
}
