package generator

import (
	"errors"
	"fmt"

	appmodule "github.com/nurcahyaari/kite/lib/generator/appmodule"
	"github.com/nurcahyaari/kite/lib/generator/infrastructure"
	"github.com/nurcahyaari/kite/lib/generator/misc"
	"github.com/nurcahyaari/kite/lib/impl"

	"github.com/nurcahyaari/kite/utils/fs"
	"github.com/nurcahyaari/kite/utils/logger"
	"github.com/nurcahyaari/kite/utils/pkg"
)

type NewAppOption struct {
	impl.GeneratorOptions
}

func NewApp(opt NewAppOption) impl.AppGenerator {
	return NewAppOption{
		GeneratorOptions: opt.GeneratorOptions,
	}
}

func (o NewAppOption) Run() error {
	logger.Infoln(fmt.Sprintf("Creating %s", o.AppName))
	err := o.createAppDir()
	if err != nil {
		return errors.New(fmt.Sprintf("%s already created", o.AppName))
	}

	// list package
	appPkg := pkg.AppPackages{
		Packages: []string{
			"github.com/rs/zerolog",
			"github.com/spf13/viper",
			"github.com/go-sql-driver/mysql",
			"github.com/jmoiron/sqlx",
			"github.com/go-chi/chi/v5",
			"github.com/golang-jwt/jwt",
			"github.com/google/wire",
			"github.com/google/wire/cmd/wire",
			"github.com/swaggo/swag/cmd/swag",
		},
	}

	// create .env
	envOption := misc.ConfigureEnvOption{
		GeneratorOptions: o.GeneratorOptions,
	}
	env := misc.NewConfigureEnv(envOption)
	env.Run()

	// create .gitignore
	gitignoreOption := misc.GitignoreOptions{
		GeneratorOptions: o.GeneratorOptions,
	}
	gitignore := misc.NewGitignore(gitignoreOption)
	gitignore.Run()

	// create go mod
	err = fs.GoModInit(o.ProjectPath, o.GoModName)
	if err != nil {
		o.Rollback()
		return errors.New("There was a go mod file in this Folder")
	}
	logger.Infoln("Success to run go mod init on your project")

	appOption := ApplicationOption{
		GeneratorOptions: o.GeneratorOptions,
	}
	app := NewApplicationGenerator(appOption)
	err = app.Run()
	if err != nil {
		o.Rollback()
		return err
	}

	// create config folder
	configOption := misc.ConfigOption{
		GeneratorOptions: o.GeneratorOptions,
	}
	config := misc.NewConfig(configOption)
	err = config.Run()
	if err != nil {
		o.Rollback()
		return err
	}

	// create infrastructure folder
	infrastructureOption := infrastructure.InfrastuctureOption{
		GeneratorOptions: o.GeneratorOptions,
	}

	infrastructure := infrastructure.NewInfrastructure(infrastructureOption)
	err = infrastructure.Run()
	if err != nil {
		o.Rollback()
		return err
	}

	// create internal
	internalOption := InternalOption{
		GeneratorOptions: o.GeneratorOptions,
	}
	internal := NewInternal(internalOption)
	err = internal.Run()
	if err != nil {
		o.Rollback()
		return err
	}

	// create src
	moduleOption := appmodule.ModulesOption{
		GeneratorOptions: o.GeneratorOptions,
	}

	module := appmodule.NewModules(moduleOption)
	err = module.Run()
	if err != nil {
		o.Rollback()
		return err
	}

	err = appPkg.InstallPackage()
	if err != nil {
		return errors.New("error when installing package")
	}

	// add wire
	wireOption := misc.WireOptions{
		GeneratorOptions: o.GeneratorOptions,
	}
	wire := misc.NewWire(wireOption)
	err = wire.Run()
	if err != nil {
		return errors.New("error when creating wire")
	}

	fs.GoGenerateRun(o.ProjectPath)

	fs.GoFormat(o.ProjectPath, o.GoModName)

	logger.Infoln(fmt.Sprintf("Your App '%s' already created", o.GoModName))

	return nil
}

func (o NewAppOption) Rollback() {
	fs.DeleteFolder(o.ProjectPath)
}

func (o NewAppOption) createAppDir() error {
	err := fs.CreateFolderIsNotExist(o.ProjectPath)
	return err
}
