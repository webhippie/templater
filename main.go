package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var (
	version string
)

func main() {
	app := cli.NewApp()
	app.Name = "templater"
	app.Version = version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "A template processor for environment variables"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "output",
			Usage:  "Different output than STDOUT",
			EnvVar: "TEMPLATER_OUTPUT",
			Value:  "",
		},
		cli.BoolFlag{
			Name:   "force",
			Usage:  "Force to overwrite the output",
			EnvVar: "TEMPLATER_FORCE",
		},
	}

	app.Action = func(c *cli.Context) {
		fmt.Println("Hello!")
	}

	app.Run(os.Args)
}
