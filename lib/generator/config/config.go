package config

import (
	"go/parser"

	"github.com/nurcahyaari/kite/lib/ast"
	"github.com/nurcahyaari/kite/templates/configtemplate"
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
	templateNew := configtemplate.NewConfigTemplate(configtemplate.ConfigTemplateData{
		DatabaseDialeg: "mysql",
	})
	configTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	configAbstractCode := ast.NewAbstractCode(configTemplate, parser.ParseComments)
	configAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"log\"",
	})
	configAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"sync\"",
	})
	configAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/spf13/viper\"",
	})
	// after manipulate the code, rebuild the code
	err = configAbstractCode.RebuildCode()
	if err != nil {
		return err
	}
	// get the manipulate code
	configCode := configAbstractCode.GetCode()

	return fs.CreateFileIfNotExist(s.ConfigPath, "config.go", configCode)
}
