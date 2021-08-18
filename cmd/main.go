package main

import (
	"fmt"
	"github.com/aozora0000/idea"
	"github.com/urfave/cli/v2"
	"os"
	"path"
)

var (
	version = "local"
)

func init() {
	err := idea.MakeShelf()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	app := &cli.App{
		Name:                 path.Base(os.Args[0]),
		Usage:                "alias subcommand from file",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:   "path",
				Action: idea.Path,
			},
			{
				Name:   "create",
				Action: idea.Create,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "psyche",
						Aliases:  []string{"P"},
						Required: true,
					},
				},
			},
			{
				Name:    "decrypt",
				Action:  idea.Decrypt,
				Aliases: []string{"dec"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "psyche",
						Aliases:  []string{"P"},
						Required: true,
					},
				},
			},
			{
				Name:    "encrypt",
				Action:  idea.Encrypt,
				Aliases: []string{"enc"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "psyche",
						Aliases:  []string{"P"},
						Required: true,
					},
				},
			},
		},
		Version: version,
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
