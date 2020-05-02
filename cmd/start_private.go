package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/nestoroprysk/chunksender/config"
	"github.com/nestoroprysk/chunksender/parser"
	"github.com/nestoroprysk/chunksender/sender"

	"github.com/spf13/cobra"
)

const (
	bookPath = "Letters.txt"
)

func StartPrivate(env *config.Env) *cobra.Command {
	result := &cobra.Command{
		Use: "start-private",
		RunE: func(cmd *cobra.Command, args []string) error {
			letters, err := parser.Parse(bookPath)
			if err != nil {
				return err
			}

			// TODO: maybe exit with context
			// TODO: provide flag which changes the interval
			senders := sender.NewSenders(env.Cfg)
			t := time.Tick(4 * time.Hour)
			i := 0

			send := func() {
				letter := letters[i]
				body := strings.Join(letter.Body, "<br><br>")

				for _, s := range senders {
					err := s.Send(letter.Title, body)
					// TODO: log in some nice way
					if err != nil {
						fmt.Errorf("Failed to send: %q (err: %v)\n", body, err)
						continue
					}

					fmt.Printf("Sent: %q\n", body)
				}

				i++
			}
			send()

			for _ = range t {
				if i >= len(letters) {
					break
				}

				send()
			}

			// TODO: log in a normal way
			fmt.Print("Exiting")

			// TODO: exit in some way
			return nil
		},
	}

	return result
}
