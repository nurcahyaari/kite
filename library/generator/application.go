package generator

import (
	"fmt"

	"github.com/nurcahyaari/kite/library/impl"
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
	"github.com/nurcahyaari/kite/utils/logger"
)

type ApplicationOption struct {
	impl.GeneratorOptions
}

func NewApplicationGenerator(options ApplicationOption) impl.AppGenerator {
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
		Data: map[string]interface{}{
			"GoGenerate": []string{
				"//go:generate go run github.com/google/wire/cmd/wire",
				"//go:generate go run github.com/swaggo/swag/cmd/swag init",
			},
		},
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
