package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/nurcahyaari/kite/config"
	"github.com/nurcahyaari/kite/internal/logger"
	clirouter "github.com/nurcahyaari/kite/src/protocol/cli"
	"github.com/theckman/yacspin"

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

	spinner, err := yacspin.New(yacspin.Config{
		Frequency:         100 * time.Millisecond,
		CharSet:           yacspin.CharSets[57],
		SuffixAutoColon:   true,
		StopCharacter:     "✔",
		StopFailCharacter: "✘",
		StopColors:        []string{"fgGreen"},
		StopFailColors:    []string{"fgRed"},
	})
	if err != nil {
		return
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
				spinner.Message(" Creating Application ")
				spinner.StopMessage(" New Application was created ")
				spinner.Start()
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
				spinner.Message(" Creating Domain ")
				spinner.StopMessage(" New Domain was created ")
				spinner.Start()
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
				spinner.Message(" Creating Handler ")
				spinner.StopMessage(" New Handler was created ")
				spinner.Start()
				return s.clirouter.CreateNewHandler(ctx, path)
			},
		},
		{
			Name:        "module",
			Description: "Make a new module",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Value: "",
					Usage: "module name",
				},
				&cli.StringFlag{
					Name:  "path",
					Value: "",
					Usage: "Path of a new module",
				},
			},
			Action: func(ctx *cli.Context) error {
				spinner.Message(" Creating Module ")
				spinner.StopMessage(" New Module was created ")
				spinner.Start()
				return s.clirouter.CreateNewModule(ctx, path)
			},
		},
		{
			Name:        "module",
			Description: "Make a new module",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Value: "",
					Usage: "module name",
				},
				&cli.StringFlag{
					Name:  "path",
					Value: "",
					Usage: "Path of a new module",
				},
			},
			Action: func(ctx *cli.Context) error {
				spinner.Message(" Creating Module ")
				spinner.StopMessage(" New Module was created ")
				spinner.Start()
				return s.clirouter.CreateNewModule(ctx, path)
			},
		},
		{
			Name:        "version",
			Description: "get app version",
			Aliases: []string{
				"v",
			},
			Action: func(ctx *cli.Context) error {
				fmt.Printf("v%s \n", config.Get().Application.Version)
				return nil
			},
		},
	}

	err = s.cli.Run(os.Args)
	if err != nil {
		spinner.StopFailMessage(fmt.Sprintf(" %s ", err.Error()))
		spinner.StopFail()
	}

	spinner.Stop()
}
