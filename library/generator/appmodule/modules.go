package appmodule

import (
	"github.com/nurcahyaari/kite/library/generator/model"

	"github.com/nurcahyaari/kite/utils/fs"
)

type ModulesOption struct {
	Options
	IsModule   bool
	DirName    string
	DirPath    string
	ModuleName string
	ModulePath string
}

func NewModules(opt ModulesOption) AppGenerator {
	dirName := "src"
	dirPath := fs.ConcatDirPath(opt.ProjectPath, dirName)
	modulesPath := fs.ConcatDirPath(fs.ConcatDirPath(dirPath, "modules"), opt.ModuleName)

	if opt.GoModName == "" {
		opt.GoModName = fs.GetGoModName(opt.ProjectPath)
	}

	return &ModulesOption{
		Options:    opt.Options,
		IsModule:   opt.IsModule,
		DirName:    dirName,
		DirPath:    dirPath,
		ModuleName: opt.ModuleName,
		ModulePath: modulesPath,
	}
}

func (o ModulesOption) Run() error {
	var err error

	if o.IsModule {
		err = o.createModulesDir()
		if err != nil {
			return err
		}
	} else {
		err = o.createSrcDir()
		if err != nil {
			return err
		}
	}

	fs.GoFormat(o.ProjectPath, o.GoModName)

	return nil
}

func (o ModulesOption) createSrcDir() error {
	// create src
	err := fs.CreateFolderIsNotExist(o.DirPath)
	if err != nil {
		return err
	}

	// create handler folder
	err = o.createHandlerDir()
	if err != nil {
		return err
	}

	// create modules folder
	err = fs.CreateFolderIsNotExist(fs.ConcatDirPath(o.DirPath, "modules"))
	if err != nil {
		return err
	}

	return nil
}

func (o ModulesOption) createModulesDir() error {

	o.createHandlerDir()
	o.createRepositoryDir()
	o.createServiceDir()
	o.createEntitiesDir()
	o.createDTODir()

	return nil
}

func (o ModulesOption) createHandlerDir() error {
	handlerOption := HandlerOption{
		Options:    o.Options,
		DirPath:    o.DirPath,
		ModuleName: o.ModuleName,
	}
	handler, _ := NewHandler(handlerOption)
	err := handler.Run()
	if err != nil {
		return err
	}

	return nil
}

func (o ModulesOption) createRepositoryDir() error {

	repositoryOption := RepositoryOption{
		Options:    o.Options,
		ModuleName: o.ModuleName,
		ModulePath: o.ModulePath,
	}
	repository, _ := NewRepository(repositoryOption)
	repository.Run()

	return nil
}

func (o ModulesOption) createServiceDir() error {

	serviceOption := ServiceOption{
		Options:    o.Options,
		ModuleName: o.ModuleName,
		ModulePath: o.ModulePath,
	}
	service, _ := NewService(serviceOption)
	service.Run()

	return nil
}

func (o ModulesOption) createEntitiesDir() error {
	entitiesOption := model.EntityOption{
		Options:     o.Options,
		ModulesPath: o.ModulePath,
	}
	entity, _ := model.NewEntity(entitiesOption)
	entity.Run()

	return nil
}

func (o ModulesOption) createDTODir() error {
	dtoOption := model.DTOOption{
		Options:     o.Options,
		ModulesPath: o.ModulePath,
	}
	dto, _ := model.NewDTO(dtoOption)
	dto.Run()

	return nil
}
