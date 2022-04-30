package databasetemplate

import "github.com/nurcahyaari/kite/internal/templates"

//go:embed mysql.tmpl
var mysqlTemplateLoc string

type GitignoreTemplate interface {
	Render() (string, error)
}

type GitignoreTemplateImpl struct {
	*templates.TemplateNewImpl
}

func NewMysqlTemplate() GitignoreTemplate {
	template := templates.NewTemplateNewImpl("database", mysqlTemplateLoc)
	return &GitignoreTemplateImpl{
		TemplateNewImpl: template,
	}
}

func (s *GitignoreTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, nil)
}
