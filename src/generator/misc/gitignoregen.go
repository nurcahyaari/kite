package misc

import (
	"github.com/nurcahyaari/kite/src/templates/misctemplate"
	"github.com/nurcahyaari/kite/src/utils/fs"
)

type GitIgnoreGen interface {
	CreateGitignoreFiles() error
}

type GitIgnoreGenImpl struct {
	ProjectPath string
}

func NewGitignore(projectPath string) GitIgnoreGen {
	return &GitIgnoreGenImpl{
		ProjectPath: projectPath,
	}
}

func (s GitIgnoreGenImpl) CreateGitignoreFiles() error {
	templateNew := misctemplate.NewGitignoreTemplate()
	gitignoreTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	return fs.CreateFileIfNotExist(s.ProjectPath, ".gitignore", gitignoreTemplate)
}