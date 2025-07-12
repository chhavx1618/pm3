package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"pm3/manager"
)

func main() {
	app := &cli.App{
		Name:    "pm3",
		Usage:   "my own pm2 lol",
		Version: "1.0.0",
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "Start a new process",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return fmt.Errorf("missing <file> argument")
					}
					manager.Start(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "List all processes",
				Action: func(c *cli.Context) error {
					manager.List()
					return nil
				},
			},
			{
				Name:  "stop",
				Usage: "Stop a process",
				Action: func(c *cli.Context) error {
					manager.Stop(c.Args().First())
					return nil
				},
			},
			{
				Name:  "restart",
				Usage: "Restart a process",
				Action: func(c *cli.Context) error {
					manager.Restart(c.Args().First())
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "Delete a process",
				Action: func(c *cli.Context) error {
					manager.Delete(c.Args().First())
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
