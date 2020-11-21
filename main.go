package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	cfg := Config{}

	app := &cli.App{
		Name:  "Tweet Random Repo",
		Usage: "Tweet random repositories",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "topic",
				Aliases:     []string{"t"},
				Value:       "golang",
				Usage:       "topic for searching repos",
				Destination: &cfg.Topic,
			},
			&cli.StringFlag{
				Name:        "lang",
				Aliases:     []string{"l"},
				Value:       "",
				Usage:       "language for searching repos",
				Destination: &cfg.Language,
			},
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "./config.json",
				Usage:       "path to config file",
				Destination: &cfg.ConfigFile,
			},
			&cli.Int64Flag{
				Name:        "time",
				Aliases:     []string{"x"},
				Value:       15,
				Usage:       "periodicity of tweet",
				Destination: &cfg.Periodicity,
			},
		},
		Action: func(c *cli.Context) error {
			e := cfg.SetConfigAccess()
			if e != nil {
				panic(e)
			}
			for {
				r := getRepo(cfg)
				tweetRepo(cfg, r)
				time.Sleep(time.Duration(cfg.Periodicity) * time.Minute)
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
