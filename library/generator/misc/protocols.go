package misc

import (
	"fmt"
	"strings"

	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

const (
	ProtocolHttp = "http"
)

type ProtocolType int

const (
	Http ProtocolType = iota
)

func (r ProtocolType) ToString() string {
	protocolType := ""
	switch r {
	case Http:
		protocolType = ProtocolHttp
	}
	return protocolType
}

type ProtocolOption struct {
	Options
	DirPath    string
	IsModule   bool
	ModuleName string
	RouteType  string
	RoutePath  string
}

func NewProtocols(opt ProtocolOption) AppGenerator {
	if opt.IsModule {
		opt.DirPath = fs.ConcatDirPath(opt.DirPath, "handlers")
	} else {
		opt.DirPath = fs.ConcatDirPath(opt.DirPath, "protocols")
	}
	return ProtocolOption{
		Options:    opt.Options,
		IsModule:   opt.IsModule,
		ModuleName: opt.ModuleName,
		DirPath:    opt.DirPath,
		RouteType:  opt.RouteType,
		RoutePath:  fs.ConcatDirPath(opt.DirPath, opt.RouteType),
	}
}

func (o ProtocolOption) Run() error {
	var err error

	if o.IsModule {
		err = o.createProtocolModuleDir()
	} else {
		err = o.createProtocolInternalDir()
	}

	return err
}

func (o ProtocolOption) createProtocolDir() error {
	err := fs.CreateFolderIsNotExist(o.RoutePath)
	if err != nil {
		return err
	}
	return nil
}

func (o ProtocolOption) createProtocolInternalDir() error {
	err := fs.CreateFolderIsNotExist(o.DirPath)
	if err != nil {
		return err
	}

	err = o.createProtocolDir()
	if err != nil {
		return err
	}

	switch o.RouteType {
	case ProtocolHttp:
		err = o.createProtocolInternalHttp()
	}

	if err != nil {
		return err
	}

	return nil
}

func (o ProtocolOption) createProtocolInternalHttp() error {
	var err error

	err = o.createInternalHttpErrorDir()
	if err != nil {
		return err
	}

	err = o.createInternalHttpMiddlewareDir()
	if err != nil {
		return err
	}

	err = o.createInternalHttpResponseDir()
	if err != nil {
		return err
	}

	err = o.createInternalHttpRouteDir()
	if err != nil {
		return err
	}

	err = o.createProtocolInternalHttpFile()
	if err != nil {
		return err
	}

	return nil
}

func (o ProtocolOption) createProtocolInternalHttpFile() error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "http",
		Import: []templates.ImportedPackage{
			{
				FilePath: "fmt",
			},
			{
				FilePath: fs.ConcatDirPath(o.GoModName, "config"),
			},
			{
				FilePath: fs.ConcatDirPath(o.GoModName, "internal/protocols/http/router"),
			},
			{
				FilePath: "net/http",
			},
			{
				FilePath: "github.com/go-chi/chi/v5",
			},
		},
		IsDependency: true,
		Dependency: templates.Dependency{
			HaveInterface:  true,
			DependencyName: "Http",
			FuncParams: []templates.DependencyFuncParam{
				{
					ParamName:     "HttpRouter",
					ParamDataType: "*router.HttpRouterImpl",
				},
			},
			DependencyMethod: []templates.DependencyMethod{
				{
					MethodImpl: `func (p *HttpImpl) setupRouter(app *chi.Mux) {
						p.HttpRouter.Router(app)
					}`,
				},
				{
					Method: "Listen()",
					MethodImpl: `func (p *HttpImpl) Listen() {
						app := chi.NewRouter()

						p.setupRouter(app)

						http.ListenAndServe(fmt.Sprintf(":%d", config.Get().Application.Port), app)
					}`,
				},
			},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	err = fs.CreateFileIfNotExist(o.RoutePath, fmt.Sprintf("%s.go", o.RouteType), templateString)
	if err != nil {
		return nil
	}

	return nil
}

func (o ProtocolOption) createInternalHttpErrorDir() error {
	path := fs.ConcatDirPath(o.RoutePath, "errors")
	err := fs.CreateFolderIsNotExist(path)
	if err != nil {
		return err
	}

	err = o.createInternalErrorHttpFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (o ProtocolOption) createInternalErrorHttpFile(path string) error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "errors",
		Template:    templates.ProtocolHttpChiErrorTemplate,
		Import: []templates.ImportedPackage{
			{
				FilePath: "fmt",
			},
			{
				FilePath: "net/http",
			},
			{
				FilePath: "regexp",
			},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	fs.CreateFileIfNotExist(path, "error.go", templateString)

	return nil
}

func (o ProtocolOption) createInternalHttpMiddlewareDir() error {
	path := fs.ConcatDirPath(o.RoutePath, "middleware")
	err := fs.CreateFolderIsNotExist(path)
	if err != nil {
		return err
	}

	err = o.createInternalMiddlewareHttpFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (o ProtocolOption) createInternalMiddlewareHttpFile(path string) error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "middleware",
		Template:    templates.ProtocolHttpChiMiddlewareTemplate,
		Import: []templates.ImportedPackage{
			{
				FilePath: "fmt",
			},
			{
				FilePath: fs.ConcatDirPath(o.GoModName, "config"),
			},
			{
				FilePath: "net/http",
			},
			{
				FilePath: "strings",
			},
			{
				FilePath: "time",
			},
			{
				Alias:    "httpresponse",
				FilePath: fs.ConcatDirPath(o.GoModName, "internal/protocols/http/response"),
			},
			{
				FilePath: fs.ConcatDirPath(o.GoModName, "internal/utils/encryption"),
			},
			{
				FilePath: "github.com/golang-jwt/jwt",
			},
			{
				FilePath: "github.com/golang-jwt/jwt/request",
			},
			{
				FilePath: "github.com/rs/zerolog/log",
			},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	fs.CreateFileIfNotExist(path, "jwt.go", templateString)

	return nil
}

func (o ProtocolOption) createInternalHttpResponseDir() error {
	path := fs.ConcatDirPath(o.RoutePath, "response")
	err := fs.CreateFolderIsNotExist(path)
	if err != nil {
		return err
	}

	err = o.createInternalHttpResponseFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (o ProtocolOption) createInternalHttpResponseFile(path string) error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "response",
		Template:    templates.ProtocolHttpChiResponseTemplate,
		Import: []templates.ImportedPackage{
			{
				FilePath: "encoding/json",
			},
			{
				FilePath: fs.ConcatDirPath(o.GoModName, "internal/protocols/http/errors"),
			},
			{
				FilePath: "net/http",
			},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	fs.CreateFileIfNotExist(path, "response.go", templateString)

	return nil
}

func (o ProtocolOption) createInternalHttpRouteDir() error {
	path := fs.ConcatDirPath(o.RoutePath, "router")
	err := fs.CreateFolderIsNotExist(path)
	if err != nil {
		return err
	}

	err = o.createInternalHttpRouteFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (o ProtocolOption) createInternalHttpRouteFile(path string) error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "router",
		Import: []templates.ImportedPackage{
			{
				FilePath: "github.com/go-chi/chi/v5",
			},
		},
		IsDependency: true,
		Dependency: templates.Dependency{
			HaveInterface:  false,
			DependencyName: "HttpRouter",
			DependencyMethod: []templates.DependencyMethod{
				{
					MethodImpl: "func (h *HttpRouterImpl) Router(r *chi.Mux) {}",
				},
			},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	err = fs.CreateFileIfNotExist(path, "route.go", templateString)
	if err != nil {
		return nil
	}

	return nil
}

// Package for module
// start over here

func (o ProtocolOption) createProtocolModuleDir() error {
	fs.CreateFolderIsNotExist(o.DirPath)

	o.createProtocolDir()
	err := o.createBaseModuleFile()
	if err != nil {
		return err
	}

	return nil
}

func (o ProtocolOption) createBaseModuleFile() error {
	var tmplBaseFile templates.Template
	var tmplModuleHandlerFile templates.Template

	switch o.RouteType {
	case ProtocolHttp:
		tmplBaseFile = o.createbaseModuleHttpFile()
		tmplModuleHandlerFile = o.createModuleHttpFile()
	}

	templateBaseFileString, err := tmplBaseFile.Render()
	if err != nil {
		return err
	}

	templateModuleHandlerFileString, err := tmplModuleHandlerFile.Render()
	if err != nil {
		return err
	}

	baseHandlerFile := fmt.Sprintf("%s_handler.go", o.RouteType)

	if o.ModuleName != "" {
		fs.CreateFileIfNotExist(o.RoutePath, fmt.Sprintf("%s.go", o.ModuleName), templateModuleHandlerFileString)
		err = o.appendModuleHandlerIntoMainHandler(fs.ConcatDirPath(o.RoutePath, baseHandlerFile))
		if err != nil {
			return err
		}
	} else {
		fs.CreateFileIfNotExist(o.RoutePath, baseHandlerFile, templateBaseFileString)
	}

	return nil
}

func (o ProtocolOption) createbaseModuleHttpFile() templates.Template {
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

func (o ProtocolOption) createModuleHttpFile() templates.Template {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "http",
	})

	return tmpl
}

func (o ProtocolOption) appendModuleHandlerIntoMainHandler(handlerFilePath string) error {
	servicePath := fmt.Sprintf("%s/src/modules/%s/service", o.GoModName, o.ModuleName)
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

	importAlias := fmt.Sprintf("%ssvc", o.ModuleName)
	newImport = append(newImport, templates.ImportedPackage{
		Alias:    importAlias,
		FilePath: servicePath,
	})
	newFuncParam = append(newFuncParam, templates.DependencyFuncParam{
		ParamName:     fmt.Sprintf("%sSvc", o.ModuleName),
		ParamDataType: fmt.Sprintf("%s.%sService", importAlias, strings.Title(o.ModuleName)),
	})

	tmpl := templates.NewTemplate(templates.Template{
		PackageName:  o.RouteType,
		Import:       newImport,
		IsDependency: true,
		Dependency: templates.Dependency{
			HaveInterface:    true,
			DependencyName:   fmt.Sprintf("%sHandler", strings.Title(o.RouteType)),
			FuncParams:       newFuncParam,
			DependencyMethod: dependencyMethods,
		},
	})

	template, err := tmpl.Render()
	if err != nil {
		return err
	}

	baseHandlerFile := fmt.Sprintf("%s_handler.go", o.RouteType)

	fs.ReplaceFile(o.RoutePath, baseHandlerFile, template)

	return nil
}
