package modulegen

import (
	"fmt"
	"go/parser"
	"strings"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/internal/utils/ast"
	"github.com/nurcahyaari/kite/src/domain/wiregen"
)

// create module, what is module
// all of these package is based on module.
// dependency in here called as a module
type ModuleGen interface {
	CreateNewModule(dto ModuleDto) error
	BuildModuleTemplate(dto ModuleDto) (string, error)
}

type ModuleGenImpl struct {
	fs      database.FileSystem
	wireGen wiregen.WireGen
}

func NewModuleGen(fs database.FileSystem, wireGen wiregen.WireGen) *ModuleGenImpl {
	return &ModuleGenImpl{fs: fs, wireGen: wireGen}
}

// CreateNewModule to create new module
func (s ModuleGenImpl) CreateNewModule(dto ModuleDto) error {
	// get package name
	filesPath, err := utils.GetGoFilesInPath(dto.Path)
	if err != nil {
		return err
	}

	if len(filesPath) > 0 {
		fileValue, err := utils.ReadFile(filesPath[0])
		if err != nil {
			return err
		}

		abstractCode := ast.NewAbstractCode(fileValue, parser.ParseComments)
		dto.PackageName = abstractCode.GetPackageName()
	} else {
		dto.PackageName = utils.GetLastDirPath(dto.Path)
	}

	template, err := s.BuildModuleTemplate(dto)
	if err != nil {
		return err
	}

	err = s.fs.CreateFileIfNotExists(dto.Path, fmt.Sprintf("%s.go", dto.FileName), template)
	if err != nil {
		return err
	}

	importSpec := ast.ImportSpec{}
	dependencySpec := ast.WireDependencyInjection{
		VarName:                   dto.ModuleName,
		TargetInjectConstructName: fmt.Sprintf("New%s", utils.CapitalizeFirstLetter(dto.ModuleName)),
		InterfaceName:             utils.CapitalizeFirstLetter(dto.ModuleName),
		StructName:                fmt.Sprintf("%sImpl", utils.CapitalizeFirstLetter(dto.ModuleName)),
	}
	if dto.PackageName != "main" {
		importSpec = ast.ImportSpec{
			Name: strings.ToLower(dto.ModuleName),
			Path: fmt.Sprintf("\"%s\"", utils.GetImportPathBasedOnProjectPath(dto.Path, dto.GomodName)),
		}
		dependencySpec.TargetInjectName = strings.ToLower(dto.ModuleName)
		dependencySpec.InterfaceLib = strings.ToLower(dto.ModuleName)
		dependencySpec.StructLib = strings.ToLower(dto.ModuleName)
	}

	err = s.wireGen.AddDependencyAfterCreatingModule(wiregen.WireAddModuleDto{
		WireDto: wiregen.WireDto{
			ProjectPath: dto.ProjectPath,
			GomodName:   dto.GomodName,
		},
		Dependency: dependencySpec,
		Import:     importSpec,
	})

	return err
}

// BuildModuleTemplate will build base file of the module, such as interface struct and the construct
func (s ModuleGenImpl) BuildModuleTemplate(dto ModuleDto) (string, error) {
	templateNew := templates.NewTemplateNewImpl(dto.PackageName, "")
	templateCode, err := templateNew.Render("", nil)
	if err != nil {
		return "", err
	}

	abstractCode := ast.NewAbstractCode(templateCode, parser.ParseComments)
	abstractCode.AddFunction(ast.FunctionSpecList{
		&ast.FunctionSpec{
			Name: fmt.Sprintf("New%s", utils.CapitalizeFirstLetter(dto.ModuleName)),
			Args: *dto.AddFuncArgs,
			Returns: &ast.FunctionReturnSpecList{
				&ast.FunctionReturnSpec{
					IsPointer: true,
					IsStruct:  true,
					DataType:  fmt.Sprintf("%sImpl", utils.CapitalizeFirstLetter(dto.ModuleName)),
					Return:    fmt.Sprintf("%sImpl", utils.CapitalizeFirstLetter(dto.ModuleName)),
				},
			},
		},
	})
	if dto.AddFuncArgToReturn != nil {
		abstractCode.AddFunctionArgsToReturn(*dto.AddFuncArgToReturn)
	}
	abstractCode.AddStructs(ast.StructSpecList{
		&ast.StructSpec{
			Name: fmt.Sprintf("%sImpl", utils.CapitalizeFirstLetter(dto.ModuleName)),
		},
	})
	if dto.AddStructVarDecl != nil {
		abstractCode.AddStructVarDecl(*dto.AddStructVarDecl)
	}
	abstractCode.AddInterfaces(ast.InterfaceSpecList{
		&ast.InterfaceSpec{
			Name:       utils.CapitalizeFirstLetter(dto.ModuleName),
			StructName: fmt.Sprintf("%sImpl", utils.CapitalizeFirstLetter(dto.ModuleName)),
		},
	})
	if dto.AddImport != nil {
		abstractCode.AddImport(*dto.AddImport)
	}
	err = abstractCode.RebuildCode()
	if err != nil {
		return "", err
	}
	templateBaseFileString := abstractCode.GetCode()

	return templateBaseFileString, nil
}
