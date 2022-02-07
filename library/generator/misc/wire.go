package misc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/nurcahyaari/kite/library/impl"
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

// TODO: implement auto add modules into wire.go

type Service struct {
	ServicePath     string
	ServiceName     string
	ServiceImplName string
}

type Repository struct {
	RepositoryPath     string
	RepositoryName     string
	RepositoryNameImpl string
}

type WireModuleOption struct {
	ModuleName string
	Service    Service
	Repository Repository
}

type WireModulesOption []*WireModuleOption

type WireOptions struct {
	impl.GeneratorOptions
	IsNewModule bool
}

func NewWire(options WireOptions) impl.AppGenerator {
	return WireOptions{
		GeneratorOptions: options.GeneratorOptions,
		IsNewModule:      options.IsNewModule,
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
	tmpl := templates.NewTemplate(templates.Template{
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
			{
				FilePath: fmt.Sprintf("%s/infrastructure", o.GoModName),
			},
		},
	})

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

func (o WireOptions) readExistingModules() (WireModulesOption, error) {
	wireModuleOpt := WireModulesOption{}
	moduleDir := fs.ConcatDirPath(o.ProjectPath, "src/modules/")
	files, err := ioutil.ReadDir(moduleDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		wireModuleOpt = append(wireModuleOpt, &WireModuleOption{
			ModuleName: f.Name(),
			Service: Service{
				ServicePath:     fs.ConcatDirPath(moduleDir, fmt.Sprintf("%s/service", f.Name())),
				ServiceName:     fmt.Sprintf("%sService", strings.Title(f.Name())),
				ServiceImplName: fmt.Sprintf("%sServiceImpl", strings.Title(f.Name())),
			},
			Repository: Repository{
				RepositoryPath:     fs.ConcatDirPath(moduleDir, fmt.Sprintf("%s/repository", f.Name())),
				RepositoryName:     fmt.Sprintf("%sRepository", strings.Title(f.Name())),
				RepositoryNameImpl: fmt.Sprintf("%sRepositoryImpl", strings.Title(f.Name())),
			},
		})

	}

	return wireModuleOpt, nil
}

func (o WireOptions) appendDependencyToWire() error {
	// wireModuleOpt, err := o.readExistingModules()
	// if err != nil {
	// 	return err
	// }

	return nil
}
