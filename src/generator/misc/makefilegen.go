package misc

import (
	"github.com/nurcahyaari/kite/src/templates/misctemplate"
	"github.com/nurcahyaari/kite/src/utils/fs"
)

type MakefileGen interface {
	CreateMakefilefile() error
}

type MakefileGenImpl struct {
	ProjectPath string
}

func NewMakefile(projectPath string) MakefileGen {
	return &MakefileGenImpl{
		ProjectPath: projectPath,
	}
}

func (s MakefileGenImpl) CreateMakefilefile() error {
	templateNew := misctemplate.NewMakefileTemplate()
	makefileTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	return fs.CreateFileIfNotExist(s.ProjectPath, "makefile", makefileTemplate)
}
