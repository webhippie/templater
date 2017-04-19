package main

import (
	"os"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/webhippie/templater/cmd"
	"github.com/webhippie/templater/config"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if env := os.Getenv("TEMPLATER_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:      "templater",
		Version:   config.Version,
		Usage:     "A template processor for environment variables",
		Compiled:  time.Now(),
		ArgsUsage: "<inputfile>",

		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},

		Flags:  cmd.Flags(),
		Before: cmd.Before(),
		Action: cmd.Action(),
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
