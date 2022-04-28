package configgen

import (
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates/configtemplate"
	"github.com/nurcahyaari/kite/internal/utils/ast"
)

type ConfigGen interface {
	CreateConfigDir(option ConfigDto) error
	CreateConfigFile(option ConfigDto) error
}

type ConfigGenImpl struct {
	fs database.FileSystem
}

func NewConfig(
	fs database.FileSystem,
) *ConfigGenImpl {
	return &ConfigGenImpl{
		fs: fs,
	}
}

func (s ConfigGenImpl) CreateConfigDir(option ConfigDto) error {
	err := s.fs.CreateFolderIfNotExists(option.ConfigPath)
	if err != nil {
		return err
	}

	return nil
}

func (s ConfigGenImpl) CreateConfigFile(option ConfigDto) error {
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

	return s.fs.CreateFileIfNotExists(option.ConfigPath, "config.go", configCode)
}
