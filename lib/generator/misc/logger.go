package misc

import (
	"github.com/nurcahyaari/kite/lib/impl"
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
	"github.com/nurcahyaari/kite/utils/logger"
)

type LoggerOption struct {
	impl.GeneratorOptions
	InfrastructurePath string
	DirName            string
	DirPath            string
	QueueType          string
}

func NewLogger(options LoggerOption) impl.AppGenerator {
	options.DirName = "logger"
	options.DirPath = fs.ConcatDirPath(options.InfrastructurePath, options.DirName)

	return options
}

func (o LoggerOption) Run() error {

	o.createLoggerDir()
	o.createDefaultLoggerFile()

	return nil
}

func (o LoggerOption) createLoggerDir() error {
	logger.Info("Creating internal/logger directory... ")
	err := fs.CreateFolderIsNotExist(o.DirPath)
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")
	return nil
}

func (o LoggerOption) createDefaultLoggerFile() error {
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

	err = fs.CreateFileIfNotExist(o.DirPath, "log.go", templateString)
	if err != nil {
		return err
	}
	logger.InfoSuccessln("success")
	return nil
}
