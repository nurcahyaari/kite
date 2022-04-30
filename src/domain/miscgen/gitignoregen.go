package miscgen

import (
	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates/misctemplate"
)

type GitIgnoreGen interface {
	CreateGitignoreFile(dto MiscDto) error
}

type GitIgnoreGenImpl struct {
	ProjectPath string
	fs          database.FileSystem
}

func NewGitignore(fs database.FileSystem) *GitIgnoreGenImpl {
	return &GitIgnoreGenImpl{
		fs: fs,
	}
}

func (s GitIgnoreGenImpl) CreateGitignoreFile(dto MiscDto) error {
	templateNew := misctemplate.NewGitignoreTemplate()
	gitignoreTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	return s.fs.CreateFileIfNotExists(dto.ProjectPath, ".gitignore", gitignoreTemplate)
}
