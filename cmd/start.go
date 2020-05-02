package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/nestoroprysk/chunksender/config"
	"github.com/nestoroprysk/chunksender/constant"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func Start(env *config.Env) *cobra.Command {
	result := &cobra.Command{
		Use:   "start",
		Short: "Start spamming and annoying everyone",

		RunE: func(cmd *cobra.Command, args []string) error {
			self, err := os.Executable()
			if err != nil {
				return err
			}

			// TODO: check if it is ok to do that
			c := exec.Command(self, "start-private", "--config", env.ConfigPath)
			if err := c.Start(); err != nil {
				return err
			}

			// TODO: direct convert
			pid := strconv.Itoa(c.Process.Pid)
			if err = afero.WriteFile(env.Fs, constant.PIDFilePath, []byte(pid), 0644); err != nil {
				return err
			}

			// TODO: nodmal logging
			fmt.Printf("Started spammer at the pid %v\n", pid)

			// TODO: return error of type
			return nil
		},
	}

	return result
}
