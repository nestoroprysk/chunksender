package main

import (
	"context"
	"log"

	"sender/cmd"
	"sender/config"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func main() {
	env := &config.Env{
		Fs: afero.NewOsFs(),
		BackgroundCtx: context.Background(),
	}

	root := cmd.Root(env, config.InitEnvFromOptions)
	// TODO: make start private private
	root.AddCommand(cmd.Start(env), cmd.Stop(env))
	for _, cmd := range []*cobra.Command{
		cmd.StartPrivate(env),
	} {
		cmd.Hidden = true
		root.AddCommand(cmd)
	}

	if err := root.Execute(); err != nil {
		// TODO: usual error
		log.Fatal(err)
	}
}
