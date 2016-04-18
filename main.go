package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/webhippie/templater/cmd"
	"github.com/webhippie/templater/config"
)

var (
	version    string
	versionSha string
)

func main() {
	app := cli.NewApp()
	app.Name = "templater"
	app.Version = config.Version
	app.Usage = "A template processor for environment variables"
	app.ArgsUsage = "<inputfile>"

	app.Authors = []cli.Author{
		{"Thomas Boerger", "thomas@webhippie.de"},
	}

	app.HideHelp = true

	app.Before = cmd.Before()
	app.Flags = cmd.Flags()
	app.Action = cmd.Action()

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Show the help, so what you see now",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Print the current version of that tool",
	}

	app.Run(os.Args)
}
