package infrastructure

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/lib/ast"
	"github.com/nurcahyaari/kite/templates/infrastructuretemplate/databasetemplate"
	"github.com/nurcahyaari/kite/utils/fs"
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
	CreateMysqlConnection() error
}

type DatabaseGenImpl struct {
	AppName            string
	InfrastructurePath string
	GomodName          string
	DatabaseType       DbType
}

func NewDatabaseGen(dbGenImpl DatabaseGenImpl) *DatabaseGenImpl {
	return &dbGenImpl
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
		s.InfrastructurePath,
		fmt.Sprintf("%s.go", s.DatabaseType.ToString()),
		databaseCode,
	)
}
