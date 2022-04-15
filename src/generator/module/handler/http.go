package handler

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/src/ast"
	"github.com/nurcahyaari/kite/src/generator/protocol"
	"github.com/nurcahyaari/kite/src/templates"
	"github.com/nurcahyaari/kite/src/utils/fs"
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

	templateNew := templates.NewTemplateNewImpl("http", "")
	templateCode, err := templateNew.Render("", nil)
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddFunction(ast.FunctionSpecList{
		&ast.FunctionSpec{
			Name: "Router",
			StructSpec: &ast.FunctionStructSpec{
				Name:      "h",
				DataTypes: "HttpHandlerImpl",
			},
			Args: ast.FunctionArgList{
				&ast.FunctionArg{
					IsPointer: true,
					Name:      "r",
					LibName:   "chi",
					DataType:  "Mux",
				},
			},
		},
	})
	abstractCode.AddFunction(ast.FunctionSpecList{
		&ast.FunctionSpec{
			Name: "NewHttpHandler",
			Returns: &ast.FunctionReturnSpecList{
				&ast.FunctionReturnSpec{
					IsPointer: true,
					IsStruct:  true,
					DataType:  "HttpHandlerImpl",
					Return:    "HttpHandlerImpl",
				},
			},
		},
	})
	abstractCode.AddStructs(ast.StructSpecList{
		&ast.StructSpec{
			Name: "HttpHandlerImpl",
		},
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/go-chi/chi/v5\"",
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateBaseFileString := abstractCode.GetCode()

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

	if tmplModuleHandlerFile != nil {
		templateModuleHandlerFileString, err := tmplModuleHandlerFile.Render("", "")
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
	}

	return nil
}

func (s *HttpHandlerGenImpl) createModuleHttpFile() *templates.TemplateNewImpl {
	tmpl := templates.NewTemplateNewImpl("http", "")

	return tmpl
}

func (s *HttpHandlerGenImpl) appendModuleHandlerIntoBaseHandler(handlerFilePath string) error {
	servicePath := fmt.Sprintf("%s/src/module/%s/service", s.GomodName, s.ModuleName)
	val, err := fs.ReadFile(handlerFilePath)
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(val, parser.ParseComments)

	importAlias := fmt.Sprintf("%ssvc", s.ModuleName)
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
				DataType: "Service",
			},
		},
	})
	abstractCode.AddStructVarDecl(ast.StructArgList{
		&ast.StructArg{
			StructName: "HttpHandlerImpl",
			Name:       fmt.Sprintf("%sSvc", s.ModuleName),
			DataType: ast.StructDtypes{
				LibName:  importAlias,
				TypeName: "Service",
			},
			IsPointer: false,
		},
	})
	abstractCode.AddFunctionArgsToReturn(ast.FunctionReturnArgsSpec{
		FuncName:      "NewHttpHandler",
		ReturnName:    "HttpHandlerImpl",
		DataTypeKey:   fmt.Sprintf("%sSvc", s.ModuleName),
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
