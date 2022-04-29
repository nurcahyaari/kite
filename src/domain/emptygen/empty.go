package emptygen

import (
	"fmt"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates"
	"github.com/nurcahyaari/kite/internal/utils"
)

type EmptyGen interface {
	CreateEmptyGolangFile(dto EmptyDto) error
}

type EmptyGenImpl struct {
	fs database.FileSystem
}

func NewEmptyGen(fs database.FileSystem) *EmptyGenImpl {
	return &EmptyGenImpl{
		fs: fs,
	}
}

func (s EmptyGenImpl) CreateEmptyGolangFile(dto EmptyDto) error {
	template := templates.NewTemplateNewImpl(dto.PackageName, "")
	templateCode, err := template.Render("", nil)
	if err != nil {
		return err
	}

	if !s.fs.IsFileExists(utils.ConcatDirPath(dto.Path, dto.FileName)) {
		s.fs.CreateFileIfNotExists(dto.Path, fmt.Sprintf("%s.go", dto.FileName), templateCode)
	}

	return nil
}
