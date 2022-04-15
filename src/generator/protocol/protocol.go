package protocol

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/src/ast"
	"github.com/nurcahyaari/kite/src/templates/protocoltemplate/httptemplate/chitemplate"
	"github.com/nurcahyaari/kite/src/utils/fs"
)

type ProtocolGen interface {
	CreateInternalProtocolDir() error
	CreateProtocolTypeDir() error
	CreateProtocolDir() error
}

type ProtocolGenImpl struct {
	ProtocolType ProtocolType
	ProtocolPath string
	GomodName    string
}

// default protocol is http
func NewProtocolGen(
	protocol string,
	InternalPath string,
	GomodName string,
) *ProtocolGenImpl {
	protocolType := NewProtocolType(protocol)
	protocolPath := fs.ConcatDirPath(InternalPath, "protocol")
	return &ProtocolGenImpl{
		ProtocolType: protocolType,
		ProtocolPath: protocolPath,
		GomodName:    GomodName,
	}
}

func (s *ProtocolGenImpl) CreateProtocolTypeDir() error {
	dirPath := fs.ConcatDirPath(s.ProtocolPath, s.ProtocolType.ToString())
	err := fs.CreateFolderIsNotExist(dirPath)
	if err != nil {
		return err
	}

	switch s.ProtocolType {
	case Http:
		err = s.createProtocolInternalHttp(dirPath)
	}

	if err != nil {
		return err
	}

	return nil
}

// create protocol under src dir
func (s *ProtocolGenImpl) CreateProtocolDir() error {
	ProtocolTypePath := fs.ConcatDirPath(s.ProtocolPath, s.ProtocolType.ToString())
	err := fs.CreateFolderIsNotExist(ProtocolTypePath)
	if err != nil {
		return err
	}
	return nil
}

// create protocol under internal dir
func (s *ProtocolGenImpl) createProtocolInternalHttp(path string) error {
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

func (s *ProtocolGenImpl) CreateInternalProtocolDir() error {
	err := fs.CreateFolderIsNotExist(s.ProtocolPath)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProtocolGenImpl) createInternalHttpErrorDir(path string) error {
	path = fs.ConcatDirPath(path, "errors")
	err := fs.CreateFolderIsNotExist(path)
	if err != nil {
		return err
	}

	err = s.createInternalErrorHttpFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProtocolGenImpl) createInternalErrorHttpFile(path string) error {
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

	fs.CreateFileIfNotExist(path, "error.go", errorTemplateString)

	return nil
}

func (s *ProtocolGenImpl) createInternalHttpMiddlewareDir(path string) error {
	path = fs.ConcatDirPath(path, "middleware")
	err := fs.CreateFolderIsNotExist(path)
	if err != nil {
		return err
	}

	err = s.createInternalMiddlewareHttpFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProtocolGenImpl) createInternalMiddlewareHttpFile(path string) error {
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
		Path: fmt.Sprintf("\"%s\"", fs.ConcatDirPath(s.GomodName, "config")),
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
		Path: fmt.Sprintf("\"%s\"", fs.ConcatDirPath(s.GomodName, "internal/protocol/http/response")),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s\"", fs.ConcatDirPath(s.GomodName, "internal/utils/encryption")),
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

	fs.CreateFileIfNotExist(path, "jwt.go", middlewareTemplateString)

	return nil
}

func (s *ProtocolGenImpl) createInternalHttpResponseDir(path string) error {
	path = fs.ConcatDirPath(path, "response")
	err := fs.CreateFolderIsNotExist(path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpResponseFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProtocolGenImpl) createInternalHttpResponseFile(path string) error {
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
		Path: fmt.Sprintf("\"%s\"", fs.ConcatDirPath(s.GomodName, "internal/protocol/http/errors")),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}

	responseTemplateString = abstractCode.GetCode()

	fs.CreateFileIfNotExist(path, "response.go", responseTemplateString)

	return nil
}

func (s *ProtocolGenImpl) createInternalHttpRouteDir(path string) error {
	path = fs.ConcatDirPath(path, "router")
	err := fs.CreateFolderIsNotExist(path)
	if err != nil {
		return err
	}

	err = s.createInternalHttpRouteFile(path)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProtocolGenImpl) createInternalHttpRouteFile(path string) error {
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
		Path: fmt.Sprintf("\"%s\"", fs.ConcatDirPath(s.GomodName, "src/handler/http")),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateCode = abstractCode.GetCode()

	err = fs.CreateFileIfNotExist(path, "route.go", templateCode)
	if err != nil {
		return nil
	}

	return nil
}

func (s *ProtocolGenImpl) createProtocolInternalHttpFile(path string) error {
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
		Path: fmt.Sprintf("\"%s\"", fs.ConcatDirPath(s.GomodName, "config")),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s\"", fs.ConcatDirPath(s.GomodName, "internal/protocol/http/router")),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateBaseFileString := abstractCode.GetCode()

	// abstractCode = ast.NewAbstractCode(templateBaseFileString, parser.ParseComments)
	// abstractCode.AddFunctionCaller("setupRouter", ast.CallerSpec{
	// 	Func: ast.CallerFunc{
	// 		Name: ast.CallerSelecterExpr{
	// 			Name:     "p",
	// 			Selector: "HttpRouter",
	// 		},
	// 		Selector: "Router",
	// 	},
	// })
	// abstractCode.AddArgsToCallExpr(ast.CallerSpec{
	// 	Func: ast.CallerFunc{
	// 		Name: ast.CallerSelecterExpr{
	// 			Name:     "p",
	// 			Selector: "HttpRouter",
	// 		},
	// 		Selector: "Router",
	// 	},
	// 	Args: ast.CallerArgList{
	// 		&ast.CallerArg{
	// 			SelectorStmt: &ast.CallerArgSelectorStmt{
	// 				DataType: "app",
	// 			},
	// 		},
	// 	},
	// })
	// err = abstractCode.RebuildCode()
	// if err != nil {
	// 	return err
	// }
	// templateBaseFileString = abstractCode.GetCode()

	err = fs.CreateFileIfNotExist(path, fmt.Sprintf("%s.go", "http"), templateBaseFileString)
	if err != nil {
		return nil
	}

	return nil
}
