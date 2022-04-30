package dbgen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/src/domain/dbgen/databasetype"
)

type DatabaseGen interface {
	CreateDatabaseDir(dto DatabaseDto) error
	CreateDatabaseConnection(dto DatabaseDto) error
}

type DatabaseGenImpl struct {
	fs       database.FileSystem
	mysqlGen databasetype.MysqlGen
}

func NewDatabaseGen(
	fs database.FileSystem,
	mysqlGen databasetype.MysqlGen,
) *DatabaseGenImpl {
	return &DatabaseGenImpl{
		fs:       fs,
		mysqlGen: mysqlGen,
	}
}

func (s DatabaseGenImpl) CreateDatabaseDir(dto DatabaseDto) error {
	err := s.fs.CreateFolderIfNotExists(dto.DatabasePath)
	if err != nil {
		return err
	}

	return nil
}

func (s DatabaseGenImpl) CreateDatabaseConnection(dto DatabaseDto) error {
	var err error
	databaseTypeDto := databasetype.DatabaseTypeDto{
		Path:        dto.DatabasePath,
		GomodName:   dto.GomodName,
		ProjectPath: dto.ProjectPath,
	}
	switch dto.DatabaseType {
	case DbMysql:
		err = s.mysqlGen.CreateMysqlConnection(databaseTypeDto)
	}
	return err
}
