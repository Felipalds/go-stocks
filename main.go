package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

func buy(c *cli.Context) error {
	stock := c.Args().Get(0)
	log.Infof("Buying! %s\n", stock)
	return nil
}

func sell(c *cli.Context) error {
	stock := c.Args().Get(0)
	log.Infof("Selling %s\n", stock)
	return nil
}

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "buy",
				Aliases: []string{"b"},
				Action:  buy,
			},
			{
				Name:    "sell",
				Aliases: []string{"s"},
				Action:  sell,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
