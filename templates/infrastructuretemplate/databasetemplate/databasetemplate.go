package databasetemplate

import (
	_ "embed"

	"github.com/nurcahyaari/kite/templates"
)

type databaseTemplateLoc string

//go:embed mysql.tmpl
var mysqlTemplateLoc databaseTemplateLoc

type DatabaseType int

const (
	DatabaseMysql DatabaseType = iota
)

func (s DatabaseType) GetDatabaseTemplate() string {
	template := ""
	switch s {
	case DatabaseMysql:
		template = string(mysqlTemplateLoc)
	}
	return template
}

type DatabaseTemplateData struct {
	DatabaseType DatabaseType
}

type DatabaseTemplate interface {
	Render() (string, error)
}

type DatabaseTemplateImpl struct {
	*templates.TemplateNewImpl
	Data DatabaseTemplateData
}

func NewDatabaseTemplate(data DatabaseTemplateData) DatabaseTemplate {
	template := templates.NewTemplateNewImpl("infrastructure", data.DatabaseType.GetDatabaseTemplate())
	return &DatabaseTemplateImpl{
		TemplateNewImpl: template,
		Data:            data,
	}
}

func (s *DatabaseTemplateImpl) Render() (string, error) {
	return s.TemplateNewImpl.Render(s.TemplateLocation, s.Data)
}
