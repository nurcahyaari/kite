package module

import (
	"fmt"
	"strings"

	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type ServiceGen interface {
	CreateServiceDir() error
	CreateServiceFile() error
}

type ServiceGenImpl struct {
	ServicePath string
	ModuleName  string
	GomodName   string
}

func NewServiceGen(moduleName, modulePath, gomodName string) *ServiceGenImpl {
	ServicePath := fs.ConcatDirPath(modulePath, "service")
	return &ServiceGenImpl{
		ServicePath: ServicePath,
		ModuleName:  moduleName,
		GomodName:   gomodName,
	}
}

func (s *ServiceGenImpl) CreateServiceDir() error {
	err := fs.CreateFolderIsNotExist(s.ServicePath)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceGenImpl) CreateServiceFile() error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName:  "service",
		IsDependency: true,
		Import: []templates.ImportedPackage{
			{
				Alias:    fmt.Sprintf("%srepo", strings.ToLower(s.ModuleName)),
				FilePath: fmt.Sprintf("%s/src/module/%s/repository", s.GomodName, s.ModuleName),
			},
		},
		Dependency: templates.Dependency{
			HaveInterface:  true,
			DependencyName: fmt.Sprintf("%sService", strings.Title(s.ModuleName)),
			FuncParams: []templates.DependencyFuncParam{
				{
					ParamName:     fmt.Sprintf("%sRepo", strings.Title(s.ModuleName)),
					ParamDataType: fmt.Sprintf("%srepo.%sRepository", strings.ToLower(s.ModuleName), strings.Title(s.ModuleName)),
				},
			},
			DependencyMethod: []templates.DependencyMethod{},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	err = fs.CreateFileIfNotExist(s.ServicePath, "service.go", templateString)
	if err != nil {
		return err
	}

	return nil
}
