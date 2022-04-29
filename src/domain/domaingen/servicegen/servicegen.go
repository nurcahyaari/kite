package servicegen

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates"
	"github.com/nurcahyaari/kite/internal/utils/ast"
	"github.com/nurcahyaari/kite/src/domain/modulegen"
)

type ServiceGen interface {
	CreateService(dto ServiceDto) error
	CreateServiceDir(dto ServiceDto) error
	CreateServiceFile(dto ServiceDto) error
}

type ServiceGenImpl struct {
	fs        database.FileSystem
	moduleGen modulegen.ModuleGen
}

func NewServiceGen(
	fs database.FileSystem,
	moduleGen modulegen.ModuleGen,
) *ServiceGenImpl {
	return &ServiceGenImpl{
		fs:        fs,
		moduleGen: moduleGen,
	}
}

func (s ServiceGenImpl) CreateService(dto ServiceDto) error {
	s.CreateServiceDir(dto)
	return s.CreateServiceFile(dto)
}

func (s ServiceGenImpl) CreateServiceDir(dto ServiceDto) error {
	err := s.fs.CreateFolderIfNotExists(dto.Path)
	if err != nil {
		return err
	}

	return nil
}

func (s ServiceGenImpl) CreateServiceFile(dto ServiceDto) error {
	templateNew := templates.NewTemplateNewImpl("repository", "")
	templateCode, err := templateNew.Render("", nil)
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddFunction(ast.FunctionSpecList{
		&ast.FunctionSpec{
			Name: "NewService",
			Args: ast.FunctionArgList{
				&ast.FunctionArg{
					Name:     fmt.Sprintf("%sRepo", dto.DomainName),
					LibName:  fmt.Sprintf("%srepo", dto.DomainName),
					DataType: "Repository",
				},
			},
			Returns: &ast.FunctionReturnSpecList{
				&ast.FunctionReturnSpec{
					IsPointer: true,
					IsStruct:  true,
					DataType:  "ServiceImpl",
					Return:    "ServiceImpl",
				},
			},
		},
	})
	abstractCode.AddFunctionArgsToReturn(ast.FunctionReturnArgsSpec{
		FuncName:      "NewService",
		ReturnName:    "ServiceImpl",
		DataTypeKey:   fmt.Sprintf("%sRepo", dto.DomainName),
		DataTypeValue: fmt.Sprintf("%sRepo", dto.DomainName),
	})
	abstractCode.AddStructs(ast.StructSpecList{
		&ast.StructSpec{
			Name: "ServiceImpl",
		},
	})
	abstractCode.AddStructVarDecl(ast.StructArgList{
		&ast.StructArg{
			StructName: "ServiceImpl",
			Name:       fmt.Sprintf("%sRepo", dto.DomainName),
			DataType: ast.StructDtypes{
				LibName:  fmt.Sprintf("%srepo", dto.DomainName),
				TypeName: "Repository",
			},
		},
	})
	abstractCode.AddInterfaces(ast.InterfaceSpecList{
		&ast.InterfaceSpec{
			Name:       "Service",
			StructName: "ServiceImpl",
		},
	})
	abstractCode.AddImport(ast.ImportSpec{
		Name: fmt.Sprintf("%srepo", dto.DomainName),
		Path: fmt.Sprintf("\"%s/src/domains/%s/repository\"", dto.GomodName, dto.DomainName),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateBaseFileString := abstractCode.GetCode()

	err = s.fs.CreateFileIfNotExists(dto.Path, "service.go", templateBaseFileString)
	if err != nil {
		return err
	}

	return nil
}
