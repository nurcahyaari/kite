package wiregen

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/internal/utils/ast"
	"gopkg.in/guregu/null.v4"
)

type WireGen interface {
	CreateWireFiles(dto WireDto) error
	CreateWireEntryPoint(dto WireEntryPointDto) error
	AddDependencyAfterCreatingModule(dto WireAddModuleDto) error
}

type WireGenImpl struct {
	fs database.FileSystem
}

func NewWire(
	fs database.FileSystem,
) *WireGenImpl {
	return &WireGenImpl{
		fs: fs,
	}
}

func (s WireGenImpl) CreateWireFiles(dto WireDto) error {
	templateNew := templates.NewTemplateNewImpl("main", "")
	templateCode, err := templateNew.Render("", nil)
	if err != nil {
		return err
	}

	// TODO: remove this
	// since I don't know how to add new line with ast standard lib, so I use this
	templateCode = fmt.Sprintf("\n%s", templateCode)

	abstractCode := ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/google/wire\"",
	})
	abstractCode.AddCommentOutsideFunction(ast.Comment{
		Value: "//+build wireinject",
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateCode = abstractCode.GetCode()

	return s.fs.CreateFileIfNotExists(dto.ProjectPath, "wire.go", templateCode)
}

func (s WireGenImpl) CreateWireEntryPoint(dto WireEntryPointDto) error {
	wirePath := utils.ConcatDirPath(dto.ProjectPath, "wire.go")
	val, err := utils.ReadFile(wirePath)
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(val, parser.ParseComments)
	abstractCode.AddImport(dto.Import)
	abstractCode.AddFunction(ast.FunctionSpecList{
		&ast.FunctionSpec{
			Name:    dto.FunctionName,
			Returns: dto.Return,
		},
	})
	abstractCode.AddFunctionCaller(dto.FunctionName, ast.CallerSpec{
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
	astCode := abstractCode.GetCode()

	return s.fs.ReplaceFile(dto.ProjectPath, "wire.go", astCode)
}

func (s WireGenImpl) AddDependencyAfterCreatingModule(dto WireAddModuleDto) error {
	wirePath := utils.ConcatDirPath(dto.ProjectPath, "wire.go")
	val, err := utils.ReadFile(wirePath)
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(val, parser.ParseComments)
	abstractCode.AddWireDependencyInjection(dto.Dependency)
	funcName := null.String{}
	if dto.FunctionName != "" {
		funcName = null.StringFrom(dto.FunctionName)
	}
	abstractCode.AddArgsToCallExpr(
		funcName,
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
						Name: dto.Dependency.VarName,
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
	abstractCode.AddImport(dto.Import)
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	astCode = abstractCode.GetCode()

	return s.fs.ReplaceFile(dto.ProjectPath, "wire.go", astCode)
}
