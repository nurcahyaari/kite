package main

import (
	"fmt"
	"os"

	"github.com/nurcahyaari/kite/lib/generator"

	"github.com/nurcahyaari/kite/utils/fs"
	"github.com/nurcahyaari/kite/utils/logger"
	cli "github.com/urfave/cli/v2"
)

// TODO: beautify the code

func main() {
	var err error
	app := cli.NewApp()
	app.Name = "kite"
	app.Description = "Projects Generator for Golang inspired by Clean Code Arch"

	path, err := os.Getwd()
	if err != nil {
		logger.Errorln(err)
		return
	}

	app.Commands = []*cli.Command{
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
			Action: func(c *cli.Context) error {
				gomodName := c.String("name")
				if c.String("path") != "" {
					path = fmt.Sprintf("%s/", c.String("path"))
				}

				appGenerator := generator.NewAppNew(generator.ProjectInfo{
					GoModName:   gomodName,
					ProjectPath: fs.ConcatDirPath(path, gomodName),
				})

				err = appGenerator.CreateNewApp()

				return err
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
					Usage: "Path of projects",
				},
			},
			Action: func(c *cli.Context) error {
				var err error
				if c.String("path") != "" {
					path = fmt.Sprintf("%s/", c.String("path"))
				}
				moduleName := c.String("name")
				gomodName := fs.GetGoModName(path)

				moduleGen := generator.NewModule(moduleName, generator.ProjectInfo{
					GoModName:   gomodName,
					ProjectPath: path,
				})
				err = moduleGen.CreateNewModule()

				return err
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Errorf("Error ", err)
	}
}
