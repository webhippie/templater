package command

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	generateCmd = &cobra.Command{
		Use:     "generate [template]",
		Aliases: []string{"gen"},
		Short:   "Generate output from template",
		Run:     generateAction,
		Args: func(_ *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("missing template argument")
			}

			return nil
		},
	}

	defaultGeneratePrefix = ""
	defaultGenerateOutput = ""
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.PersistentFlags().String("prefix", defaultGeneratePrefix, "Prefix of the env variables")
	viper.SetDefault("generate.prefix", defaultGeneratePrefix)
	viper.BindPFlag("generate.prefix", generateCmd.PersistentFlags().Lookup("prefix"))

	generateCmd.PersistentFlags().String("output", defaultGenerateOutput, "Different output than stdout")
	viper.SetDefault("generate.output", defaultGenerateOutput)
	viper.BindPFlag("generate.output", generateCmd.PersistentFlags().Lookup("output"))
}

func generateAction(_ *cobra.Command, args []string) {
	v := viper.New()
	source := args[0]

	if viper.GetString("prefix") != "" {
		v.SetEnvPrefix(
			viper.GetString("prefix"),
		)
	}

	v.AutomaticEnv()

	tmpl, err := template.New(
		path.Base(source),
	).Funcs(
		template.FuncMap{
			"envGet":      v.Get,
			"envBool":     v.GetBool,
			"envFloat":    v.GetFloat64,
			"envInt":      v.GetInt,
			"envString":   v.GetString,
			"envTime":     v.GetTime,
			"envDuration": v.GetDuration,
			"envIsSet":    v.IsSet,
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
		source,
	)

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to read template")

		os.Exit(1)
	}

	var handle io.Writer

	if viper.GetString("output") == "" {
		handle = os.Stdout
	} else {
		var (
			err error
		)

		handle, err = os.Create(
			viper.GetString("output"),
		)

		if err != nil {
			log.Error().
				Err(err).
				Msg("failed to write config")

			os.Exit(2)
		}
	}

	writer := bufio.NewWriter(
		handle,
	)

	defer writer.Flush()

	if exe := tmpl.Execute(
		writer,
		&struct {
			Source string
			Prefix string
			Output string
		}{
			Source: source,
			Prefix: viper.GetString("prefix"),
			Output: viper.GetString("output"),
		},
	); exe != nil {
		log.Error().
			Err(err).
			Msg("failed to parse template")

		os.Exit(3)
	}
}
