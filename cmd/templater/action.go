package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/viper"
	"github.com/webhippie/templater/pkg/config"
	"gopkg.in/urfave/cli.v2"
)

// Action is the general handle for this CLI tool.
func Action() cli.ActionFunc {
	return func(c *cli.Context) error {
		if c.Args().Len() == 0 {
			cli.ShowAppHelp(c)
			os.Exit(1)
		} else {
			config.Source = c.Args().First()
		}

		if config.Prefix != "" {
			viper.SetEnvPrefix(config.Prefix)
		}

		viper.AutomaticEnv()

		tmpl, err := template.New(
			path.Base(config.Source),
		).Funcs(
			template.FuncMap{
				"envGet":      viper.Get,
				"envBool":     viper.GetBool,
				"envFloat":    viper.GetFloat64,
				"envInt":      viper.GetInt,
				"envString":   viper.GetString,
				"envTime":     viper.GetTime,
				"envDuration": viper.GetDuration,
				"envIsSet":    viper.IsSet,
				"base":        path.Base,
				"dir":         path.Dir,
				"datetime":    time.Now,
				"split":       strings.Split,
				"join":        strings.Join,
				"toUpper":     strings.ToUpper,
				"toLower":     strings.ToLower,
				"contains":    strings.Contains,
				"replace":     strings.Replace,
			},
		).ParseFiles(
			config.Source,
		)

		if err != nil {
			fmt.Println("Failed to read template")

			if config.Debug {
				fmt.Println(err)
			}

			os.Exit(1)
		}

		var handle io.Writer

		if config.Output == "" {
			handle = os.Stdout
		} else {
			var err error

			handle, err = os.Create(
				config.Output,
			)

			if err != nil {
				fmt.Println("Failed to write config")

				if config.Debug {
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
			&struct {
				Prefix string
				Source string
				Output string
				Debug  bool
			}{
				Prefix: config.Prefix,
				Source: config.Source,
				Output: config.Output,
				Debug:  config.Debug,
			},
		)

		if exe != nil {
			fmt.Println("Failed to parse template")

			if config.Debug {
				fmt.Println(exe)
			}

			os.Exit(3)
		}

		return nil
	}
}
