package infrastructuregen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/logger"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/src/domain/cachegen"
	"github.com/nurcahyaari/kite/src/domain/dbgen"
)

type InfrastructureGen interface {
	CreateInfrastructureDir(dto InfrastructureDto) error
	GenerateInfrastructure(dto InfrastructureDto) error
}

type InfrastructureGenImpl struct {
	cacheGen cachegen.CacheGen
	dbGen    dbgen.DatabaseGen
	fs       database.FileSystem
}

func NewInfrastructureGen(
	cacheGen cachegen.CacheGen,
	dbGen dbgen.DatabaseGen,
	fs database.FileSystem,
) *InfrastructureGenImpl {
	return &InfrastructureGenImpl{
		cacheGen: cacheGen,
		dbGen:    dbGen,
		fs:       fs,
	}
}

func (s *InfrastructureGenImpl) CreateInfrastructureDir(dto InfrastructureDto) error {
	logger.Info("Creating infrastructure directory... ")
	err := s.fs.CreateFolderIfNotExists(dto.InfrastructurePath)
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")

	return nil
}

func (s *InfrastructureGenImpl) GenerateInfrastructure(dto InfrastructureDto) error {
	databasePath := utils.ConcatDirPath(dto.InfrastructurePath, "database")

	dbDto := dbgen.DBOption{
		GomodName:    dto.GomodName,
		DatabasePath: databasePath,
		DatabaseType: dto.DatabaseType,
	}
	err := s.dbGen.CreateDatabaseDir(dbDto)
	if err != nil {
		return err
	}

	err = s.dbGen.CreateMysqlConnection(dbDto)
	if err != nil {
		return err
	}

	return nil
}
