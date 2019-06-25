package main

import (
	"errors"
	"fmt"
	pf "github.com/ru-lai/pathfinder"
	rw "github.com/ru-lai/restwell"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
)

var addr = "127.0.0.1:12345"

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
				if pf.DoesExist(c.Args().First()) == false {
					err = errors.New(fmt.Sprintf("The file path \"%s\" does not exist. Enter a valid path and try again", c.Args().First()))
					return err
				}

				jStr, err := ioutil.ReadFile(c.Args().First())
				if err != nil {
					return err
				}

				mux := rw.CreateMuxFromJSON(string(jStr))
				s := http.Server{Handler: mux, Addr: addr}
				fmt.Printf("Now listening on %s...\n", addr)
				log.Fatal(s.ListenAndServe())

				return nil
			},
		},
	}

	app.Run(args)

	return resp, err
}
