package main

import (
	"bark-cli/commands"
	"bark-cli/flags"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	appName = "bark-cli"
	usage   = "A simple cli for bark"
	version = "1.0.0"
)

func main() {
	app := cli.NewApp()

	setupApplicationMeta(app)
	flags.SetupApplicationFlags(app)
	commands.SetupApplicationCommand(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setupApplicationMeta(app *cli.App) {
	app.Name = appName
	app.Usage = usage
	app.Version = version
}
