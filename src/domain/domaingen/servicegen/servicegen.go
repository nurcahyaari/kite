package servicegen

import (
	"fmt"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/internal/utils/ast"
	"github.com/nurcahyaari/kite/src/domain/modulegen"
	"github.com/nurcahyaari/kite/src/domain/protocolgen"
	"github.com/nurcahyaari/kite/src/domain/wiregen"
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
	wireGen     wiregen.WireGen
}

func NewServiceGen(
	fs database.FileSystem,
	moduleGen modulegen.ModuleGen,
	protocolGen protocolgen.ProtocolGen,
	wireGen wiregen.WireGen,
) *ServiceGenImpl {
	return &ServiceGenImpl{
		fs:          fs,
		moduleGen:   moduleGen,
		protocolGen: protocolGen,
		wireGen:     wireGen,
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
	var addFuncArgs *ast.FunctionArgList
	var addFuncArgToReturn *ast.FunctionReturnArgsSpec
	var addStructVarDecl *ast.StructArgList
	var addImport *ast.ImportSpec
	if dto.IsInjectRepo {
		addFuncArgs = &ast.FunctionArgList{
			&ast.FunctionArg{
				Name:     fmt.Sprintf("%sRepo", dto.DomainName),
				LibName:  fmt.Sprintf("%srepo", dto.DomainName),
				DataType: fmt.Sprintf("%sRepository", utils.CapitalizeFirstLetter(dto.DomainName)),
			},
		}
		addFuncArgToReturn = &ast.FunctionReturnArgsSpec{
			FuncName:      fmt.Sprintf("New%sService", utils.CapitalizeFirstLetter(dto.DomainName)),
			ReturnName:    fmt.Sprintf("%sServiceImpl", utils.CapitalizeFirstLetter(dto.DomainName)),
			DataTypeKey:   fmt.Sprintf("%sRepo", dto.DomainName),
			DataTypeValue: fmt.Sprintf("%sRepo", dto.DomainName),
		}
		addStructVarDecl = &ast.StructArgList{
			&ast.StructArg{
				StructName: fmt.Sprintf("%sServiceImpl", utils.CapitalizeFirstLetter(dto.DomainName)),
				Name:       fmt.Sprintf("%sRepo", dto.DomainName),
				DataType: ast.StructDtypes{
					LibName:  fmt.Sprintf("%srepo", dto.DomainName),
					TypeName: fmt.Sprintf("%sRepository", utils.CapitalizeFirstLetter(dto.DomainName)),
				},
			},
		}
		addImport = &ast.ImportSpec{
			Name: fmt.Sprintf("%srepo", dto.DomainName),
			Path: fmt.Sprintf("\"%s/src/domains/%s/repository\"", dto.GomodName, dto.DomainName),
		}
	}

	err := s.moduleGen.CreateNewModule(modulegen.ModuleDto{
		PackageName:        "service",
		FileName:           "service",
		ModuleName:         fmt.Sprintf("%sService", utils.CapitalizeFirstLetter(dto.DomainName)),
		Path:               dto.Path,
		ProjectPath:        dto.ProjectPath,
		GomodName:          dto.GomodName,
		AddFuncArgs:        addFuncArgs,
		AddFuncArgToReturn: addFuncArgToReturn,
		AddStructVarDecl:   addStructVarDecl,
		AddImport:          addImport,
	})

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

	return err
}
