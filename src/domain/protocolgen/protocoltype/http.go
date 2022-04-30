package protocoltype

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates"
	"github.com/nurcahyaari/kite/internal/templates/protocoltemplate/httptemplate/chitemplate"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/internal/utils/ast"
	"github.com/nurcahyaari/kite/src/domain/emptygen"
	"github.com/nurcahyaari/kite/src/domain/wiregen"
)

type ProtocolHttpGen interface {
	createInternalHttpErrorDir(dto ProtocolDto) error
	createInternalErrorHttpFile(dto ProtocolDto) error
	createInternalHttpMiddlewareDir(dto ProtocolDto) error
	createInternalMiddlewareHttpFile(dto ProtocolDto) error
	createInternalHttpResponseDir(dto ProtocolDto) error
	createInternalHttpResponseFile(dto ProtocolDto) error
	createProtocolInternalHttpFile(dto ProtocolDto) error
	createInternalHttpRouteDir(dto ProtocolDto) error
	createInternalHttpRouteFile(dto ProtocolDto) error
	CreateProtocolSrcHttpBaseFile(dto ProtocolDto) error
	CreateProtocolInternalHttp(dto ProtocolDto) error
	CreateProtocolSrcHttpHandler(dto ProtocolDto) error
}

type ProtocolHttpGenImpl struct {
	fs       database.FileSystem
	emptyGen emptygen.EmptyGen
	wireGen  wiregen.WireGen
}

func NewProtocolHttp(
	fs database.FileSystem,
	emptyGen emptygen.EmptyGen,
	wireGen wiregen.WireGen,
) *ProtocolHttpGenImpl {
	return &ProtocolHttpGenImpl{
		fs:       fs,
		emptyGen: emptyGen,
		wireGen:  wireGen,
	}
}

// create internal/http/error/ directory
func (s *ProtocolHttpGenImpl) createInternalHttpErrorDir(dto ProtocolDto) error {
	dto.Path = utils.ConcatDirPath(dto.Path, "errors")
	err := s.fs.CreateFolderIfNotExists(dto.Path)
	if err != nil {
		return err
	}

	err = s.createInternalErrorHttpFile(dto)
	if err != nil {
		return err
	}

	return nil
}

// create internal/http/error/error.go file
func (s *ProtocolHttpGenImpl) createInternalErrorHttpFile(dto ProtocolDto) error {
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

	s.fs.CreateFileIfNotExists(dto.Path, "error.go", errorTemplateString)

	return nil
}

// create internal/http/middleware/ directory
func (s *ProtocolHttpGenImpl) createInternalHttpMiddlewareDir(dto ProtocolDto) error {
	dto.Path = utils.ConcatDirPath(dto.Path, "middleware")
	err := s.fs.CreateFolderIfNotExists(dto.Path)
	if err != nil {
		return err
	}

	err = s.createInternalMiddlewareHttpFile(dto)
	if err != nil {
		return err
	}

	return nil
}

// create internal/http/middleware/middleware.go file
func (s *ProtocolHttpGenImpl) createInternalMiddlewareHttpFile(dto ProtocolDto) error {
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
		Path: "\"net/http\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"strings\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(dto.GomodName, "config")),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"time\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Name: "httpresponse",
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(dto.GomodName, "internal/protocols/http/response")),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(dto.GomodName, "internal/utils/encryption")),
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

	s.fs.CreateFileIfNotExists(dto.Path, "jwt.go", middlewareTemplateString)

	return nil
}

// create internal/http/response/ directory
func (s *ProtocolHttpGenImpl) createInternalHttpResponseDir(dto ProtocolDto) error {
	dto.Path = utils.ConcatDirPath(dto.Path, "response")
	err := s.fs.CreateFolderIfNotExists(dto.Path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpResponseFile(dto)
	if err != nil {
		return err
	}

	return nil
}

// create internal/http/response/response.go file
func (s *ProtocolHttpGenImpl) createInternalHttpResponseFile(dto ProtocolDto) error {
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
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(dto.GomodName, "internal/protocols/http/errors")),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}

	responseTemplateString = abstractCode.GetCode()

	s.fs.CreateFileIfNotExists(dto.Path, "response.go", responseTemplateString)

	return nil
}

