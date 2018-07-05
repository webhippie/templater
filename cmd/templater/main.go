package main

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/webhippie/templater/pkg/version"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	if env := os.Getenv("TEMPLATER_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:      "templater",
		Version:   version.Version.String(),
		Usage:     "a template processor for environment variables",
		Compiled:  time.Now(),
		ArgsUsage: "<inputfile>",

		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},

		Flags:  Flags(),
		Before: Before(),
		Action: Action(),
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
