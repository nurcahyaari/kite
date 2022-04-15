package generator

import (
	"fmt"

	"github.com/nurcahyaari/kite/src/generator/module"
	"github.com/nurcahyaari/kite/src/utils/fs"
	"github.com/nurcahyaari/kite/src/utils/logger"
)

type Module interface {
	CreateNewModule() error
}

// module.go is usefull to create a new module
type ModuleImpl struct {
	// The name of the module
	Name string

	Info ProjectInfo
}

func NewModule(moduleName string, info ProjectInfo) Module {
	return &ModuleImpl{
		Name: moduleName,
		Info: info,
	}
}

func (s *ModuleImpl) CreateNewModule() error {
	logger.Info(fmt.Sprintf("Creating %s file... ", s.Name))
	moduleDirParentPath := fs.ConcatDirPath(s.Info.ProjectPath, "src")
	moduleGen := module.NewModuleGen(s.Info.ProjectPath, s.Name, s.Info.GoModName)

	if !fs.IsFolderExist(moduleDirParentPath) {
		err := moduleGen.CreateSrcDir()
		if err != nil {
			return err
		}
	}

	moduleGen.CreateBaseModuleDir()

	err := moduleGen.CreateNewModule()
	if err != nil {
		return err
	}

	err = moduleGen.CreateDtoDir()
	if err != nil {
		return err
	}

	err = moduleGen.CreateEntityDir()
	if err != nil {
		return err
	}

	err = moduleGen.CreateRepositoryDir()
	if err != nil {
		return err
	}
	err = moduleGen.CreateRepositoryFile()
	if err != nil {
		return err
	}

	err = moduleGen.CreateServiceDir()
	if err != nil {
		return err
	}
	err = moduleGen.CreateServiceFile()
	if err != nil {
		return err
	}

	moduleGen.CreateBaseHandlerDir()
	err = moduleGen.CreateHandlerBaseModuleFile()
	if err != nil {
		return err
	}

	moduleGen.AppendModuleToWire()

	fs.GoFormat(s.Info.ProjectPath, s.Info.GoModName)

	logger.Infoln(fmt.Sprintf("Your Module '%s' already created under '%s' App", s.Name, s.Info.GoModName))

	return nil
}
