package infrastructure

import (
	"github.com/nurcahyaari/kite/utils/fs"
)

type CachedOption struct {
	Options
	InfrastructurePath string
	DirName            string
	DirPath            string
	DBType             string
}

func NewCached(options CachedOption) AppGenerator {
	options.DirName = "cached"
	options.DirPath = fs.ConcatDirPath(options.InfrastructurePath, options.DirName)

	return options
}

func (o CachedOption) Run() error {

	// o.createCacheDir()

	return nil
}

// func (o CachedOption) createCacheDir() error {
// 	logger.Info("Create infrastructure/cached directory... ")
// 	err := fs.CreateFolderIsNotExist(o.DirPath)
// 	if err != nil {
// 		return err
// 	}
// 	logger.InfoSuccessln("success")
// 	return nil
// }
