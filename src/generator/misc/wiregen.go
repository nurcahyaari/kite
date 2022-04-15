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
	AddDependencyAfterCreatingModule(importSpec ast.ImportSpec, dependency ast.WireDependencyInjection) error
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
				&ast.CallerArg{
					SelectorStmt: &ast.CallerArgSelectorStmt{
						LibName:  "http",
						DataType: "NewHttp",
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

func (s WireGenImpl) AddDependencyAfterCreatingModule(importSpec ast.ImportSpec, dependency ast.WireDependencyInjection) error {
	wirePath := fs.ConcatDirPath(s.ProjectPath, "wire.go")
	val, err := fs.ReadFile(wirePath)
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(val, parser.ParseComments)
	abstractCode.AddWireDependencyInjection(dependency)
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
						Name: dependency.VarName,
					},
				},
			},
		},
	)
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	astCode := abstractCode.GetCode()

	abstractCode = ast.NewAbstractCode(astCode, parser.ParseComments)
	abstractCode.AddImport(importSpec)
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	astCode = abstractCode.GetCode()

	return fs.ReplaceFile(s.ProjectPath, "wire.go", astCode)
}
