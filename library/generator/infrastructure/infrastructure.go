package infrastructure

import (
	"github.com/nurcahyaari/kite/utils/fs"
	"github.com/nurcahyaari/kite/utils/logger"
)

type InfrastuctureOption struct {
	Options
	DirName            string
	DirPath            string
	InfrastructureName string
	InfrastructurePath string
}

func NewInfrastructure(opt InfrastuctureOption) AppGenerator {
	dirName := "infrastructure"
	dirPath := fs.ConcatDirPath(opt.ProjectPath, dirName)
	InfrastructurePath := fs.ConcatDirPath(dirPath, opt.InfrastructureName)

	return InfrastuctureOption{
		Options:            opt.Options,
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
		Options:            o.Options,
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
		Options:            o.Options,
		InfrastructurePath: o.DirPath,
	}
	db := NewCached(cachedOption)
	db.Run()
}

func (o InfrastuctureOption) createQueue() {
	queueOption := QueueOption{
		Options:            o.Options,
		InfrastructurePath: o.DirPath,
	}
	db := NewQueue(queueOption)
	db.Run()
}
