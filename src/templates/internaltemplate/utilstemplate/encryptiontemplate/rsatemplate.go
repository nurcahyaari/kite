package encryptiontemplate

import (
	_ "embed"

	"github.com/nurcahyaari/kite/src/templates"
)

//go:embed rsa.tmpl
var rsaReaderTemplate string

type RsaTemplate interface {
	Render() (string, error)
}

type RsaTemplateImpl struct {
	*templates.TemplateNewImpl
}

func NewRsaTemplate() RsaTemplate {
	template := templates.NewTemplateNewImpl("encryption", rsaReaderTemplate)
	return &RsaTemplateImpl{
		TemplateNewImpl: template,
	}
}

func (s *RsaTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, nil)
}
