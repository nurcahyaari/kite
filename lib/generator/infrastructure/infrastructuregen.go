package infrastructure

import (
	"github.com/nurcahyaari/kite/utils/fs"
	"github.com/nurcahyaari/kite/utils/logger"
)

type InfrastructureGen interface {
	CacheGen
	DatabaseGen
	CreateInfrastructureDir() error
	GenerateInfrastructure() error
}

type InfrastructureGenImpl struct {
	GomodName          string
	InfrastructurePath string
	*CacheGenImpl
	*DatabaseGenImpl
}

func NewInfrastructureGen(
	projectPath string,
	gomodName string,
) InfrastructureGen {
	infrastructurePath := fs.ConcatDirPath(projectPath, "infrastructure")
	appName := fs.GetAppNameBasedOnGoMod(gomodName)

	return &InfrastructureGenImpl{
		InfrastructurePath: infrastructurePath,
		CacheGenImpl:       NewCacheGen(),
		DatabaseGenImpl: NewDatabaseGen(
			DatabaseGenImpl{
				AppName:            appName,
				GomodName:          gomodName,
				InfrastructurePath: infrastructurePath,
			},
		),
	}
}

func (s *InfrastructureGenImpl) CreateInfrastructureDir() error {
	logger.Info("Creating infrastructure directory... ")
	err := fs.CreateFolderIsNotExist(s.InfrastructurePath)
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")

	return nil
}

func (s *InfrastructureGenImpl) GenerateInfrastructure() error {
	return s.CreateMysqlConnection()
}
