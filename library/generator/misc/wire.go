package misc

import (
	"fmt"

	"github.com/nurcahyaari/kite/library/impl"
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type WireModuleOptions struct {
	ServicePath    string
	RepositoryPath string
}

type WireOptions struct {
	impl.GeneratorOptions
	WireModuleOptions
	IsNewModule bool
}

func NewWire(options WireOptions) impl.AppGenerator {
	return WireOptions{
		GeneratorOptions:  options.GeneratorOptions,
		WireModuleOptions: options.WireModuleOptions,
		IsNewModule:       options.IsNewModule,
	}
}

func (o WireOptions) Run() error {
	if o.IsNewModule {
		err := o.appendDependencyToWire()
		if err != nil {
			return err
		}
	} else {
		err := o.createWireFile()
		if err != nil {
			return err
		}
	}

	return nil
}

func (o WireOptions) createWireFile() error {
	var tmpl templates.Template
	if o.WireModuleOptions.ServicePath == "" {
		tmpl = templates.NewTemplate(templates.Template{
			PackageName: "main",
			Template:    templates.WireTemplate,
			Header:      "//+build wireinject",
			Import: []templates.ImportedPackage{
				{
					FilePath: "github.com/google/wire",
				},
				{
					Alias:    "httprouter",
					FilePath: fmt.Sprintf("%s/internal/protocols/http/router", o.GoModName),
				},
				{
					Alias:    "httphandler",
					FilePath: fmt.Sprintf("%s/src/handlers/http", o.GoModName),
				},
				{
					FilePath: fmt.Sprintf("%s/internal/protocols/http", o.GoModName),
				},
			},
		})
	} else {

	}

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	err = fs.CreateFileIfNotExist(o.ProjectPath, "wire.go", templateString)
	if err != nil {
		return nil
	}

	return nil
}

func (o WireOptions) appendDependencyToWire() error {
	return nil
}
