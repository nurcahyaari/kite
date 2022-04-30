package envgen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates/configtemplate"
)

type EnvGen interface {
	CreateEnvExampleFile(path string) error
	CreateEnvFile(path string) error
	AddConfigToEnv(option EnvOption) error
}

type EnvGenImpl struct {
	fs database.FileSystem
}

func NewEnvGen(fs database.FileSystem) *EnvGenImpl {
	return &EnvGenImpl{
		fs: fs,
	}
}

func (s *EnvGenImpl) CreateEnvExampleFile(path string) error {
	templateNew := configtemplate.NewEnvTemplate(configtemplate.EnvTemplateData{
		DatabaseDialeg: "mysql",
	})
	envTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	return s.fs.CreateFileIfNotExists(path, ".env.example", envTemplate)
}

func (s *EnvGenImpl) CreateEnvFile(path string) error {
	templateNew := configtemplate.NewEnvTemplate(configtemplate.EnvTemplateData{
		DatabaseDialeg: "mysql",
	})
	envTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	return s.fs.CreateFileIfNotExists(path, ".env", envTemplate)
}

func (s *EnvGenImpl) AddConfigToEnv(option EnvOption) error {
	return nil
}
