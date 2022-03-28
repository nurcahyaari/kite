package misc

import (
	"fmt"

	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type WireGen interface {
	CreateWireFiles() error
}

type WireGenImpl struct {
	GomodName   string
	ProjectPath string
}

func NewWire(
	projectPath string,
	gomodName string,
) WireGen {
	return &WireGenImpl{
		ProjectPath: projectPath,
		GomodName:   gomodName,
	}
}

func (s WireGenImpl) CreateWireFiles() error {
	tmpl := templates.NewTemplate(templates.Template{
		PackageName: "main",
		Template:    templates.WireTemplate,
		Header:      "//+build wireinject",
		Import: []templates.ImportedPackage{
			{
				FilePath: "github.com/google/wire",
			},
			{
				Alias:    "httprouter",
				FilePath: fmt.Sprintf("%s/internal/protocol/http/router", s.GomodName),
			},
			{
				Alias:    "httphandler",
				FilePath: fmt.Sprintf("%s/src/handlers/http", s.GomodName),
			},
			{
				FilePath: fmt.Sprintf("%s/internal/protocol/http", s.GomodName),
			},
			{
				FilePath: fmt.Sprintf("%s/infrastructure", s.GomodName),
			},
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	return fs.CreateFileIfNotExist(s.ProjectPath, "wire.go", templateString)
}
