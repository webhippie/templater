package main

import (
	"github.com/webhippie/templater/pkg/config"
	"gopkg.in/urfave/cli.v2"
)

// Flags defines all available flags for this command.
func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "prefix, p",
			Usage:       "Prefix of the environemt variables",
			EnvVars:     []string{"TEMPLATER_PREFIX"},
			Destination: &config.Prefix,
			Value:       "",
		},
		&cli.StringFlag{
			Name:        "output, o",
			Usage:       "Different output than STDOUT",
			EnvVars:     []string{"TEMPLATER_OUTPUT"},
			Destination: &config.Output,
			Value:       "",
		},
		&cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Print a bit more debug output",
			EnvVars:     []string{"TEMPLATER_DEBUG"},
			Destination: &config.Debug,
		},
	}
}
