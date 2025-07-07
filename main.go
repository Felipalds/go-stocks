package main

import (
	"os"

	"github.com/Felipalds/go-stocks/pkg/commands"
	"github.com/Felipalds/go-stocks/pkg/db"
	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

func main() {

	database, err := db.InitDB("data.db")
	path := "./init.sql"
	if err != nil {
		log.Fatal(err)
	}
	if err := db.CreateTables(database, path); err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:  "Go Stocks",
		Usage: "Track stocks buys and sells",
		Commands: []*cli.Command{
			commands.TradeCommand(database),
			commands.ListCommand(database),
			commands.ResumeCommand(database),
			commands.UpdatePriceCommand(database),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
