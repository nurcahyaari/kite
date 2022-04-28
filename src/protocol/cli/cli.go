package cli

import (
	"fmt"

	"github.com/nurcahyaari/kite/internal/utils"
	"github.com/nurcahyaari/kite/src/domain/generator"

	cli "github.com/urfave/cli/v2"
)

type CliRouter interface {
	CreateNewApps(ctx *cli.Context, path string) error
	CreateNewDomain(ctx *cli.Context, path string) error
	// TODO: need to implement
	CreateNewHandler(ctx *cli.Context, path string) error
	CreateNewModule(ctx *cli.Context, path string) error
}

type CliRouterImpl struct {
	path   string
	appGen generator.AppGenNew
}

func NewCliRouter(
	gen generator.AppGenNew,
) *CliRouterImpl {
	return &CliRouterImpl{
		appGen: gen,
	}
}

func (s CliRouterImpl) CreateNewApps(ctx *cli.Context, path string) error {
	gomodName := ctx.String("name")
	protocolType := ctx.String("protocolType")
	s.path = path
	if ctx.String("path") != "" {
		s.path = fmt.Sprintf("%s/", ctx.String("path"))
	}

	err := s.appGen.CreateNewApp(generator.ProjectInfo{
		GoModName:    gomodName,
		ProjectPath:  utils.ConcatDirPath(s.path, gomodName),
		ProtocolType: protocolType,
	})
	// appGenerator := generator.NewAppNew(generator.ProjectInfo{
	// 	GoModName:   gomodName,
	// 	ProjectPath: utils.ConcatDirPath(s.path, gomodName),
	// })

	// err := appGenerator.CreateNewApp()

	return err
}

func (s CliRouterImpl) CreateNewDomain(ctx *cli.Context, path string) error {
	// var err error
	// if ctx.String("path") != "" {
	// 	s.path = fmt.Sprintf("%s/", ctx.String("path"))
	// }
	// moduleName := ctx.String("name")
	// gomodName := utils.GetGoModName(s.path)

	// moduleGen := generator.NewDomain(moduleName, generator.ProjectInfo{
	// 	GoModName:   gomodName,
	// 	ProjectPath: s.path,
	// })
	// err = moduleGen.CreateNewModule()

	// return err
	return nil
}

func (s CliRouterImpl) CreateNewHandler(ctx *cli.Context, path string) error {
	// TODO: write code to implement handler here
	// handler can be http, grpc, amqp, or whatever you want
	// handler just a interface and the impl
	return nil
}

func (s CliRouterImpl) CreateNewModule(ctx *cli.Context, path string) error {
	// TODO: write code to implement create module here
	// all of dependency is a module, module = dependency in this code generator
	return nil
}
