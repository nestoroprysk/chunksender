package cmd

import (
	"fmt"
	"os/exec"

	"chunksender/config"
	"chunksender/constant"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func Stop(env *config.Env) *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stop spamming and annoying everyone",

		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: resolve path
			// TODO: validate pid
			pid, err := afero.ReadFile(env.Fs, constant.PIDFilePath)
			if err != nil {
				return err
			}

			c := exec.Command("kill", string(pid))
			if err = c.Run(); err != nil {
				return err
			}

			// TODO: normal logging
			fmt.Printf("Stopped spammer at the pid %v\n", string(pid))

			// TODO: provide type
			return nil
		},
	}
}
