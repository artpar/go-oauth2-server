package main

import (
	"log"
	"os"

	"github.com/artpar/go-oauth2-server/commands"
	"github.com/codegangsta/cli"
)

var (
	cliApp *cli.App
)

func init() {
	// Initialise a CLI app
	cliApp = cli.NewApp()
	cliApp.Name = "go-oauth2-server"
	cliApp.Usage = "Go OAuth 2.0 Server"
	cliApp.Author = "Richard Knop"
	cliApp.Email = "risoknop@gmail.com"
	cliApp.Version = "0.0.0"
}

func main() {
	// Set the CLI app commands
	cliApp.Commands = []cli.Command{
		{
			Name:  "migrate",
			Usage: "run migrations",
			Action: func(c *cli.Context) {
				if err := commands.Migrate(); err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:  "loaddata",
			Usage: "load data from fixture",
			Action: func(c *cli.Context) {
				if err := commands.LoadData(c.Args()); err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:  "runserver",
			Usage: "run web server",
			Action: func(c *cli.Context) {
				if err := commands.RunServer(); err != nil {
					log.Fatal(err)
				}
			},
		},
	}

	// Run the CLI app
	cliApp.Run(os.Args)
}
