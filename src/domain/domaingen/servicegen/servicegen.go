package servicegen

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/internal/utils/ast"
	"github.com/nurcahyaari/kite/src/domain/modulegen"
	"github.com/nurcahyaari/kite/src/domain/protocolgen"
)

type ServiceGen interface {
	CreateService(dto ServiceDto) error
	CreateServiceDir(dto ServiceDto) error
	CreateServiceFile(dto ServiceDto) error
}

type ServiceGenImpl struct {
	fs          database.FileSystem
	moduleGen   modulegen.ModuleGen
	protocolGen protocolgen.ProtocolGen
}

func NewServiceGen(
	fs database.FileSystem,
	moduleGen modulegen.ModuleGen,
	protocolGen protocolgen.ProtocolGen,
) *ServiceGenImpl {
	return &ServiceGenImpl{
		fs:          fs,
		moduleGen:   moduleGen,
		protocolGen: protocolGen,
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
	addFuncArgs := ast.FunctionArgList{}
	addFuncArgToReturn := ast.FunctionReturnArgsSpec{}
	addStructVarDecl := ast.StructArgList{}
	addImport := ast.ImportSpec{}
	if dto.IsInjectRepo {
		addFuncArgs = ast.FunctionArgList{
			&ast.FunctionArg{
				Name:     fmt.Sprintf("%sRepo", dto.DomainName),
				LibName:  fmt.Sprintf("%srepo", dto.DomainName),
				DataType: "Repository",
			},
		}
		addFuncArgToReturn = ast.FunctionReturnArgsSpec{
			FuncName:      "NewService",
			ReturnName:    "ServiceImpl",
			DataTypeKey:   fmt.Sprintf("%sRepo", dto.DomainName),
			DataTypeValue: fmt.Sprintf("%sRepo", dto.DomainName),
		}
		addStructVarDecl = ast.StructArgList{
			&ast.StructArg{
				StructName: "ServiceImpl",
				Name:       fmt.Sprintf("%sRepo", dto.DomainName),
				DataType: ast.StructDtypes{
					LibName:  fmt.Sprintf("%srepo", dto.DomainName),
					TypeName: "Repository",
				},
			},
		}
		addImport = ast.ImportSpec{
			Name: fmt.Sprintf("%srepo", dto.DomainName),
			Path: fmt.Sprintf("\"%s/src/domains/%s/repository\"", dto.GomodName, dto.DomainName),
		}
	}
	abstractCode.AddFunction(ast.FunctionSpecList{
		&ast.FunctionSpec{
			Name: "NewService",
			Args: addFuncArgs,
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
	abstractCode.AddFunctionArgsToReturn(addFuncArgToReturn)
	abstractCode.AddStructs(ast.StructSpecList{
		&ast.StructSpec{
			Name: "ServiceImpl",
		},
	})
	abstractCode.AddStructVarDecl(addStructVarDecl)
	abstractCode.AddInterfaces(ast.InterfaceSpecList{
		&ast.InterfaceSpec{
			Name:       "Service",
			StructName: "ServiceImpl",
		},
	})
	abstractCode.AddImport(addImport)
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	templateBaseFileString := abstractCode.GetCode()

	err = s.fs.CreateFileIfNotExists(dto.Path, "service.go", templateBaseFileString)
	if err != nil {
		return err
	}

	// inject service into handlers
	if dto.ProjectPath != "" && dto.IsInjectToHandler {
		handlerPath := utils.ConcatDirPath(dto.ProjectPath, "src", "handlers")
		handlerList, err := s.fs.ReadFolderList(handlerPath)
		if err != nil {
			return err
		}

		for _, h := range handlerList {
			protocolGenDto := protocolgen.ProtocolDto{
				Path:         handlerPath,
				ProtocolType: protocolgen.NewProtocolType(h),
				DomainName:   dto.DomainName,
				GomodName:    dto.GomodName,
			}

			err = s.protocolGen.InjectDomainServiceIntoHandler(protocolGenDto)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
