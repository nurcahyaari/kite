package generator

import (
	"errors"
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/src/ast"
	"github.com/nurcahyaari/kite/src/generator/config"
	"github.com/nurcahyaari/kite/src/generator/infrastructure"
	"github.com/nurcahyaari/kite/src/generator/internal"
	"github.com/nurcahyaari/kite/src/generator/misc"
	"github.com/nurcahyaari/kite/src/generator/module"
	"github.com/nurcahyaari/kite/src/templates/misctemplate"
	"github.com/nurcahyaari/kite/src/utils/fs"
	"github.com/nurcahyaari/kite/src/utils/logger"
	"github.com/nurcahyaari/kite/src/utils/pkg"
)

type App interface {
	CreateNewApp() error
	// Private
	createMainApp() error
	rollback()
}

// App.go is useful to creating a new app
type AppImpl struct {
	Info ProjectInfo
	// The name of app
	Name string
}

func NewAppNew(info ProjectInfo) App {
	return &AppImpl{
		Name: fs.GetAppNameBasedOnGoMod(info.GoModName),
		Info: info,
	}
}

func (s AppImpl) CreateNewApp() error {
	logger.Infoln(fmt.Sprintf("Creating %s", s.Name))

	err := fs.CreateFolderIsNotExist(s.Info.ProjectPath)
	if err != nil {
		return errors.New(fmt.Sprintf("%s already created", s.Name))
	}

	// list package
	appPkg := pkg.AppPackages{
		Packages: []string{
			"github.com/rs/zerolog",
			"github.com/spf13/viper",
			"github.com/go-sql-driver/mysql",
			"github.com/jmoiron/sqlx",
			"github.com/go-chi/chi/v5",
			"github.com/go-chi/cors",
			"github.com/golang-jwt/jwt",
			"github.com/google/wire",
			"github.com/google/wire/cmd/wire",
			"github.com/swaggo/swag/cmd/swag",
			"github.com/swaggo/http-swagger",
		},
	}

	err = fs.Gitinit(s.Info.ProjectPath)
	if err != nil {
		logger.Warnln(err.Error())
	}

	// create gitignore
	gitignoreGen := misc.NewGitignore(s.Info.ProjectPath)
	gitignoreGen.CreateGitignoreFile()

	// create makefile
	makefileGen := misc.NewMakefile(s.Info.ProjectPath)
	makefileGen.CreateMakefilefile()

	// create go.mod
	err = fs.GoModInit(s.Info.ProjectPath, s.Info.GoModName)
	if err != nil {
		s.rollback()
		return err
	}

	// setup env
	envGen := config.NewEnvGen(s.Info.ProjectPath)
	// create .env.example
	err = envGen.CreateEnvExampleFile()
	if err != nil {
		s.rollback()
		return err
	}
	// create .env
	err = envGen.CreateEnvFile()
	if err != nil {
		s.rollback()
		return err
	}

	// create config
	configGen := config.NewConfig(s.Info.ProjectPath, s.Name)
	err = configGen.CreateConfigDir()
	if err != nil {
		s.rollback()
		return err
	}

	err = configGen.CreateConfigFile()
	if err != nil {
		s.rollback()
		return err
	}

	// create infra file
	infraGen := infrastructure.NewInfrastructureGen(
		s.Info.ProjectPath,
		s.Info.GoModName,
	)
	err = infraGen.CreateInfrastructureDir()
	if err != nil {
		s.rollback()
		return err
	}

	err = infraGen.CreateDatabaseDir()
	if err != nil {
		s.rollback()
		return err
	}

	err = infraGen.GenerateInfrastructure()
	if err != nil {
		s.rollback()
		return err
	}

	// create internal
	internalGen := internal.NewInternal(s.Info.ProjectPath, s.Info.GoModName)
	err = internalGen.CreateInternalDir()
	if err != nil {
		s.rollback()
		return err
	}

	err = internalGen.CreateUtilDir()
	if err != nil {
		s.rollback()
		return err
	}

	err = internalGen.CreateRsaReader()
	if err != nil {
		s.rollback()
		return err
	}

	err = internalGen.CreateLoggerDir()
	if err != nil {
		s.rollback()
		return err
	}

	err = internalGen.CreateDefaultLoggerFile()
	if err != nil {
		s.rollback()
		return err
	}

	err = internalGen.CreateInternalProtocolDir()
	if err != nil {
		s.rollback()
		return err
	}

	err = internalGen.CreateProtocolTypeDir()
	if err != nil {
		s.rollback()
		return err
	}

	// create module parent path
	moduleGen := module.NewModuleGen(s.Info.ProjectPath, "", s.Info.GoModName)
	err = moduleGen.CreateSrcDir()
	if err != nil {
		s.rollback()
		return err
	}

	err = moduleGen.CreateBaseHandlerDir()
	if err != nil {
		s.rollback()
		return err
	}

	err = moduleGen.CreateHandlerBaseFile()
	if err != nil {
		return err
	}

	moduleGen.CreateBaseModuleDir()

	// create main file
	err = s.createMainApp()
	if err != nil {
		s.rollback()
		return err
	}

	// wire
	wireGen := misc.NewWire(s.Info.ProjectPath, s.Info.GoModName)
	err = wireGen.CreateWireFiles()
	if err != nil {
		s.rollback()
		return err
	}

	err = appPkg.InstallPackage()
	if err != nil {
		return errors.New("error when installing package")
	}

	fs.GoGenerateRun(s.Info.ProjectPath)
	fs.GoFormat(s.Info.ProjectPath, s.Info.GoModName)

	logger.Infoln(fmt.Sprintf("Your App '%s' already created", s.Info.GoModName))

	return nil
}

func (s AppImpl) createMainApp() error {
	logger.Info("Create main.go file... ")

	templateNew := misctemplate.NewMainTemplate()
	mainTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}
	mainAbstractCode := ast.NewAbstractCode(mainTemplate, parser.ParseComments)
	mainAbstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s/internal/logger\"", s.Info.GoModName),
	})
	mainAbstractCode.AddCommentOutsideFunction(ast.Comment{
		Value: "//go:generate go run github.com/google/wire/cmd/wire",
	})
	mainAbstractCode.AddCommentOutsideFunction(ast.Comment{
		Value: "//go:generate go run github.com/swaggo/swag/cmd/swag init",
	})
	err = mainAbstractCode.RebuildCode()
	if err != nil {
		return err
	}

	mainTemplate = mainAbstractCode.GetCode()

	err = fs.CreateFileIfNotExist(s.Info.ProjectPath, "main.go", mainTemplate)
	if err != nil {
		return err
	}

	logger.InfoSuccessln("success")
	return nil
}

func (s AppImpl) rollback() {
	fs.DeleteFolder(s.Info.ProjectPath)
}
