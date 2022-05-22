package repositorygen

import (
	"fmt"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/internal/utils/ast"
	"github.com/nurcahyaari/kite/src/domain/modulegen"
	"github.com/nurcahyaari/kite/src/domain/wiregen"
)

type RepositoryGen interface {
	CreateRepository(dto RepositoryDto) error
	CreateRepositoryDir(dto RepositoryDto) error
	CreateRepositoryFile(dto RepositoryDto) error
}

type RepositoryGenImpl struct {
	fs        database.FileSystem
	moduleGen modulegen.ModuleGen
	wireGen   wiregen.WireGen
}

func NewRepositoryGen(
	fs database.FileSystem,
	moduleGen modulegen.ModuleGen,
	wireGen wiregen.WireGen,
) *RepositoryGenImpl {
	return &RepositoryGenImpl{
		fs:        fs,
		moduleGen: moduleGen,
		wireGen:   wireGen,
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
	var addFuncArgs *ast.FunctionArgList
	var addFuncArgToReturn *ast.FunctionReturnArgsSpec
	var addStructVarDecl *ast.StructArgList
	var addImport *ast.ImportSpec

	addFuncArgs = &ast.FunctionArgList{
		&ast.FunctionArg{
			IsPointer: true,
			Name:      "db",
			LibName:   "database",
			DataType:  "MysqlImpl",
		},
	}
	addFuncArgToReturn = &ast.FunctionReturnArgsSpec{
		FuncName:      fmt.Sprintf("New%sRepository", utils.CapitalizeFirstLetter(dto.DomainName)),
		ReturnName:    fmt.Sprintf("%sRepositoryImpl", utils.CapitalizeFirstLetter(dto.DomainName)),
		DataTypeKey:   "db",
		DataTypeValue: "db",
	}
	addStructVarDecl = &ast.StructArgList{
		&ast.StructArg{
			StructName: fmt.Sprintf("%sRepositoryImpl", utils.CapitalizeFirstLetter(dto.DomainName)),
			IsPointer:  true,
			Name:       "db",
			DataType: ast.StructDtypes{
				LibName:  "database",
				TypeName: "MysqlImpl",
			},
		},
	}
	addImport = &ast.ImportSpec{
		Path: fmt.Sprintf("\"%s\"", utils.ConcatDirPath(dto.GomodName, "infrastructure", "database")),
	}

	err := s.moduleGen.CreateNewModule(modulegen.ModuleDto{
		PackageName:        "repository",
		FileName:           "repository",
		ModuleName:         fmt.Sprintf("%sRepository", utils.CapitalizeFirstLetter(dto.DomainName)),
		Path:               dto.Path,
		ProjectPath:        dto.ProjectPath,
		GomodName:          dto.GomodName,
		AddFuncArgs:        addFuncArgs,
		AddFuncArgToReturn: addFuncArgToReturn,
		AddStructVarDecl:   addStructVarDecl,
		AddImport:          addImport,
	})

	return err
}
