package dbgen

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/logger"
	"github.com/nurcahyaari/kite/internal/templates/infrastructuretemplate/databasetemplate"
	"github.com/nurcahyaari/kite/internal/utils/ast"
)

type DbType int

const (
	DbMysql DbType = iota
)

const (
	MysqlCode string = "mysql"
)

func (s DbType) ToDatabaseTemplateType() databasetemplate.DatabaseType {
	var dbType databasetemplate.DatabaseType
	switch s {
	case DbMysql:
		dbType = databasetemplate.DatabaseMysql
	}
	return dbType
}

func (s DbType) ToString() string {
	var dbType string
	switch s {
	case DbMysql:
		dbType = MysqlCode
	}
	return dbType
}

type DatabaseGen interface {
	CreateDatabaseDir(option DBOption) error
	CreateMysqlConnection(option DBOption) error
}

type DatabaseGenImpl struct {
	fs database.FileSystem
}

func NewDatabaseGen(
	fs database.FileSystem,
) *DatabaseGenImpl {
	// dbGenImpl.DatabasePath = utils.ConcatDirPath(dbGenImpl.InfrastructurePath, "database")
	// dbGenImpl.fs = database.NewFileSystem(dbGenImpl.DatabasePath)
	return &DatabaseGenImpl{
		fs: fs,
	}
}

func (s *DatabaseGenImpl) CreateDatabaseDir(option DBOption) error {
	logger.Info("Creating infrastructure/database directory... ")
	err := s.fs.CreateFolderIfNotExists(option.DatabasePath)
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")

	return nil
}

func (s *DatabaseGenImpl) CreateMysqlConnection(option DBOption) error {
	templateNew := databasetemplate.NewDatabaseTemplate(databasetemplate.DatabaseTemplateData{
		DatabaseType: option.DatabaseType.ToDatabaseTemplateType(),
	})
	databaseTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	databaseAbstractCode := ast.NewAbstractCode(databaseTemplate, parser.ParseComments)
	databaseAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"fmt\"",
	})
	databaseAbstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s/config\"", option.GomodName),
	})
	databaseAbstractCode.AddImport(ast.ImportSpec{
		Name: "_",
		Path: "\"github.com/go-sql-driver/mysql\"",
	})
	databaseAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/jmoiron/sqlx\"",
	})
	databaseAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/rs/zerolog/log\"",
	})

	err = databaseAbstractCode.RebuildCode()
	if err != nil {
		return err
	}
	// get the manipulate code
	databaseCode := databaseAbstractCode.GetCode()

	return s.fs.CreateFileIfNotExists(
		option.DatabasePath,
		fmt.Sprintf("%s.go", option.DatabaseType.ToString()),
		databaseCode,
	)
}
