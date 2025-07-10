package main

import (
	"os"

	"github.com/Felipalds/go-stocks/pkg/commands"
	"github.com/Felipalds/go-stocks/pkg/db"
	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

func main() {

	err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:  "Go Stocks",
		Usage: "Track stocks buys and sells",
		Commands: []*cli.Command{
			commands.TradeCommand(),
			commands.ListCommand(),
			commands.ResumeCommand(),
			commands.UpdatePriceCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
