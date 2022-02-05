package misc

import (
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type ConfigOption struct {
	Options
	DirName string
	DirPath string
}

func NewConfig(options ConfigOption) AppGenerator {
	options.DirName = "config"
	options.DirPath = fs.ConcatDirPath(options.ProjectPath, options.DirName)

	return options
}

func (o ConfigOption) Run() error {

	err := o.createConfigDir()
	if err != nil {
		return err
	}
	err = o.createBaseFile()
	if err != nil {
		return err
	}

	return nil
}

func (o ConfigOption) createConfigDir() error {
	err := fs.CreateFolderIsNotExist(o.DirPath)
	if err != nil {
		return err
	}

	return nil
}

func (o ConfigOption) createBaseFile() error {
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
			"AppName":  o.Options.AppName,
			"DBDialeg": o.DefaultDBDialeg,
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	err = fs.CreateFileIfNotExist(o.DirPath, "config.go", templateString)
	if err != nil {
		return err
	}

	return nil
}