// create internal/http/router/ directory
func (s *ProtocolHttpGenImpl) createInternalHttpRouteDir(dto ProtocolDto) error {
	dto.Path = utils.ConcatDirPath(dto.Path, "router")
	err := s.fs.CreateFolderIfNotExists(dto.Path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpRouteFile(dto)
	if err != nil {
		return err
	}

	return nil
}

// create internal/http/router/route.go file
func (s *ProtocolHttpGenImpl) createInternalHttpRouteFile(dto ProtocolDto) error {
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
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(dto.GomodName, "src/handlers/http")),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateCode = abstractCode.GetCode()

	err = s.fs.CreateFileIfNotExists(dto.Path, "route.go", templateCode)
	if err != nil {
		return nil
	}

	err = s.wireGen.AddDependencyAfterCreatingModule(wiregen.WireAddModuleDto{
		WireDto: wiregen.WireDto{
			ProjectPath: dto.ProjectPath,
			GomodName:   dto.GomodName,
		},
		Dependency: ast.WireDependencyInjection{
			VarName:                   "httpRouter",
			TargetInjectName:          "httprouter",
			TargetInjectConstructName: "NewHttpRouter",
		},
		Import: ast.ImportSpec{
			Name: "httprouter",
			Path: fmt.Sprintf("\"%s\"", utils.GetImportPathBasedOnProjectPath(dto.Path, dto.GomodName)),
		},
	})
	if err != nil {
		return err
	}

	return nil
}

// create internal/http/http.go file
func (s *ProtocolHttpGenImpl) createProtocolInternalHttpFile(dto ProtocolDto) error {
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
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(dto.GomodName, "config")),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(dto.GomodName, "internal/protocols/http/router")),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateBaseFileString := abstractCode.GetCode()

	err = s.fs.CreateFileIfNotExists(dto.Path, fmt.Sprintf("%s.go", "http"), templateBaseFileString)
	if err != nil {
		return nil
	}

	err = s.wireGen.CreateWireEntryPoint(wiregen.WireEntryPointDto{
		WireDto: wiregen.WireDto{
			FunctionName: "InitHttpProtocol",
		},
		Import: ast.ImportSpec{
			Path: fmt.Sprintf("\"%s\"", utils.GetImportPathBasedOnProjectPath(dto.Path, dto.GomodName)),
		},
		Return: &ast.FunctionReturnSpecList{
			&ast.FunctionReturnSpec{
				IsPointer: true,
				IsStruct:  true,
				LibName:   "http",
				DataType:  "HttpImpl",
				Return:    "HttpImpl",
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

// create internal/http directory and all of the assets
func (s *ProtocolHttpGenImpl) CreateProtocolInternalHttp(dto ProtocolDto) error {
	var err error

	err = s.createProtocolInternalHttpFile(dto)
	if err != nil {
		return err
	}

	err = s.createInternalHttpErrorDir(dto)
	if err != nil {
		return err
	}

	err = s.createInternalHttpMiddlewareDir(dto)
	if err != nil {
		return err
	}

	err = s.createInternalHttpResponseDir(dto)
	if err != nil {
		return err
	}

	err = s.createInternalHttpRouteDir(dto)
	if err != nil {
		return err
	}

	return nil
}

// create src/http directory and the assets
func (s *ProtocolHttpGenImpl) CreateProtocolSrcHttpBaseFile(dto ProtocolDto) error {
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
	if !s.fs.IsFileExists(utils.ConcatDirPath(dto.Path, baseHandlerFile)) {
		s.fs.CreateFileIfNotExists(dto.Path, baseHandlerFile, templateBaseFileString)
	}

	err = s.wireGen.AddDependencyAfterCreatingModule(wiregen.WireAddModuleDto{
		WireDto: wiregen.WireDto{
			ProjectPath: dto.ProjectPath,
			GomodName:   dto.GomodName,
		},
		Dependency: ast.WireDependencyInjection{
			VarName:                   "httpHandler",
			TargetInjectName:          "httphandler",
			TargetInjectConstructName: "NewHttpHandler",
		},
		Import: ast.ImportSpec{
			Name: "httprouter",
			Path: fmt.Sprintf("\"%s\"", utils.GetImportPathBasedOnProjectPath(dto.Path, dto.GomodName)),
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (s ProtocolHttpGenImpl) CreateProtocolSrcHttpHandler(dto ProtocolDto) error {
	if exist := s.fs.IsFileExists(utils.ConcatDirPath(dto.Path, "http_handler.go")); !exist {
		s.CreateProtocolSrcHttpBaseFile(dto)
	}

	return s.emptyGen.CreateEmptyGolangFile(emptygen.EmptyDto{
		Path:        dto.Path,
		FileName:    dto.Name,
		PackageName: "http",
	})
}
