package main

import (
	"fmt"
	"github.com/urfave/cli"
)

func StartCli(args []string) ([]string, error) {
	var resp []string
	var err error
	app := cli.NewApp()
	app.Name = "restwell-cli"
	app.Version = "0.1.0"
	app.Usage = "Create simple rest servers from a json file. Profit."
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Tyler Boright",
			Email: "ru.lai.development@gmail.com",
		},
	}

	app.Action = func(c *cli.Context) error {
		resp = append(resp, "Welcome to restwell. type help for usage examples.")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "createRestServer",
			Aliases: []string{"crs"},
			Usage:   "Creates a rest server from a json file stored on disk",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	app.Run(args)

	return resp, err
}
