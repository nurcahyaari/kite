package module

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/src/ast"
	"github.com/nurcahyaari/kite/src/templates"
	"github.com/nurcahyaari/kite/src/utils/fs"
)

type RepositoryGen interface {
	CreateRepositoryDir() error
	CreateRepositoryFile() error
}

type RepositoryGenImpl struct {
	RepositoryPath string
	ModuleName     string
	GomodName      string
}

func NewRepositoryGen(moduleName, modulePath, gomodName string) *RepositoryGenImpl {
	RepositoryPath := fs.ConcatDirPath(modulePath, "repository")
	return &RepositoryGenImpl{
		RepositoryPath: RepositoryPath,
		ModuleName:     moduleName,
		GomodName:      gomodName,
	}
}

func (s *RepositoryGenImpl) CreateRepositoryDir() error {
	err := fs.CreateFolderIsNotExist(s.RepositoryPath)
	if err != nil {
		return err
	}

	return nil
}

func (s *RepositoryGenImpl) CreateRepositoryFile() error {
	templateNew := templates.NewTemplateNewImpl("repository", "")
	templateCode, err := templateNew.Render("", nil)
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddFunction(ast.FunctionSpecList{
		&ast.FunctionSpec{
			Name: "NewRepository",
			Args: ast.FunctionArgList{
				&ast.FunctionArg{
					IsPointer: true,
					Name:      "db",
					LibName:   "infrastructure",
					DataType:  "MysqlImpl",
				},
			},
			Returns: &ast.FunctionReturnSpecList{
				&ast.FunctionReturnSpec{
					IsPointer: true,
					IsStruct:  true,
					DataType:  "RepositoryImpl",
					Return:    "RepositoryImpl",
				},
			},
		},
	})
	abstractCode.AddFunctionArgsToReturn(ast.FunctionReturnArgsSpec{
		FuncName:      "NewRepository",
		ReturnName:    "RepositoryImpl",
		DataTypeKey:   "db",
		DataTypeValue: "db",
	})
	abstractCode.AddStructs(ast.StructSpecList{
		&ast.StructSpec{
			Name: "RepositoryImpl",
		},
	})
	abstractCode.AddStructVarDecl(ast.StructArgList{
		&ast.StructArg{
			StructName: "RepositoryImpl",
			IsPointer:  true,
			Name:       "db",
			DataType: ast.StructDtypes{
				LibName:  "infrastructure",
				TypeName: "MysqlImpl",
			},
		},
	})
	abstractCode.AddInterfaces(ast.InterfaceSpecList{
		&ast.InterfaceSpec{
			Name:       "Repository",
			StructName: "RepositoryImpl",
		},
	})
	abstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s\"", fs.ConcatDirPath(s.GomodName, "infrastructure")),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateBaseFileString := abstractCode.GetCode()

	err = fs.CreateFileIfNotExist(s.RepositoryPath, "repository.go", templateBaseFileString)
	if err != nil {
		return err
	}

	return nil
}
