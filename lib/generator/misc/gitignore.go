package misc

import (
	"github.com/nurcahyaari/kite/lib/impl"
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type GitignoreOptions struct {
	impl.GeneratorOptions
}

func NewGitignore(options GitignoreOptions) impl.AppGenerator {
	return options
}

func (o GitignoreOptions) Run() error {
	o.createGitignoreFiles()
	return nil
}

func (o GitignoreOptions) createGitignoreFiles() error {
	tmpl := templates.NewTemplate(templates.Template{
		Template: templates.GitignoreTemplate,
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	err = fs.CreateFileIfNotExist(o.ProjectPath, ".gitignore", templateString)
	if err != nil {
		return err
	}
	return nil
}
