package generator

import (
	"fmt"

	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
	"github.com/nurcahyaari/kite/utils/logger"
)

type ApplicationOption struct {
	GeneratorOptions
}

func NewApplicationGenerator(options ApplicationOption) AppGenerator {
	return options
}

func (o ApplicationOption) Run() error {

	o.createMainFile()

	return nil
}

func (o ApplicationOption) createMainFile() error {
	logger.Info("Create main.go file... ")

	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "main",
		Template:    templates.MainTemplate,
		Import: []templates.ImportedPackage{
			{
				FilePath: fmt.Sprintf("%s/internal/logger", o.GoModName),
			},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	err = fs.CreateFileIfNotExist(o.ProjectPath, "main.go", templateString)
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")
	return nil
}
