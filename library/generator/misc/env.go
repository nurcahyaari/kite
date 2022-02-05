package misc

import (
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type ConfigureEnvOption struct {
	Options
}

func NewConfigureEnv(opt ConfigureEnvOption) AppGenerator {
	return ConfigureEnvOption{
		Options: opt.Options,
	}
}

func (o ConfigureEnvOption) Run() error {
	tmpl := templates.NewTemplate(templates.Template{
		Template: templates.EnvTemplate,
		Data: map[string]interface{}{
			"DBDialeg": o.DefaultDBDialeg,
		},
	})

	templateString, err := tmpl.Render()
	if err != nil {
		return err
	}

	fs.CreateFileIfNotExist(o.ProjectPath, ".env", templateString)
	fs.CreateFileIfNotExist(o.ProjectPath, ".env.example", templateString)

	return nil
}
