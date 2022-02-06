package misc

import (
	"github.com/nurcahyaari/kite/library/impl"
	"github.com/nurcahyaari/kite/templates"
	"github.com/nurcahyaari/kite/utils/fs"
)

type ConfigureEnvOption struct {
	impl.GeneratorOptions
}

func NewConfigureEnv(opt ConfigureEnvOption) impl.AppGenerator {
	return ConfigureEnvOption{
		GeneratorOptions: opt.GeneratorOptions,
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
