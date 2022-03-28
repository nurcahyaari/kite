package logger

import (
	"github.com/nurcahyaari/kite/templates"
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

	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "logger",
		Template:    templates.Loggertemplate,
		Import: []templates.ImportedPackage{
			{
				FilePath: "os",
			},
			{
				FilePath: "time",
			},
			{
				FilePath: "github.com/rs/zerolog",
			},
			{
				FilePath: "github.com/rs/zerolog/log",
			},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	err = fs.CreateFileIfNotExist(s.LoggerDir, "log.go", templateString)
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")
	return nil
}
