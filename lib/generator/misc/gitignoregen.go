package misc

import (
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
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
	tmpl := templates.NewTemplate(templates.Template{
		Template: templates.GitignoreTemplate,
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	return fs.CreateFileIfNotExist(s.ProjectPath, ".gitignore", templateString)
}
