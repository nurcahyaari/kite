package config

import (
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type ConfigGen interface {
	CreateConfigDir() error
	CreateConfigFile() error
}

type ConfigGenImpl struct {
	AppName    string
	ConfigPath string
}

func NewConfig(projectPath string,
	appName string,
) ConfigGen {
	configPath := fs.ConcatDirPath(projectPath, "config")
	return &ConfigGenImpl{
		ConfigPath: configPath,
		AppName:    appName,
	}
}

func (s *ConfigGenImpl) CreateConfigDir() error {
	err := fs.CreateFolderIsNotExist(s.ConfigPath)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConfigGenImpl) CreateConfigFile() error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "config",
		Template:    templates.ConfigTemplate,
		Import: []templates.ImportedPackage{
			{
				FilePath: "log",
			},
			{
				FilePath: "sync",
			},
			{
				FilePath: "github.com/spf13/viper",
			},
		},
		Data: map[string]interface{}{
			"AppName":  s.AppName,
			"DBDialeg": "mysql",
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	return fs.CreateFileIfNotExist(s.ConfigPath, "config.go", templateString)
}
