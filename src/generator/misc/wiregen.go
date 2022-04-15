package misc

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/src/ast"
	"github.com/nurcahyaari/kite/src/templates"
	"github.com/nurcahyaari/kite/src/utils/fs"
)

type WireGen interface {
	CreateWireFiles() error
}

type WireGenImpl struct {
	GomodName   string
	ProjectPath string
}

func NewWire(
	projectPath string,
	gomodName string,
) WireGen {
	return &WireGenImpl{
		ProjectPath: projectPath,
		GomodName:   gomodName,
	}
}

func (s WireGenImpl) CreateWireFiles() error {
	templateNew := templates.NewTemplateNewImpl("main", "")
	templateCode, err := templateNew.Render("", nil)
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddCommentOutsideFunction(ast.Comment{
		Value: "//+build wireinject",
	})
	abstractCode.AddFunction(ast.FunctionSpecList{
		&ast.FunctionSpec{
			Name: "InitHttpProtocol",
			Returns: &ast.FunctionReturnSpecList{
				&ast.FunctionReturnSpec{
					IsPointer: true,
					IsStruct:  true,
					LibName:   "http",
					DataType:  "HttpImpl",
					Return:    "HttpImpl",
				},
			},
		},
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateCode = abstractCode.GetCode()

	abstractCode = ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddFunctionCaller("InitHttpProtocol", ast.CallerSpec{
		Func: ast.CallerFunc{
			Name: ast.CallerSelecterExpr{
				Name: "wire",
			},
			Selector: "Build",
		},
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateCode = abstractCode.GetCode()

	abstractCode = ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddWireDependencyInjection(ast.WireDependencyInjection{
		VarName:                   "httpRouter",
		TargetInjectName:          "httprouter",
		TargetInjectConstructName: "NewHttpRouter",
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateCode = abstractCode.GetCode()

	abstractCode = ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddWireDependencyInjection(ast.WireDependencyInjection{
		VarName:                   "httpHandler",
		TargetInjectName:          "httphandler",
		TargetInjectConstructName: "NewHttpHandler",
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateCode = abstractCode.GetCode()

	abstractCode = ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddWireDependencyInjection(ast.WireDependencyInjection{
		VarName:                   "storages",
		TargetInjectName:          "db",
		TargetInjectConstructName: "NewMysqlClient",
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateCode = abstractCode.GetCode()

	abstractCode = ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/google/wire\"",
	})
	abstractCode.AddImport(ast.ImportSpec{
		Name: "httprouter",
		Path: fmt.Sprintf("\"%s/internal/protocol/http/router\"", s.GomodName),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Name: "httphandler",
		Path: fmt.Sprintf("\"%s/src/handler/http\"", s.GomodName),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s/internal/protocol/http\"", s.GomodName),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Name: "db",
		Path: fmt.Sprintf("\"%s/infrastructure/database\"", s.GomodName),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateCode = abstractCode.GetCode()

	abstractCode = ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddArgsToCallExpr(
		ast.CallerSpec{
			Func: ast.CallerFunc{
				Name: ast.CallerSelecterExpr{
					Name: "wire",
				},
				Selector: "Build",
			},
			Args: ast.CallerArgList{
				&ast.CallerArg{
					Ident: &ast.CallerArgIdent{
						Name: "storages",
					},
				},
				&ast.CallerArg{
					Ident: &ast.CallerArgIdent{
						Name: "httpHandler",
					},
				},
				&ast.CallerArg{
					Ident: &ast.CallerArgIdent{
						Name: "httpRouter",
					},
				},
			},
		},
	)
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateCode = abstractCode.GetCode()

	return fs.CreateFileIfNotExist(s.ProjectPath, "wire.go", templateCode)
}
