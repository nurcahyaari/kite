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
	CreateNewHandler(ctx *cli.Context, path string) error
	CreateNewModule(ctx *cli.Context, path string) error
}

type CliRouterImpl struct {
	appGenerator generator.AppGenerator
}

func NewCliRouter(
	appGenerator generator.AppGenerator,
) *CliRouterImpl {
	return &CliRouterImpl{
		appGenerator: appGenerator,
	}
}

func (s CliRouterImpl) CreateNewApps(ctx *cli.Context, path string) error {
	gomodName := ctx.String("name")
	protocolType := ctx.String("protocolType")
	if ctx.String("path") != "" {
		path = fmt.Sprintf("%s/", ctx.String("path"))
	}

	dto := generator.AppNewDto{
		ProjectInfo: generator.ProjectInfo{
			GoModName:    gomodName,
			ProjectPath:  utils.ConcatDirPath(path, gomodName),
			ProtocolType: protocolType,
			Name:         gomodName,
		},
	}

	if err := dto.Validate(); err != nil {
		return err
	}

	err := s.appGenerator.CreateNewApp(dto)

	return err
}

func (s CliRouterImpl) CreateNewDomain(ctx *cli.Context, path string) error {
	if ctx.String("path") != "" {
		path = fmt.Sprintf("%s/", ctx.String("path"))
	}
	moduleName := ctx.String("name")
	isCreateDomainFolderOnly := ctx.Bool("create-only-folder")

	gomodName := utils.GetGoModName(path)
	domainDto := generator.DomainNewDto{
		ProjectInfo: generator.ProjectInfo{
			GoModName:   gomodName,
			Name:        moduleName,
			ProjectPath: path,
		},
		IsCreateDomainFolderOnly: isCreateDomainFolderOnly,
	}

	if err := domainDto.Validate(); err != nil {
		return err
	}

	err := s.appGenerator.CreateNewDomain(domainDto)

	return err
}

func (s CliRouterImpl) CreateNewHandler(ctx *cli.Context, path string) error {
	// TODO: write code to implement handler here
	// handler can be http, grpc, amqp, or whatever you want
	// handler just a interface and the impl

	if ctx.String("path") != "" {
		path = fmt.Sprintf("%s/", ctx.String("path"))
	}
	moduleName := ctx.String("name")
	gomodName := utils.GetGoModName(path)

	handlerDto := generator.HandlerNewDto{
		ProjectInfo: generator.ProjectInfo{
			GoModName:   gomodName,
			Name:        moduleName,
			ProjectPath: path,
		},
	}

	if err := handlerDto.Validate(); err != nil {
		return err
	}

	err := s.appGenerator.CreateNewHandler(handlerDto)

	return err
}

func (s CliRouterImpl) CreateNewModule(ctx *cli.Context, path string) error {
	// all of dependency is a module, module = dependency in this code generator

	if ctx.String("path") != "" {
		path = fmt.Sprintf("%s/", ctx.String("path"))
	}
	moduleName := ctx.String("name")

	moduleDto := generator.ModuleNewDto{
		ProjectInfo: generator.ProjectInfo{
			Name:        moduleName,
			ProjectPath: path,
		},
	}

	if err := moduleDto.Validate(); err != nil {
		return err
	}

	err := s.appGenerator.CreateNewModule(moduleDto)

	return err
}
