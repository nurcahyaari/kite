package handler

import (
	"fmt"
	"strings"

	"github.com/nurcahyaari/kite/lib/generator/protocol"
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type HttpHandlerGen interface {
	CreateHttpHandlerBaseDir() error
	CreateHttpHandlerBaseFile() error
}

type HttpHandlerGenImpl struct {
	ModuleName  string
	HandlerPath string
	GomodName   string
}

func NewHttpHandlerGen(moduleName, modulePath, gomodName string) *HttpHandlerGenImpl {
	handlerPath := fs.ConcatDirPath(modulePath, protocol.Http.ToString())
	return &HttpHandlerGenImpl{
		ModuleName:  moduleName,
		HandlerPath: handlerPath,
		GomodName:   gomodName,
	}
}

func (s *HttpHandlerGenImpl) CreateHttpHandlerBaseDir() error {
	err := fs.CreateFolderIsNotExist(s.HandlerPath)
	if err != nil {
		return err
	}

	return nil
}

func (s *HttpHandlerGenImpl) CreateHttpHandlerBaseFile() error {
	var tmplBaseFile templates.Template
	var tmplModuleHandlerFile templates.Template

	tmplBaseFile = s.createbaseModuleHttpFile()
	tmplModuleHandlerFile = s.createModuleHttpFile()

	templateBaseFileString, err := tmplBaseFile.Render()
	if err != nil {
		return err
	}

	templateModuleHandlerFileString, err := tmplModuleHandlerFile.Render()
	if err != nil {
		return err
	}

	baseHandlerFile := fmt.Sprintf("%s_handler.go", protocol.Http.ToString())

	if !fs.IsFileExist(fs.ConcatDirPath(s.HandlerPath, baseHandlerFile)) {
		fs.CreateFileIfNotExist(s.HandlerPath, baseHandlerFile, templateBaseFileString)
	}

	if !fs.IsFileExist(fs.ConcatDirPath(s.HandlerPath, fmt.Sprintf("%s.go", s.ModuleName))) {
		fs.CreateFileIfNotExist(s.HandlerPath, fmt.Sprintf("%s.go", s.ModuleName), templateModuleHandlerFileString)
	}
	err = s.appendModuleHandlerIntoMainHandler(fs.ConcatDirPath(s.HandlerPath, baseHandlerFile))
	if err != nil {
		return err
	}

	return nil
}

func (s *HttpHandlerGenImpl) createbaseModuleHttpFile() templates.Template {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "http",
		Import: []templates.ImportedPackage{
			{
				FilePath: "github.com/go-chi/chi/v5",
			},
		},
		IsDependency: true,
		Dependency: templates.Dependency{
			HaveInterface:  true,
			DependencyName: "HttpHandler",
			FuncParams:     []templates.DependencyFuncParam{},
			DependencyMethod: []templates.DependencyMethod{
				{
					Method:     "Router(r *chi.Mux)",
					MethodImpl: "func (h *HttpHandlerImpl) Router(r *chi.Mux) {}",
				},
			},
		},
	})

	return tmpl
}

func (s *HttpHandlerGenImpl) createModuleHttpFile() templates.Template {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "http",
	})

	return tmpl
}

func (s *HttpHandlerGenImpl) appendModuleHandlerIntoMainHandler(handlerFilePath string) error {
	servicePath := fmt.Sprintf("%s/src/module/%s/service", s.GomodName, s.ModuleName)
	val, err := fs.ReadFile(handlerFilePath)
	if err != nil {
		return err
	}

	importedPackages := fs.ReadImportedPackages(val)
	dependencyInjected := fs.ReadStructWithObject(val)
	methodList := fs.ReadInterfaceWithMethod(val)
	methodImplList := fs.ReadMethodImpl(val)

	newImport := []templates.ImportedPackage{}
	for _, i := range importedPackages {
		newImport = append(newImport, templates.ImportedPackage{
			Alias:    i.Alias,
			FilePath: i.FilePath,
		})
	}

	newFuncParam := []templates.DependencyFuncParam{}
	for _, d := range dependencyInjected {
		newFuncParam = append(newFuncParam, templates.DependencyFuncParam{
			ParamName:     d.ObjectName,
			ParamDataType: d.ObjectDataType,
		})
	}

	dependencyMethods := []templates.DependencyMethod{}
	for _, d := range methodList {
		dependencyMethods = append(dependencyMethods, templates.DependencyMethod{
			Method:     d.Method,
			MethodImpl: "",
		})
	}

	for i, _ := range methodImplList {
		dependencyMethods[i].MethodImpl = methodImplList[i]
	}

	importAlias := fmt.Sprintf("%ssvc", s.ModuleName)
	newImport = append(newImport, templates.ImportedPackage{
		Alias:    importAlias,
		FilePath: servicePath,
	})
	newFuncParam = append(newFuncParam, templates.DependencyFuncParam{
		ParamName:     fmt.Sprintf("%sSvc", s.ModuleName),
		ParamDataType: fmt.Sprintf("%s.%sService", importAlias, strings.Title(s.ModuleName)),
	})

	tmpl := templates.NewTemplate(templates.Template{
		PackageName:  protocol.Http.ToString(),
		Import:       newImport,
		IsDependency: true,
		Dependency: templates.Dependency{
			HaveInterface:    true,
			DependencyName:   fmt.Sprintf("%sHandler", strings.Title(protocol.Http.ToString())),
			FuncParams:       newFuncParam,
			DependencyMethod: dependencyMethods,
		},
	})

	template, err := tmpl.Render()
	if err != nil {
		return err
	}

	baseHandlerFile := fmt.Sprintf("%s_handler.go", protocol.ProtocolHttp)

	fs.ReplaceFile(s.HandlerPath, baseHandlerFile, template)

	return nil
}
