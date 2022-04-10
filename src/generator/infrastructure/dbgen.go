package infrastructure

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/src/ast"
	"github.com/nurcahyaari/kite/src/templates/infrastructuretemplate/databasetemplate"
	"github.com/nurcahyaari/kite/src/utils/fs"
	"github.com/nurcahyaari/kite/src/utils/logger"
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
	CreateDatabaseDir() error
	CreateMysqlConnection() error
}

type DatabaseGenImpl struct {
	AppName            string
	InfrastructurePath string
	DatabasePath       string
	GomodName          string
	DatabaseType       DbType
}

func NewDatabaseGen(dbGenImpl DatabaseGenImpl) *DatabaseGenImpl {
	dbGenImpl.DatabasePath = fs.ConcatDirPath(dbGenImpl.InfrastructurePath, "database")
	return &dbGenImpl
}

func (s *DatabaseGenImpl) CreateDatabaseDir() error {
	logger.Info("Creating infrastructure/database directory... ")
	err := fs.CreateFolderIsNotExist(s.DatabasePath)
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")

	return nil
}

func (s *DatabaseGenImpl) CreateMysqlConnection() error {
	templateNew := databasetemplate.NewDatabaseTemplate(databasetemplate.DatabaseTemplateData{
		DatabaseType: s.DatabaseType.ToDatabaseTemplateType(),
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
		Path: fmt.Sprintf("\"%s/config\"", s.GomodName),
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

	return fs.CreateFileIfNotExist(
		s.DatabasePath,
		fmt.Sprintf("%s.go", s.DatabaseType.ToString()),
		databaseCode,
	)
}
