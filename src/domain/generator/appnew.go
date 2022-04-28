package generator

import (
	"fmt"
	"go/parser"

	"github.com/nurcahyaari/kite/infrastructure/database"
	"github.com/nurcahyaari/kite/internal/logger"
	"github.com/nurcahyaari/kite/internal/templates/misctemplate"
	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/internal/utils/ast"
	"github.com/nurcahyaari/kite/src/domain/configgen"
	"github.com/nurcahyaari/kite/src/domain/dbgen"
	"github.com/nurcahyaari/kite/src/domain/envgen"
	"github.com/nurcahyaari/kite/src/domain/infrastructuregen"
	"github.com/nurcahyaari/kite/src/domain/internalgen"
	"github.com/nurcahyaari/kite/src/domain/misc"
	"github.com/nurcahyaari/kite/src/domain/protocolgen"
	"github.com/nurcahyaari/kite/src/domain/srcgen"
)

type AppGenNew interface {
	CreateNewApp(info ProjectInfo) error
	// Private
	createMainApp(info ProjectInfo) error
	rollback(path string)
}

type AppGenNewImpl struct {
	fs                database.FileSystem
	configGen         configgen.ConfigGen
	envGen            envgen.EnvGen
	wireGen           misc.WireGen
	internalGen       internalgen.InternalGen
	infrastructureGen infrastructuregen.InfrastructureGen
	srcGen            srcgen.SrcGen
}

func NewApp(
	fs database.FileSystem,
	configGen configgen.ConfigGen,
	envGen envgen.EnvGen,
	wireGen misc.WireGen,
	internalGen internalgen.InternalGen,
	infrastructureGen infrastructuregen.InfrastructureGen,
	srcGen srcgen.SrcGen,
) *AppGenNewImpl {
	return &AppGenNewImpl{
		fs:                fs,
		configGen:         configGen,
		envGen:            envGen,
		wireGen:           wireGen,
		internalGen:       internalGen,
		infrastructureGen: infrastructureGen,
		srcGen:            srcGen,
	}
}

func (s AppGenNewImpl) CreateNewApp(info ProjectInfo) error {
	s.fs.CreateFolderIfNotExists(info.ProjectPath)

	// setup all path
	configPath := utils.ConcatDirPath(info.ProjectPath, "config")
	internalPath := utils.ConcatDirPath(info.ProjectPath, "internal")
	infrastructurePath := utils.ConcatDirPath(info.ProjectPath, "infrastructure")
	srcPath := utils.ConcatDirPath(info.ProjectPath, "src")

	// create config module
	configGenDto := configgen.ConfigDto{
		ConfigPath: configPath,
		AppName:    info.Name,
	}
	err := s.configGen.CreateConfigDir(configGenDto)
	if err != nil {
		return err
	}

	err = s.configGen.CreateConfigFile(configGenDto)
	if err != nil {
		return err
	}

	// create internal module
	internalDto := internalgen.InternalDto{
		Path:      internalPath,
		GomodName: info.GoModName,
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
		GomodName:          info.GoModName,
		DatabaseType:       dbgen.DbMysql,
		InfrastructurePath: infrastructurePath,
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
		GomodName:    info.GoModName,
		ProtocolType: protocolgen.NewProtocolType(info.ProtocolType),
	}
	err = s.srcGen.CreateSrcDirectory(srcDto)
	if err != nil {
		return err
	}

	// create file in project path dir level
	s.envGen.CreateEnvFile(info.ProjectPath)
	s.envGen.CreateEnvExampleFile(info.ProjectPath)

	s.wireGen.CreateWireFiles(misc.MiscDto{
		ProjectPath: info.ProjectPath,
		GomodName:   info.GoModName,
	})

	err = s.createMainApp(info)
	if err != nil {
		return err
	}
	return nil
}

func (s AppGenNewImpl) createMainApp(info ProjectInfo) error {
	logger.Info("Create main.go file... ")

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
		Path: fmt.Sprintf("\"%s/internal/logger\"", info.GoModName),
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

	err = s.fs.CreateFileIfNotExists(info.ProjectPath, "main.go", mainTemplate)
	if err != nil {
		return err
	}

	logger.InfoSuccessln("success")
	return nil
}

func (s AppGenNewImpl) rollback(path string) {
	s.fs.DeleteFolder(path)
}
