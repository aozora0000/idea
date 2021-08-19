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
		Usage:                "encrypt decrypt tool from psyche",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:   "path",
				Action: idea.Path,
				Usage:  "display psyches directory",
			},
			{
				Name:   "create",
				Action: idea.Create,
				Usage:  "create psyche",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "psyche",
						Aliases:  []string{"P"},
						Required: true,
						Usage:    "psyche name",
					},
				},
			},
			{
				Name:    "decrypt",
				Action:  idea.Decrypt,
				Usage:   "decrypt string by psyche",
				Aliases: []string{"dec"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "psyche",
						Aliases:  []string{"P"},
						Required: true,
						Usage:    "psyche name",
					},
				},
			},
			{
				Name:    "encrypt",
				Action:  idea.Encrypt,
				Usage:   "encrypt string by psyche",
				Aliases: []string{"enc"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "psyche",
						Aliases:  []string{"P"},
						Required: true,
						Usage:    "psyche name",
					},
					&cli.BoolFlag{
						Name:    "with-command",
						Aliases: []string{"W"},
						Usage:   "with decrypt command",
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
