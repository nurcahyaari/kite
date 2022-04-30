package databasetype

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates/infrastructuretemplate/databasetemplate"
	"github.com/nurcahyaari/kite/internal/utils/ast"
)

type MysqlGen interface {
	CreateMysqlConnection(dto DatabaseTypeDto) error
}

type MysqlGenImpl struct {
	fs database.FileSystem
}

func NewMysqlGen(fs database.FileSystem) *MysqlGenImpl {
	return &MysqlGenImpl{
		fs: fs,
	}
}

func (s *MysqlGenImpl) CreateMysqlConnection(dto DatabaseTypeDto) error {
	templateNew := databasetemplate.NewMysqlTemplate()
	databaseTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	databaseAbstractCode := ast.NewAbstractCode(databaseTemplate, parser.ParseComments)
	databaseAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"fmt\"",
	})
	databaseAbstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s/config\"", dto.GomodName),
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
		dto.Path,
		fmt.Sprintf("%s.go", "mysql"),
		databaseCode,
	)
}
