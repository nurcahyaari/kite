package misc

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/internal/utils/ast"
)

type WireGen interface {
	CreateWireFiles(dto MiscDto) error
	AddDependencyAfterCreatingModule(dto MiscDto, importSpec ast.ImportSpec, dependency ast.WireDependencyInjection) error
}

type WireGenImpl struct {
	fs database.FileSystem
}

func NewWire(
	fs database.FileSystem,
	// projectPath string,
	// gomodName string,
) *WireGenImpl {
	return &WireGenImpl{
		fs: fs,
		// ProjectPath: projectPath,
		// GomodName:   gomodName,
		// fs:          database.NewFileSystem(projectPath),
	}
}

func (s WireGenImpl) CreateWireFiles(dto MiscDto) error {
	templateNew := templates.NewTemplateNewImpl("main", "")
	templateCode, err := templateNew.Render("", nil)
	if err != nil {
		return err
	}

	// TODO: remove this
	// since I don't know how to add new line with ast standard lib, so I use this
	templateCode = fmt.Sprintf("\n%s", templateCode)

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
		Path: fmt.Sprintf("\"%s/internal/protocol/http/router\"", dto.GomodName),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Name: "httphandler",
		Path: fmt.Sprintf("\"%s/src/handler/http\"", dto.GomodName),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s/internal/protocol/http\"", dto.GomodName),
	})
	abstractCode.AddImport(ast.ImportSpec{
		Name: "db",
		Path: fmt.Sprintf("\"%s/infrastructure/database\"", dto.GomodName),
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

	return s.fs.CreateFileIfNotExists(dto.ProjectPath, "wire.go", templateCode)
}

func (s WireGenImpl) AddDependencyAfterCreatingModule(dto MiscDto, importSpec ast.ImportSpec, dependency ast.WireDependencyInjection) error {
	wirePath := utils.ConcatDirPath(dto.ProjectPath, "wire.go")
	val, err := utils.ReadFile(wirePath)
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

	return s.fs.ReplaceFile(dto.ProjectPath, "wire.go", astCode)
}
