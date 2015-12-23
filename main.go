package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
)

var (
	version string
	prefix  string
	source  string
	output  string
	debug   bool
)

type Payload struct {
	Prefix string
	Source string
	Output string
	Debug  bool
}

func main() {
	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Show the help, so what you see now",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Print the current version of that tool",
	}

	app := cli.NewApp()
	app.Name = "templater"
	app.Version = version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "A template processor for environment variables"
	app.ArgsUsage = "<inputfile>"

	app.HideHelp = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "prefix, p",
			Usage:       "Prefix of the environemt variables",
			EnvVar:      "TEMPLATER_PREFIX",
			Destination: &prefix,
			Value:       "",
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "Different output than STDOUT",
			EnvVar:      "TEMPLATER_OUTPUT",
			Destination: &output,
			Value:       "",
		},
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Print a bit more debug output",
			EnvVar:      "TEMPLATER_DEBUG",
			Destination: &debug,
		},
		cli.HelpFlag,
	}

	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 0 {
			cli.ShowAppHelp(c)
			os.Exit(1)
		} else {
			source = c.Args().First()
		}

		if prefix != "" {
			viper.SetEnvPrefix(prefix)
		}

		viper.AutomaticEnv()

		tmpl, err := template.New(
			path.Base(source),
		).Funcs(
			template.FuncMap{
				"Get":                viper.Get,
				"GetBool":            viper.GetBool,
				"GetFloat64":         viper.GetFloat64,
				"GetInt":             viper.GetInt,
				"GetString":          viper.GetString,
				"GetStringMap":       viper.GetStringMap,
				"GetStringMapString": viper.GetStringMapString,
				"GetTime":            viper.GetTime,
				"GetDuration":        viper.GetDuration,
				"IsSet":              viper.IsSet,
			},
		).ParseFiles(
			source,
		)

		if err != nil {
			fmt.Println("Failed to read template")

			if debug {
				fmt.Println(err)
			}

			os.Exit(1)
		}

		var handle io.Writer

		if output == "" {
			handle = os.Stdout
		} else {
			var err error

			handle, err = os.Create(
				output,
			)

			if err != nil {
				fmt.Println("Failed to write config")

				if debug {
					fmt.Println(err)
				}

				os.Exit(2)
			}
		}

		writer := bufio.NewWriter(
			handle,
		)

		defer writer.Flush()

		exe := tmpl.Execute(
			writer,
			&Payload{
				Prefix: prefix,
				Source: source,
				Output: output,
				Debug:  debug,
			},
		)

		if exe != nil {
			fmt.Println("Failed to parse template")

			if debug {
				fmt.Println(exe)
			}

			os.Exit(3)
		}
	}

	app.Run(os.Args)
}
