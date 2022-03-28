package config

import (
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type EnvGen interface {
	CreateEnvExampleFile() error
	CreateEnvFile() error
}

type EnvGenImpl struct {
	ProjectPath string
}

func NewEnvGen(projectPath string) EnvGen {
	return &EnvGenImpl{
		ProjectPath: projectPath,
	}
}

func (s *EnvGenImpl) CreateEnvExampleFile() error {
	tmpl := templates.NewTemplate(templates.Template{
		Template: templates.EnvTemplate,
		Data: map[string]interface{}{
			"DBDialeg": "mysql",
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	return fs.CreateFileIfNotExist(s.ProjectPath, ".env.example", templateString)
}

func (s *EnvGenImpl) CreateEnvFile() error {
	tmpl := templates.NewTemplate(templates.Template{
		Template: templates.EnvTemplate,
		Data: map[string]interface{}{
			"DBDialeg": "mysql",
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	return fs.CreateFileIfNotExist(s.ProjectPath, ".env", templateString)
}
