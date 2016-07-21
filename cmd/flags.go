package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/webhippie/templater/config"
)

// Flags defines all available flags for this command.
func Flags() []cli.Flag {
	return []cli.Flag{
		cli.BoolTFlag{
			Name:        "update, u",
			Usage:       "Enable auto updates",
			EnvVar:      "TEMPLATER_UPDATE",
			Destination: &config.Update,
		},
		cli.StringFlag{
			Name:        "prefix, p",
			Usage:       "Prefix of the environemt variables",
			EnvVar:      "TEMPLATER_PREFIX",
			Destination: &config.Prefix,
			Value:       "",
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "Different output than STDOUT",
			EnvVar:      "TEMPLATER_OUTPUT",
			Destination: &config.Output,
			Value:       "",
		},
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Print a bit more debug output",
			EnvVar:      "TEMPLATER_DEBUG",
			Destination: &config.Debug,
		},
	}
}
