package modulegen

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/internal/templates"
	"github.com/nurcahyaari/kite/internal/utils/ast"
)

// create module, what is module
// all of these package is based on module.
// dependency in here called as a module
type ModuleGen interface {
	BuildModuleTemplate(dto ModuleDto) (string, error)
}

type ModuleGenImpl struct {
}

func NewModuleGen() *ModuleGenImpl {
	return &ModuleGenImpl{}
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
			Name: fmt.Sprintf("New%s", dto.ModuleName),
			Returns: &ast.FunctionReturnSpecList{
				&ast.FunctionReturnSpec{
					IsPointer: true,
					IsStruct:  true,
					DataType:  fmt.Sprintf("%sImpl", dto.ModuleName),
					Return:    fmt.Sprintf("%sImpl", dto.ModuleName),
				},
			},
		},
	})
	abstractCode.AddStructs(ast.StructSpecList{
		&ast.StructSpec{
			Name: fmt.Sprintf("%sImpl", dto.ModuleName),
		},
	})
	abstractCode.AddInterfaces(ast.InterfaceSpecList{
		&ast.InterfaceSpec{
			Name:       dto.ModuleName,
			StructName: fmt.Sprintf("%sImpl", dto.ModuleName),
		},
	})
	err = abstractCode.RebuildCode()
	if err != nil {
		return "", err
	}
	templateBaseFileString := abstractCode.GetCode()

	return templateBaseFileString, nil
}

// type ModuleGen interface {
// 	DtoGen
// 	EntityGen
// 	RepositoryGen
// 	ServiceGen
// 	HandlerGen
// 	// CreateSrcDir() error
// 	// CreateBaseModuleDir() error
// 	CreateNewModule() error
// 	// AppendModuleToWire() error
// }

// type ModuleGenImpl struct {
// 	// directory path is the path of the directory that store module
// 	DirectoryPath string
// 	// the base of module path
// 	BaseModulePath string
// 	// module name is the name of the module
// 	ModuleName string
// 	// module path is the place of the module
// 	ModulePath string
// 	// BaseHandlerPath  base of handler path
// 	BaseHandlerPath string
// 	// path of the project
// 	ProjectPath string
// 	// Derived module
// 	GomodName string
// 	*DtoGenImpl
// 	*EntityGenImpl
// 	*RepositoryGenImpl
// 	*ServiceGenImpl
// 	*HandlerGenImpl
// }

// func NewModuleGen(projectPath string, moduleName string, gomodName string) *ModuleGenImpl {
// 	directoryPath := fs.ConcatDirPath(
// 		projectPath, "src",
// 	)
// 	baseModulePath := fs.ConcatDirPath(directoryPath, "module")

// 	modulePath := fs.ConcatDirPath(
// 		baseModulePath,
// 		moduleName,
// 	)

// 	return &ModuleGenImpl{
// 		DirectoryPath:     directoryPath,
// 		BaseModulePath:    baseModulePath,
// 		ModulePath:        modulePath,
// 		ModuleName:        moduleName,
// 		ProjectPath:       projectPath,
// 		GomodName:         gomodName,
// 		DtoGenImpl:        NewDtoGen(modulePath),
// 		EntityGenImpl:     NewEntityGen(modulePath),
// 		RepositoryGenImpl: NewRepositoryGen(moduleName, modulePath, gomodName),
// 		ServiceGenImpl:    NewServiceGen(moduleName, modulePath, gomodName),
// 		HandlerGenImpl:    NewHandlerGen(directoryPath, moduleName, gomodName),
// 	}
// }

// // func (s *ModuleGenImpl) CreateSrcDir() error {
// // 	err := fs.CreateFolderIsNotExist(s.DirectoryPath)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	return nil
// // }

// // func (s *ModuleGenImpl) CreateBaseModuleDir() error {
// // 	// validate is project exist
// // 	if !fs.IsFolderExist(s.ProjectPath) {
// // 		return fmt.Errorf("%s project path is not exist", s.ProjectPath)
// // 	}

// // 	moduleDir := fs.ConcatDirPath(s.DirectoryPath, "module")
// // 	fs.CreateFolderIsNotExist(moduleDir)

// // 	return nil
// // }

// func (s *ModuleGenImpl) CreateNewModule() error {
// 	err := fs.CreateFolderIsNotExist(s.ModulePath)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // func (s *ModuleGenImpl) AppendModuleToWire() error {
// // 	// append repo
// // 	wireGen := misc.NewWire(s.ProjectPath, "")
// // 	wireGen.AddDependencyAfterCreatingModule(
// // 		ast.ImportSpec{
// // 			Name: fmt.Sprintf("%srepo", s.ModuleName),
// // 			Path: fmt.Sprintf("\"%s\"", fs.ConcatDirPath(s.GomodName, "src/module", s.ModuleName, "repository")),
// // 		},
// // 		ast.WireDependencyInjection{
// // 			VarName:                   fmt.Sprintf("%sRepo", s.ModuleName),
// // 			TargetInjectName:          fmt.Sprintf("%srepo", s.ModuleName),
// // 			TargetInjectConstructName: "NewRepository",
// // 			InterfaceLib:              fmt.Sprintf("%srepo", s.ModuleName),
// // 			InterfaceName:             "Repository",
// // 			StructLib:                 fmt.Sprintf("%srepo", s.ModuleName),
// // 			StructName:                "RepositoryImpl",
// // 		},
// // 	)

// // 	// append service
// // 	wireGen.AddDependencyAfterCreatingModule(
// // 		ast.ImportSpec{
// // 			Name: fmt.Sprintf("%ssvc", s.ModuleName),
// // 			Path: fmt.Sprintf("\"%s\"", fs.ConcatDirPath(s.GomodName, "src/module", s.ModuleName, "service")),
// // 		},
// // 		ast.WireDependencyInjection{
// // 			VarName:                   fmt.Sprintf("%sSvc", s.ModuleName),
// // 			TargetInjectName:          fmt.Sprintf("%ssvc", s.ModuleName),
// // 			TargetInjectConstructName: "NewService",
// // 			InterfaceLib:              fmt.Sprintf("%ssvc", s.ModuleName),
// // 			InterfaceName:             "Service",
// // 			StructLib:                 fmt.Sprintf("%ssvc", s.ModuleName),
// // 			StructName:                "ServiceImpl",
// // 		},
// // 	)

// // 	return nil
// // }
