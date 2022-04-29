package protocolhttpgen

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates"
	"github.com/nurcahyaari/kite/internal/templates/protocoltemplate/httptemplate/chitemplate"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/internal/utils/ast"
)

type ProtocolHttpGen interface {
	createInternalHttpErrorDir(path string) error
	createInternalErrorHttpFile(path string) error
	createInternalHttpMiddlewareDir(path string) error
	createInternalMiddlewareHttpFile(path string) error
	createInternalHttpResponseDir(path string) error
	createInternalHttpResponseFile(path string) error
	createProtocolInternalHttpFile(path string) error
	createInternalHttpRouteDir(path string) error
	createInternalHttpRouteFile(path string) error
	CreateProtocolSrcHttpBaseFile(path string) error
	CreateProtocolInternalHttp(path string) error
	CreateProtocolSrcHttpHandler(path string, name string) error
}

type ProtocolHttpGenImpl struct {
	Path      string
	GomodName string
	fs        database.FileSystem
}

func NewProtocolHttp(
	fs database.FileSystem,
) *ProtocolHttpGenImpl {
	return &ProtocolHttpGenImpl{
		fs: fs,
	}
}

// create internal/http/error/ directory
func (s *ProtocolHttpGenImpl) createInternalHttpErrorDir(path string) error {
	path = utils.ConcatDirPath(path, "errors")
	err := s.fs.CreateFolderIfNotExists(path)
	if err != nil {
		return err
	}

	err = s.createInternalErrorHttpFile(path)
	if err != nil {
		return err
	}

	return nil
}

// create internal/http/error/error.go file
func (s *ProtocolHttpGenImpl) createInternalErrorHttpFile(path string) error {
	errorTemplate := chitemplate.NewErrorTemplate()
	errorTemplateString, err := errorTemplate.Render()
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(errorTemplateString, parser.ParseComments)
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"fmt\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"net/http\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"regexp\"",
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}

	errorTemplateString = abstractCode.GetCode()

	s.fs.CreateFileIfNotExists(path, "error.go", errorTemplateString)

	return nil
}

// create internal/http/middleware/ directory
func (s *ProtocolHttpGenImpl) createInternalHttpMiddlewareDir(path string) error {
	path = utils.ConcatDirPath(path, "middleware")
	err := s.fs.CreateFolderIfNotExists(path)
	if err != nil {
		return err
	}

	err = s.createInternalMiddlewareHttpFile(path)
	if err != nil {
		return err
	}

	return nil
}

// create internal/http/middleware/middleware.go file
func (s *ProtocolHttpGenImpl) createInternalMiddlewareHttpFile(path string) error {
	middlewareTemplate := chitemplate.NewMiddlewareTemplate()
	middlewareTemplateString, err := middlewareTemplate.Render()
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(middlewareTemplateString, parser.ParseComments)
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"fmt\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(s.GomodName, "config")),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"net/http\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"strings\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"time\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"net/http\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Name: "httpresponse",
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(s.GomodName, "internal/protocol/http/response")),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(s.GomodName, "internal/utils/encryption")),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/golang-jwt/jwt\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/golang-jwt/jwt/request\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/rs/zerolog/log\"",
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}

	middlewareTemplateString = abstractCode.GetCode()

	s.fs.CreateFileIfNotExists(path, "jwt.go", middlewareTemplateString)

	return nil
}

// create internal/http/response/ directory
func (s *ProtocolHttpGenImpl) createInternalHttpResponseDir(path string) error {
	path = utils.ConcatDirPath(path, "response")
	err := s.fs.CreateFolderIfNotExists(path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpResponseFile(path)
	if err != nil {
		return err
	}

	return nil
}

// create internal/http/response/response.go file
func (s *ProtocolHttpGenImpl) createInternalHttpResponseFile(path string) error {
	responseTemplate := chitemplate.NewResponseTemplate()
	responseTemplateString, err := responseTemplate.Render()
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(responseTemplateString, parser.ParseComments)
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"encoding/json\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"net/http\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Name: "httperror",
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(s.GomodName, "internal/protocol/http/errors")),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}

	responseTemplateString = abstractCode.GetCode()

	s.fs.CreateFileIfNotExists(path, "response.go", responseTemplateString)

	return nil
}

