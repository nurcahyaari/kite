package module

import (
	"fmt"

	"github.com/nurcahyaari/kite/utils/fs"
)

type ModuleGen interface {
	DtoGen
	EntityGen
	RepositoryGen
	ServiceGen
	HandlerGen
	CreateSrcDir() error
	CreateBaseModuleDir() error
	CreateNewModule() error
}

type ModuleGenImpl struct {
	// directory path is the path of the directory that store module
	DirectoryPath string
	// the base of module path
	BaseModulePath string
	// module name is the name of the module
	ModuleName string
	// module path is the place of the module
	ModulePath string
	// BaseHandlerPath  base of handler path
	BaseHandlerPath string
	// path of the project
	ProjectPath string
	// Derived module
	*DtoGenImpl
	*EntityGenImpl
	*RepositoryGenImpl
	*ServiceGenImpl
	*HandlerGenImpl
}

func NewModuleGen(projectPath string, moduleName string, gomodName string) ModuleGen {
	directoryPath := fs.ConcatDirPath(
		projectPath, "src",
	)
	baseModulePath := fs.ConcatDirPath(directoryPath, "module")

	modulePath := fs.ConcatDirPath(
		baseModulePath,
		moduleName,
	)

	return &ModuleGenImpl{
		DirectoryPath:     directoryPath,
		BaseModulePath:    baseModulePath,
		ModulePath:        modulePath,
		ModuleName:        moduleName,
		ProjectPath:       projectPath,
		DtoGenImpl:        NewDtoGen(modulePath),
		EntityGenImpl:     NewEntityGen(modulePath),
		RepositoryGenImpl: NewRepositoryGen(moduleName, modulePath, gomodName),
		ServiceGenImpl:    NewServiceGen(moduleName, modulePath, gomodName),
		HandlerGenImpl:    NewHandlerGen(directoryPath, moduleName, gomodName),
	}
}

func (s *ModuleGenImpl) CreateSrcDir() error {
	err := fs.CreateFolderIsNotExist(s.DirectoryPath)
	if err != nil {
		return err
	}

	return nil
}

func (s *ModuleGenImpl) CreateBaseModuleDir() error {
	// validate is project exist
	if !fs.IsFolderExist(s.ProjectPath) {
		return fmt.Errorf("%s project path is not exist", s.ProjectPath)
	}

	moduleDir := fs.ConcatDirPath(s.DirectoryPath, "module")
	fs.CreateFolderIsNotExist(moduleDir)

	return nil
}

func (s *ModuleGenImpl) CreateNewModule() error {
	err := fs.CreateFolderIsNotExist(s.ModulePath)
	if err != nil {
		return err
	}

	return nil
}
