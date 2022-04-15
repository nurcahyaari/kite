package config

import (
	"github.com/nurcahyaari/kite/src/templates/configtemplate"
	"github.com/nurcahyaari/kite/src/utils/fs"
)

type EnvGen interface {
	CreateEnvExampleFile() error
	CreateEnvFile() error
	AddConfigToEnv() error
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
	templateNew := configtemplate.NewEnvTemplate(configtemplate.EnvTemplateData{
		DatabaseDialeg: "mysql",
	})
	envTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	return fs.CreateFileIfNotExist(s.ProjectPath, ".env.example", envTemplate)
}

func (s *EnvGenImpl) CreateEnvFile() error {
	templateNew := configtemplate.NewEnvTemplate(configtemplate.EnvTemplateData{
		DatabaseDialeg: "mysql",
	})
	envTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	return fs.CreateFileIfNotExist(s.ProjectPath, ".env", envTemplate)
}

func (s *EnvGenImpl) AddConfigToEnv() error {
	return nil
}
