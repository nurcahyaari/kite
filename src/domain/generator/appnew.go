package generator

import (
	"errors"
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/templates/misctemplate"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/internal/utils/ast"
	"github.com/nurcahyaari/kite/internal/utils/pkg"
	"github.com/nurcahyaari/kite/src/domain/configgen"
	"github.com/nurcahyaari/kite/src/domain/dbgen"
	"github.com/nurcahyaari/kite/src/domain/domaingen"
	"github.com/nurcahyaari/kite/src/domain/envgen"
	"github.com/nurcahyaari/kite/src/domain/infrastructuregen"
	"github.com/nurcahyaari/kite/src/domain/internalgen"
	"github.com/nurcahyaari/kite/src/domain/miscgen"
	"github.com/nurcahyaari/kite/src/domain/protocolgen"
	"github.com/nurcahyaari/kite/src/domain/srcgen"
	"github.com/nurcahyaari/kite/src/domain/wiregen"
)

type AppGenNew interface {
	CreateNewApp(dto AppNewDto) error
	createMainApp(dto AppNewDto) error
	rollback(path string)
}

type AppGenNewImpl struct {
	fs                database.FileSystem
	configGen         configgen.ConfigGen
	envGen            envgen.EnvGen
	wireGen           wiregen.WireGen
	internalGen       internalgen.InternalGen
	infrastructureGen infrastructuregen.InfrastructureGen
	srcGen            srcgen.SrcGen
	domainGen         domaingen.DomainGen
	gitignoreGen      miscgen.GitIgnoreGen
	makefileGen       miscgen.MakefileGen
}

func NewApp(
	fs database.FileSystem,
	configGen configgen.ConfigGen,
	envGen envgen.EnvGen,
	wireGen wiregen.WireGen,
	internalGen internalgen.InternalGen,
	infrastructureGen infrastructuregen.InfrastructureGen,
	srcGen srcgen.SrcGen,
	domainGen domaingen.DomainGen,
	gitignoreGen miscgen.GitIgnoreGen,
	makefileGen miscgen.MakefileGen,
) *AppGenNewImpl {
	return &AppGenNewImpl{
		fs:                fs,
		configGen:         configGen,
		envGen:            envGen,
		wireGen:           wireGen,
		internalGen:       internalGen,
		infrastructureGen: infrastructureGen,
		srcGen:            srcGen,
		domainGen:         domainGen,
		gitignoreGen:      gitignoreGen,
		makefileGen:       makefileGen,
	}
}

func (s AppGenNewImpl) CreateNewApp(dto AppNewDto) error {
	// creational flow
	// create top level file such as (.env, .env.example, main.go, wire.go, gitignore, gitinit, go.mod)
	// create config modules
	// create internal modules such as (protocols, logger, utils)
	// create infrastructure modules
	// create src moduels
	if !s.fs.IsFolderEmpty(dto.ProjectPath) && s.fs.IsFolderExists(dto.ProjectPath) {
		return errors.New("the folder is not empty")
	}

	if utils.IsFolderHasGoMod(dto.ProjectPath) {
		return errors.New("the folder already had go.mod")
	}

	s.fs.CreateFolderIfNotExists(dto.ProjectPath)

	// init go.mod
	err := utils.GoModInit(dto.ProjectPath, dto.GoModName)
	if err != nil {
		return err
	}

	err = utils.Gitinit(dto.ProjectPath)
	if err != nil {
		return err
	}

	// setup all path
	configPath := utils.ConcatDirPath(dto.ProjectPath, "config")
	internalPath := utils.ConcatDirPath(dto.ProjectPath, "internal")
	infrastructurePath := utils.ConcatDirPath(dto.ProjectPath, "infrastructure")
	srcPath := utils.ConcatDirPath(dto.ProjectPath, "src")

	// create file in project path dir level
	err = s.envGen.CreateEnvFile(dto.ProjectPath)
	if err != nil {
		return err
	}

	err = s.envGen.CreateEnvExampleFile(dto.ProjectPath)
	if err != nil {
		return err
	}

	err = s.wireGen.CreateWireFiles(wiregen.WireDto{
		ProjectPath: dto.ProjectPath,
		GomodName:   dto.GoModName,
	})
	if err != nil {
		return err
	}

	miscDto := miscgen.MiscDto{
		ProjectPath: dto.ProjectPath,
		GomodName:   dto.GoModName,
	}
	err = s.gitignoreGen.CreateGitignoreFile(miscDto)
	if err != nil {
		return err
	}

	err = s.makefileGen.CreateMakefilefile(miscDto)
	if err != nil {
		return err
	}

	err = s.createMainApp(dto)
	if err != nil {
		return err
	}

	// create config module
	configGenDto := configgen.ConfigDto{
		ConfigPath: configPath,
		AppName:    dto.Name,
	}
	err = s.configGen.CreateConfigDir(configGenDto)
	if err != nil {
		return err
	}

	err = s.configGen.CreateConfigFile(configGenDto)
	if err != nil {
		return err
	}

	// create internal module
	internalDto := internalgen.InternalDto{
		Path:        internalPath,
		ProjectPath: dto.ProjectPath,
		GomodName:   dto.GoModName,
	}
	err = s.internalGen.CreateInternalDir(internalDto)
	if err != nil {
		return err
	}

	err = s.internalGen.CreateInternalModules(internalDto)
	if err != nil {
		return err
	}

	// create infrastructure module
	infrastructureDto := infrastructuregen.InfrastructureDto{
		GomodName:          dto.GoModName,
		DatabaseType:       dbgen.DbMysql,
		InfrastructurePath: infrastructurePath,
		ProjectPath:        dto.ProjectPath,
	}
	err = s.infrastructureGen.CreateInfrastructureDir(infrastructureDto)
	if err != nil {
		return err
	}

	err = s.infrastructureGen.GenerateInfrastructure(infrastructureDto)
	if err != nil {
		return err
	}

	// create src module
	srcDto := srcgen.SrcDto{
		Path:         srcPath,
		GomodName:    dto.GoModName,
		ProtocolType: protocolgen.NewProtocolType(dto.ProtocolType),
	}
	err = s.srcGen.CreateSrcDirectory(srcDto)
	if err != nil {
		return err
	}

	err = s.installPackage()
	if err != nil {
		return err
	}

	return nil
}

func (s AppGenNewImpl) createMainApp(dto AppNewDto) error {
	templateNew := misctemplate.NewMainTemplate()
	mainTemplate, err := templateNew.Render()
	if err != nil {
		return err
	}

	// TODO: remove this
	// since I don't know how to add new line with ast standard lib, so I use this
	mainTemplate = fmt.Sprintf("\n%s", mainTemplate)

	mainAbstractCode := ast.NewAbstractCode(mainTemplate, parser.ParseComments)
	mainAbstractCode.AddImport(ast.ImportSpec{
		Path: "\"context\"",
	})
	mainAbstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s/config\"", dto.GoModName),
	})
	mainAbstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s/internal/graceful\"", dto.GoModName),
	})
	mainAbstractCode.AddImport(ast.ImportSpec{
		Path: fmt.Sprintf("\"%s/internal/logger\"", dto.GoModName),
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

	err = s.fs.CreateFileIfNotExists(dto.ProjectPath, "main.go", mainTemplate)
	if err != nil {
		return err
	}

	return nil
}

func (s AppGenNewImpl) installPackage() error {
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
			"github.com/nurcahyaari/sqlabst",
		},
	}

	return appPkg.InstallPackage()
}

func (s AppGenNewImpl) rollback(path string) {
	s.fs.DeleteFolder(path)
}
