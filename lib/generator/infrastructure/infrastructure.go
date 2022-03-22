package infrastructure

import (
	"github.com/nurcahyaari/kite/lib/impl"
	"github.com/nurcahyaari/kite/utils/fs"
	"github.com/nurcahyaari/kite/utils/logger"
)

type InfrastuctureOption struct {
	impl.GeneratorOptions
	DirName            string
	DirPath            string
	InfrastructureName string
	InfrastructurePath string
}

func NewInfrastructure(opt InfrastuctureOption) impl.AppGenerator {
	dirName := "infrastructure"
	dirPath := fs.ConcatDirPath(opt.ProjectPath, dirName)
	InfrastructurePath := fs.ConcatDirPath(dirPath, opt.InfrastructureName)

	return InfrastuctureOption{
		GeneratorOptions:   opt.GeneratorOptions,
		DirName:            dirName,
		DirPath:            dirPath,
		InfrastructureName: opt.InfrastructureName,
		InfrastructurePath: InfrastructurePath,
	}
}

func (o InfrastuctureOption) Run() error {

	err := o.createInfrastructureDir()
	if err != nil {
		return err
	}
	err = o.createDB()
	if err != nil {
		return err
	}
	o.createCached()
	o.createQueue()

	return nil
}

func (o InfrastuctureOption) createInfrastructureDir() error {
	logger.Info("Creating infrastructure directory... ")
	err := fs.CreateFolderIsNotExist(o.DirPath)
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")

	return nil
}

func (o InfrastuctureOption) createDB() error {
	logger.Info("Creating database connection... ")
	dbOption := DBOption{
		GeneratorOptions:   o.GeneratorOptions,
		InfrastructurePath: o.DirPath,
	}
	db := NewDB(dbOption)
	err := db.Run()
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")
	return nil
}

func (o InfrastuctureOption) createCached() {
	cachedOption := CachedOption{
		GeneratorOptions:   o.GeneratorOptions,
		InfrastructurePath: o.DirPath,
	}
	db := NewCached(cachedOption)
	db.Run()
}

func (o InfrastuctureOption) createQueue() {
	queueOption := QueueOption{
		GeneratorOptions:   o.GeneratorOptions,
		InfrastructurePath: o.DirPath,
	}
	db := NewQueue(queueOption)
	db.Run()
}