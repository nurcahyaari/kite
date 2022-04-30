package miscgen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates/misctemplate"
)

type MakefileGen interface {
	CreateMakefilefile(dto MiscDto) error
}

type MakefileGenImpl struct {
	fs database.FileSystem
}

func NewMakefileGen(
	fs database.FileSystem,
) *MakefileGenImpl {
	return &MakefileGenImpl{
		fs: fs,
	}
}

func (s MakefileGenImpl) CreateMakefilefile(dto MiscDto) error {
	templateNew := misctemplate.NewMakefileTemplate()
	makefileTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	return s.fs.CreateFileIfNotExists(dto.ProjectPath, "makefile", makefileTemplate)
}
