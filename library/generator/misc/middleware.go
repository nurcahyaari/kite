package misc

import (
	"fmt"

	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type MiddlewareOption struct {
	Options
	InternalPath   string
	DirName        string
	DirPath        string
	MiddlewareName string
}

func NewMiddleware(options MiddlewareOption) AppGenerator {
	options.DirName = "middlewares"
	options.DirPath = fs.ConcatDirPath(options.InternalPath, options.DirName)

	return options
}

func (o MiddlewareOption) Run() error {
	o.createMiddlewaresDir()

	if o.MiddlewareName != "" {
		o.createFile()
	}

	return nil
}

func (o MiddlewareOption) createMiddlewaresDir() error {
	err := fs.CreateFolderIsNotExist(o.DirPath)
	if err != nil {
		return err
	}

	return nil
}

func (o MiddlewareOption) createFile() error {
	middlewareName := fmt.Sprintf("%s.go", o.MiddlewareName)

	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "middleware",
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	fs.CreateFileIfNotExist(o.DirPath, middlewareName, templateString)
	return nil
}