// create internal/http/router/ directory
func (s *ProtocolHttpGenImpl) createInternalHttpRouteDir(path string) error {
	path = utils.ConcatDirPath(path, "router")
	err := s.fs.CreateFolderIfNotExists(path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpRouteFile(path)
	if err != nil {
		return err
	}

	return nil
}

// create internal/http/router/route.go file
func (s *ProtocolHttpGenImpl) createInternalHttpRouteFile(path string) error {
	templateNew := chitemplate.NewInternalHttpRouterTemplate()
	templateCode, err := templateNew.Render()
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddFunction(ast.FunctionSpecList{
		&ast.FunctionSpec{
			Name: "NewHttpRouter",
			StructSpec: &ast.FunctionStructSpec{
				IsConstruct: true,
				DataTypes:   "HttpRouterImpl",
			},
			Args: ast.FunctionArgList{
				&ast.FunctionArg{
					IsPointer: true,
					Name:      "handler",
					LibName:   "http",
					DataType:  "HttpHandlerImpl",
				},
			},
			Returns: &ast.FunctionReturnSpecList{
				&ast.FunctionReturnSpec{
					IsPointer: true,
					IsStruct:  true,
					DataType:  "HttpRouterImpl",
					Return:    "HttpRouterImpl",
				},
			},
		},
	})
	abstractCode.AddFunctionArgsToReturn(ast.FunctionReturnArgsSpec{
		FuncName:      "NewHttpRouter",
		ReturnName:    "HttpRouterImpl",
		DataTypeKey:   "handler",
		DataTypeValue: "handler",
	})
	abstractCode.AddStructs(ast.StructSpecList{
		&ast.StructSpec{
			Name: "HttpRouterImpl",
		},
	})
	abstractCode.AddStructVarDecl(ast.StructArgList{
		&ast.StructArg{
			StructName: "HttpRouterImpl",
			Name:       "handler",
			DataType: ast.StructDtypes{
				LibName:  "http",
				TypeName: "HttpHandlerImpl",
			},
			IsPointer: true,
		},
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/go-chi/chi/v5\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/go-chi/cors\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Name: "httpswagger",
		Path: "\"github.com/swaggo/http-swagger\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(s.GomodName, "src/handler/http")),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateCode = abstractCode.GetCode()

	err = s.fs.CreateFileIfNotExists(path, "route.go", templateCode)
	if err != nil {
		return nil
	}

	return nil
}

// create internal/http/http.go file
func (s *ProtocolHttpGenImpl) createProtocolInternalHttpFile(path string) error {
	internalHttpTemplate := chitemplate.NewInternalHttpTemplate()
	internalHttpTemplateString, err := internalHttpTemplate.Render()
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(internalHttpTemplateString, parser.ParseComments)
	abstractCode.AddFunction(ast.FunctionSpecList{
		&ast.FunctionSpec{
			Name: "NewHttp",
			StructSpec: &ast.FunctionStructSpec{
				IsConstruct: true,
				DataTypes:   "HttpImpl",
			},
			Args: ast.FunctionArgList{
				&ast.FunctionArg{
					IsPointer: true,
					Name:      "httpRouter",
					LibName:   "router",
					DataType:  "HttpRouterImpl",
				},
			},
			Returns: &ast.FunctionReturnSpecList{
				&ast.FunctionReturnSpec{
					IsPointer: true,
					IsStruct:  true,
					DataType:  "HttpImpl",
					Return:    "HttpImpl",
				},
			},
		},
	})
	abstractCode.AddFunctionArgsToReturn(ast.FunctionReturnArgsSpec{
		FuncName:      "NewHttp",
		ReturnName:    "HttpImpl",
		DataTypeKey:   "HttpRouter",
		DataTypeValue: "httpRouter",
	})
	abstractCode.AddStructs(ast.StructSpecList{
		&ast.StructSpec{
			Name: "HttpImpl",
		},
	})
	abstractCode.AddStructVarDecl(ast.StructArgList{
		&ast.StructArg{
			StructName: "HttpImpl",
			IsPointer:  true,
			Name:       "HttpRouter",
			DataType: ast.StructDtypes{
				LibName:  "router",
				TypeName: "HttpRouterImpl",
			},
		},
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"fmt\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"net/http\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/go-chi/chi/v5\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(s.GomodName, "config")),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(s.GomodName, "internal/protocol/http/router")),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateBaseFileString := abstractCode.GetCode()

	err = s.fs.CreateFileIfNotExists(path, fmt.Sprintf("%s.go", "http"), templateBaseFileString)
	if err != nil {
		return nil
	}

	return nil
}

// create internal/http directory and all of the assets
func (s *ProtocolHttpGenImpl) CreateProtocolInternalHttp(path string) error {
	var err error

	err = s.createInternalHttpErrorDir(path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpMiddlewareDir(path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpResponseDir(path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpRouteDir(path)
	if err != nil {
		return err
	}

	err = s.createProtocolInternalHttpFile(path)
	if err != nil {
		return err
	}

	return nil
}

// create src/http directory and the assets
func (s *ProtocolHttpGenImpl) CreateProtocolSrcHttpBaseFile(path string) error {
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

	baseHandlerFile := "http_handler.go"
	if !s.fs.IsFileExists(utils.ConcatDirPath(path, baseHandlerFile)) {
		s.fs.CreateFileIfNotExists(path, baseHandlerFile, templateBaseFileString)
	}
	return nil
}

func (s ProtocolHttpGenImpl) CreateProtocolSrcHttpHandler(path string, name string) error {
	if exist := s.fs.IsFileExists(utils.ConcatDirPath(path, "http_handler.go")); !exist {
		s.CreateProtocolSrcHttpBaseFile(path)
	}
	template := templates.NewTemplateNewImpl("http", "")
	templateCode, err := template.Render("", nil)
	if err != nil {
		return err
	}

	if !s.fs.IsFileExists(utils.ConcatDirPath(path, name)) {
		s.fs.CreateFileIfNotExists(path, fmt.Sprintf("%s.go", name), templateCode)
	}

	return nil
}
