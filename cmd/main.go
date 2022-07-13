package main

import (
	"os"

	"github.com/mtk3d/food-api/cmd/flags"
	"github.com/mtk3d/food-api/cmd/server"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     flags.CorsAllowOrigin,
				EnvVars:  []string{flags.CorsAllowOrigin},
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "server",
				Aliases:     []string{"s"},
				Description: "Run application to handle endpoints",
				Action:      server.Command,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
