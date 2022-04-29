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
	path      string
	appGen    generator.AppGenNew
	domainGen generator.DomainGen
}

func NewCliRouter(
	appGen generator.AppGenNew,
	domainGen generator.DomainGen,
) *CliRouterImpl {
	return &CliRouterImpl{
		appGen:    appGen,
		domainGen: domainGen,
	}
}

func (s CliRouterImpl) CreateNewApps(ctx *cli.Context, path string) error {
	gomodName := ctx.String("name")
	protocolType := ctx.String("protocolType")
	s.path = path
	if ctx.String("path") != "" {
		s.path = fmt.Sprintf("%s/", ctx.String("path"))
	}

	dto := generator.AppNewDto{
		ProjectInfo: generator.ProjectInfo{
			GoModName:    gomodName,
			ProjectPath:  utils.ConcatDirPath(s.path, gomodName),
			ProtocolType: protocolType,
		},
	}

	if err := dto.Validate(); err != nil {
		return err
	}

	err := s.appGen.CreateNewApp(dto)

	return err
}

func (s CliRouterImpl) CreateNewDomain(ctx *cli.Context, path string) error {
	// var err error
	if ctx.String("path") != "" {
		s.path = fmt.Sprintf("%s/", ctx.String("path"))
	}
	moduleName := ctx.String("name")
	isDomainFullCreational := ctx.Bool("domain-full-creational")
	gomodName := utils.GetGoModName(s.path)
	domainDto := generator.DomainNewDto{
		ProjectInfo: generator.ProjectInfo{
			GoModName:   gomodName,
			Name:        moduleName,
			ProjectPath: path,
		},
		IsDomainFullCreational: isDomainFullCreational,
	}

	if err := domainDto.Validate(); err != nil {
		return err
	}

	err := s.domainGen.CreateNewDomain(domainDto)

	return err
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
