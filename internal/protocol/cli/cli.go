package cli

import (
	"fmt"
	"os"

	"github.com/nurcahyaari/kite/config"
	"github.com/nurcahyaari/kite/internal/logger"
	clirouter "github.com/nurcahyaari/kite/src/protocol/cli"

	cli "github.com/urfave/cli/v2"
)

type CliImpl struct {
	cli       *cli.App
	clirouter clirouter.CliRouter
}

func NewCliApp(
	clir clirouter.CliRouter,
	cliapp *cli.App,
) *CliImpl {
	return &CliImpl{
		cli:       cliapp,
		clirouter: clir,
	}
}

func (s CliImpl) CreateNewCliApp() {
	var err error

	s.cli.Name = config.Get().Application.Name
	s.cli.Description = config.Get().Application.Description
	s.cli.Version = config.Get().Application.Version
	path, err := os.Getwd()
	if err != nil {
		logger.Errorln(err)
	}

	s.cli.Commands = []*cli.Command{
		{
			Name:        "new",
			Description: "make new Apps",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Value: "kite",
					Usage: "creating new Apps with name",
				},
				&cli.StringFlag{
					Name:  "path",
					Value: "",
					Usage: "Path of projects",
				},
			},
			Action: func(ctx *cli.Context) error {
				return s.clirouter.CreateNewApps(ctx, path)
			},
		},
		{
			Name:        "domain",
			Description: "Make a new domain",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Value: "",
					Usage: "domain name",
				},
				&cli.StringFlag{
					Name:  "path",
					Value: "",
					Usage: "Path of projects",
				},
				&cli.BoolFlag{
					Name:    "create-only-folder",
					Aliases: []string{"cof"},
					Usage:   "To only domain folder",
				},
			},
			Action: func(ctx *cli.Context) error {
				return s.clirouter.CreateNewDomain(ctx, path)
			},
		},
		{
			Name:        "handler",
			Description: "Make a new handler",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Value: "",
					Usage: "domain name",
				},
				&cli.StringFlag{
					Name:  "path",
					Value: "",
					Usage: "Path of projects",
				},
				&cli.StringFlag{
					Name:  "protocol-type",
					Value: "",
					Usage: "define what is the handler's protocol type such as (http, grpc, or empty)",
				},
			},
			Action: func(ctx *cli.Context) error {
				return s.clirouter.CreateNewHandler(ctx, path)
			},
		},
	}

	err = s.cli.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
