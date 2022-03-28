package appmodule

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/nurcahyaari/kite/lib/impl"
// 	"github.com/nurcahyaari/kite/templates"
// 	"github.com/nurcahyaari/kite/utils/fs"
// )

// type ServiceOption struct {
// 	impl.KiteOptions
// 	ModuleName        string
// 	ModulePath        string
// 	ModuleServicePath string
// }

// func NewService(opt ServiceOption) (impl.AppGenerator, error) {
// 	servicePath := fs.ConcatDirPath(opt.ModulePath, "service")
// 	return ServiceOption{
// 		KiteOptions:       opt.KiteOptions,
// 		ModuleName:        opt.ModuleName,
// 		ModulePath:        opt.ModulePath,
// 		ModuleServicePath: servicePath,
// 	}, nil
// }

// func (o ServiceOption) Run() error {

// 	err := o.createServicerDir()
// 	if err != nil {
// 		return err
// 	}

// 	err = o.createBaseFile()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (o ServiceOption) createServicerDir() error {
// 	err := fs.CreateFolderIsNotExist(o.ModuleServicePath)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (o ServiceOption) createBaseFile() error {
// 	tmpl := templates.NewTemplate(templates.Template{
// 		PackageName:  "service",
// 		IsDependency: true,
// 		Import: []templates.ImportedPackage{
// 			{
// 				Alias:    fmt.Sprintf("%srepo", strings.ToLower(o.ModuleName)),
// 				FilePath: fmt.Sprintf("%s/src/modules/%s/repository", o.GoModName, o.ModuleName),
// 			},
// 		},
// 		Dependency: templates.Dependency{
// 			HaveInterface:  true,
// 			DependencyName: fmt.Sprintf("%sService", strings.Title(o.ModuleName)),
// 			FuncParams: []templates.DependencyFuncParam{
// 				{
// 					ParamName:     fmt.Sprintf("%sRepo", strings.Title(o.ModuleName)),
// 					ParamDataType: fmt.Sprintf("%srepo.%sRepository", strings.ToLower(o.ModuleName), strings.Title(o.ModuleName)),
// 				},
// 			},
// 			DependencyMethod: []templates.DependencyMethod{},
// 		},
// 	})

// 	templateString, err := tmpl.Render()
// 	if err != nil {
// 		return err
// 	}

// 	err = fs.CreateFileIfNotExist(o.ModuleServicePath, "service.go", templateString)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
