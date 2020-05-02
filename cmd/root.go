package cmd

import (
	"github.com/nestoroprysk/chunksender/config"
	"github.com/nestoroprysk/chunksender/constant"

	"github.com/spf13/cobra"
)

func Root(env *config.Env, initEnv func(*config.Env, *config.EnvOptions) error) *cobra.Command {
	options := &config.EnvOptions{}

	cmd := &cobra.Command{
		Use:   "spam",
		Short: "Spam and annoy everyone",

		SilenceUsage:  true,
		SilenceErrors: true,

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initEnv(env, options)
		},

		// TODO: some run
	}

	// TODO: use util flag for path
	cmd.PersistentFlags().StringVar(&options.ConfigPath, "config", constant.DefaultConfigPath, "set a config path")

	return cmd
}
