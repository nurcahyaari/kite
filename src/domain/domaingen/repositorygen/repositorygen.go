package repositorygen

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/internal/utils/ast"
	"github.com/nurcahyaari/kite/src/domain/modulegen"
)

type RepositoryGen interface {
	CreateRepository(dto RepositoryDto) error
	CreateRepositoryDir(dto RepositoryDto) error
	CreateRepositoryFile(dto RepositoryDto) error
}

type RepositoryGenImpl struct {
	fs        database.FileSystem
	moduleGen modulegen.ModuleGen
}

func NewRepositoryGen(
	fs database.FileSystem,
	moduleGen modulegen.ModuleGen,
) *RepositoryGenImpl {
	return &RepositoryGenImpl{
		fs:        fs,
		moduleGen: moduleGen,
	}
}

func (s RepositoryGenImpl) CreateRepository(dto RepositoryDto) error {
	s.CreateRepositoryDir(dto)
	return s.CreateRepositoryFile(dto)
}

func (s RepositoryGenImpl) CreateRepositoryDir(dto RepositoryDto) error {
	err := s.fs.CreateFolderIfNotExists(dto.Path)
	if err != nil {
		return err
	}

	return nil
}

func (s RepositoryGenImpl) CreateRepositoryFile(dto RepositoryDto) error {
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
					LibName:   "database",
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
				LibName:  "database",
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
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(dto.GomodName, "infrastructure", "database")),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateBaseFileString := abstractCode.GetCode()

	err = s.fs.CreateFileIfNotExists(dto.Path, "repository.go", templateBaseFileString)
	if err != nil {
		return err
	}

	return nil
}
