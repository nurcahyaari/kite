package protocolgen

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/internal/utils/ast"
	"github.com/nurcahyaari/kite/src/domain/protocolgen/protocoltype"
)

type ProtocolGen interface {
	createProtocolDir(dto ProtocolDto) (string, error)
	CreateProtocolInternalType(dto ProtocolDto) error
	CreateProtocolSrcBaseFile(dto ProtocolDto) error
	CreateProtocolSrcHandler(dto ProtocolDto) error
	InjectDomainServiceIntoHandler(dto ProtocolDto) error
}

type ProtocolGenImpl struct {
	fs           database.FileSystem
	protocolType protocoltype.ProtocolType
}

// default protocol is http
func NewProtocolGen(
	fs database.FileSystem,
	protocolType protocoltype.ProtocolType,
) *ProtocolGenImpl {
	return &ProtocolGenImpl{
		fs:           fs,
		protocolType: protocolType,
	}
}

func (s *ProtocolGenImpl) createProtocolDir(dto ProtocolDto) (string, error) {
	dirPath := utils.ConcatDirPath(dto.Path, dto.ProtocolType.ToString())
	err := s.fs.CreateFolderIfNotExists(dto.Path)
	if err != nil {
		return "", err
	}
	return dirPath, nil
}

func (s *ProtocolGenImpl) CreateProtocolInternalType(dto ProtocolDto) error {
	dirPath, err := s.createProtocolDir(dto)
	if err != nil {
		return err
	}

	if !s.fs.IsFolderExists(dirPath) {
		s.fs.CreateFolderIfNotExists(dirPath)
	}

	protocolDto := protocoltype.ProtocolDto{
		Name:        dto.Name,
		GomodName:   dto.GomodName,
		Path:        dirPath,
		ProjectPath: dto.ProjectPath,
	}

	switch dto.ProtocolType {
	case Http:
		err = s.protocolType.CreateProtocolInternalHttp(protocolDto)
	}
	if err != nil {
		return err
	}

	return nil
}

// create protocol directory. under src or under internal
func (s *ProtocolGenImpl) CreateProtocolSrcBaseFile(dto ProtocolDto) error {
	var err error

	if dto.ProtocolType.NotEmpty() {
		path := utils.ConcatDirPath(dto.Path, dto.ProtocolType.ToString())

		if !s.fs.IsFolderExists(path) {
			s.fs.CreateFolderIfNotExists(path)
		}

		protocolDto := protocoltype.ProtocolDto{
			Name:        dto.Name,
			GomodName:   dto.GomodName,
			Path:        path,
			ProjectPath: dto.ProjectPath,
		}

		switch dto.ProtocolType {
		case Http:
			err = s.protocolType.CreateProtocolSrcHttpBaseFile(protocolDto)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *ProtocolGenImpl) CreateProtocolSrcHandler(dto ProtocolDto) error {
	var err error

	if dto.ProtocolType.NotEmpty() {
		path := utils.ConcatDirPath(dto.Path, dto.ProtocolType.ToString())

		if !s.fs.IsFolderExists(path) {
			s.fs.CreateFolderIfNotExists(path)
		}

		protocolDto := protocoltype.ProtocolDto{
			Name:        dto.Name,
			GomodName:   dto.GomodName,
			Path:        path,
			ProjectPath: dto.ProjectPath,
		}

		switch dto.ProtocolType {
		case Http:
			err = s.protocolType.CreateProtocolSrcHttpHandler(protocolDto)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *ProtocolGenImpl) InjectDomainServiceIntoHandler(dto ProtocolDto) error {
	servicePath := utils.ConcatDirPath(dto.GomodName, "src", "domains", dto.DomainName, "service")
	handlerFileName := fmt.Sprintf("%s_handler.go", dto.ProtocolType.ToString())
	handlerDirPath := utils.ConcatDirPath(dto.Path, dto.ProtocolType.ToString())
	handlerFilePath := utils.ConcatDirPath(handlerDirPath, handlerFileName)
	val, err := s.fs.ReadFile(handlerFilePath)
	if err != nil {
		return err
	}

	abstractCode := ast.NewAbstractCode(val, parser.ParseComments)

	importAlias := fmt.Sprintf("%ssvc", dto.DomainName)
	abstractCode.AddImport(ast.ImportSpec{
		Name: importAlias,
		Path: fmt.Sprintf("\"%s\"", servicePath),
	})
	abstractCode.AddFunctionArgs(ast.FunctionSpec{
		Name: "NewHttpHandler",
		Args: ast.FunctionArgList{
			&ast.FunctionArg{
				Name:     fmt.Sprintf("%sSvc", dto.DomainName),
				LibName:  importAlias,
				DataType: "Service",
			},
		},
	})
	abstractCode.AddStructVarDecl(ast.StructArgList{
		&ast.StructArg{
			StructName: "HttpHandlerImpl",
			Name:       fmt.Sprintf("%sSvc", dto.DomainName),
			DataType: ast.StructDtypes{
				LibName:  importAlias,
				TypeName: "Service",
			},
			IsPointer: false,
		},
	})
	abstractCode.AddFunctionArgsToReturn(ast.FunctionReturnArgsSpec{
		FuncName:      "NewHttpHandler",
		ReturnName:    "HttpHandlerImpl",
		DataTypeKey:   fmt.Sprintf("%sSvc", dto.DomainName),
		DataTypeValue: fmt.Sprintf("%sSvc", dto.DomainName),
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return err
	}
	newCode := abstractCode.GetCode()

	err = s.fs.ReplaceFile(handlerDirPath, handlerFileName, newCode)

	return err
}
