package handler

import (
	"fmt"
	"go/parser"
	"strings"

	"github.com/nurcahyaari/kite/lib/ast"
	"github.com/nurcahyaari/kite/lib/generator/protocol"
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type HttpHandlerGen interface {
	CreateHttpHandlerBaseDir() error
	CreateHttpHandlerBaseFile() error
	CreateHttpHandlerBaseModuleFile() error
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
	// var tmplBaseFile templates.Template
	tmplBaseFile := s.createbaseModuleHttpFile()
	templateBaseFileString, err := tmplBaseFile.Render()
	if err != nil {
		return err
	}

	baseHandlerFile := fmt.Sprintf("%s_handler.go", protocol.Http.ToString())
	if !fs.IsFileExist(fs.ConcatDirPath(s.HandlerPath, baseHandlerFile)) {
		fs.CreateFileIfNotExist(s.HandlerPath, baseHandlerFile, templateBaseFileString)
	}
	return nil
}

func (s *HttpHandlerGenImpl) CreateHttpHandlerBaseModuleFile() error {
	baseHandlerFile := fmt.Sprintf("%s_handler.go", protocol.Http.ToString())
	// var tmplModuleHandlerFile templates.Template
	tmplModuleHandlerFile := s.createModuleHttpFile()

	templateModuleHandlerFileString, err := tmplModuleHandlerFile.Render()
	if err != nil {
		return err
	}

	if !fs.IsFileExist(fs.ConcatDirPath(s.HandlerPath, fmt.Sprintf("%s.go", s.ModuleName))) {
		fs.CreateFileIfNotExist(s.HandlerPath, fmt.Sprintf("%s.go", s.ModuleName), templateModuleHandlerFileString)
	}
	err = s.appendModuleHandlerIntoBaseHandler(fs.ConcatDirPath(s.HandlerPath, baseHandlerFile))
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

func (s *HttpHandlerGenImpl) appendModuleHandlerIntoBaseHandler(handlerFilePath string) error {
	servicePath := fmt.Sprintf("%s/src/module/%s/service", s.GomodName, s.ModuleName)
	val, err := fs.ReadFile(handlerFilePath)
	if err != nil {
		return err
	}
	importAlias := fmt.Sprintf("%ssvc", s.ModuleName)

	abstractCode := ast.NewAbstractCode(val, parser.ParseComments)
	abstractCode.AddImport(ast.ImportSpec{
		Name: importAlias,
		Path: fmt.Sprintf("\"%s\"", servicePath),
	})
	abstractCode.AddFunctionArgs(ast.FunctionSpec{
		Name: "NewHttpHandler",
		Args: ast.FunctionArgList{
			&ast.FunctionArg{
				Name:     fmt.Sprintf("%sSvc", s.ModuleName),
				LibName:  importAlias,
				DataType: fmt.Sprintf("%sService", strings.Title(s.ModuleName)),
			},
		},
	})
	abstractCode.AddStructVarDecl(ast.StructArgList{
		&ast.StructArg{
			StructName: "HttpHandlerImpl",
			DataType: ast.StructDtypes{
				LibName:  importAlias,
				TypeName: fmt.Sprintf("%sService", strings.Title(s.ModuleName)),
			},
			IsPointer: false,
		},
	})
	abstractCode.AddFunctionArgsToReturn(ast.FunctionReturnArgsSpec{
		FuncName:      "NewHttpHandler",
		ReturnName:    "HttpHandlerImpl",
		DataTypeKey:   fmt.Sprintf("%sService", strings.Title(s.ModuleName)),
		DataTypeValue: fmt.Sprintf("%sSvc", s.ModuleName),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	newCode := abstractCode.GetCode()

	baseHandlerFile := fmt.Sprintf("%s_handler.go", protocol.ProtocolHttp)

	fs.ReplaceFile(s.HandlerPath, baseHandlerFile, newCode)

	return nil
}
