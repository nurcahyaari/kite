package logger

import (
	"go/parser"

	"github.com/nurcahyaari/kite/lib/ast"
	"github.com/nurcahyaari/kite/templates/internaltemplate/loggertemplate"
	"github.com/nurcahyaari/kite/utils/fs"
	"github.com/nurcahyaari/kite/utils/logger"
)

type LoggerGen interface {
	CreateLoggerDir() error
	CreateDefaultLoggerFile() error
}

type LoggerGenImpl struct {
	LoggerDir string
}

func NewLoggerGen(internalPath string) *LoggerGenImpl {
	loggerDir := fs.ConcatDirPath(internalPath, "logger")
	return &LoggerGenImpl{
		LoggerDir: loggerDir,
	}
}

func (s *LoggerGenImpl) CreateLoggerDir() error {
	logger.Info("Creating internal/logger directory... ")
	err := fs.CreateFolderIsNotExist(s.LoggerDir)
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")
	return nil
}

func (s *LoggerGenImpl) CreateDefaultLoggerFile() error {
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

	err = fs.CreateFileIfNotExist(s.LoggerDir, "log.go", loggerCode)
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")
	return nil
}
