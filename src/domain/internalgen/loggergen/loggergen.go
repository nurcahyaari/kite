package loggergen

import (
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/logger"
	"github.com/nurcahyaari/kite/internal/templates/internaltemplate/loggertemplate"
	"github.com/nurcahyaari/kite/internal/utils/ast"
)

type LoggerGen interface {
	CreateLoggerDir(dto LoggerDto) error
	CreateDefaultLoggerFile(dto LoggerDto) error
}

type LoggerGenImpl struct {
	fs database.FileSystem
}

func NewLoggerGen(
	fs database.FileSystem,
) *LoggerGenImpl {
	return &LoggerGenImpl{
		fs: fs,
	}
}

func (s LoggerGenImpl) CreateLoggerDir(dto LoggerDto) error {
	logger.Info("Creating internal/logger directory... ")
	err := s.fs.CreateFolderIfNotExists(dto.Path)
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")
	return nil
}

func (s LoggerGenImpl) CreateDefaultLoggerFile(dto LoggerDto) error {
	logger.Info("Creating internal/logger/log.go file... ")

	templateNew := loggertemplate.NewLoggerTemplate()
	loggerTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	loggerAbstractCode := ast.NewAbstractCode(loggerTemplate, parser.ParseComments)
	loggerAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"os\"",
	})
	loggerAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"time\"",
	})
	loggerAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/rs/zerolog\"",
	})
	loggerAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"github.com/rs/zerolog/log\"",
	})
	// after manipulate the code, rebuild the code
	err = loggerAbstractCode.RebuildCode()
	if err != nil {
		return err
	}
	// get the manipulate code
	loggerCode := loggerAbstractCode.GetCode()

	err = s.fs.CreateFileIfNotExists(dto.Path, "log.go", loggerCode)
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")
	return nil
}
