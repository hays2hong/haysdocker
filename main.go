package main

import (
	"os"
	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

const usage = "it's a usage..."

func main() {
	app := cli.NewApp()
	app.Name = "haysdocker"
	app.Usage = usage

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
