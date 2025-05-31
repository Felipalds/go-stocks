package commands

import (
	"database/sql"
	"fmt"

	"github.com/Felipalds/go-stocks/pkg/db"
	"github.com/Felipalds/go-stocks/pkg/helpers"
	"github.com/urfave/cli/v2"
)

func list(database *sql.DB) func(c *cli.Context) error {
	return func(c *cli.Context) error {

		trades, err := db.GetAllTrades(database)
		if err != nil {
			return err
		}

		for key, trade := range trades {
			curr := "R$"
			if trade.Currency == "USD" {
				curr = "$"
			}

			fmt.Printf("(%d) %s   -   %s   -   %f   -   %s%f\n", key, trade.Date.Format(helpers.GetLayoutDate()), trade.Ticker, trade.Quantity, curr, trade.Price)

		}

		return nil
	}
}

func ListCommand(database *sql.DB) *cli.Command {

	return &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "Lists all operations sorted",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "ticker",
				Usage: "Name of the stock that you want to sort operations.",
			},
			&cli.BoolFlag{
				Name:    "buy",
				Usage:   "Sort BUY operations",
				Aliases: []string{"b"},
			},
			&cli.BoolFlag{
				Name:    "sell",
				Usage:   "Sort SELL operations",
				Aliases: []string{"s"},
			},
			&cli.StringFlag{
				Name:  "currency",
				Usage: "Sort by currency",
			},
			&cli.StringFlag{
				Name:  "date",
				Usage: "Date of the operation in the format YYYY-MM-DD. If ommited, will sort all by newest to oldest",
			},
		},
		Action: list(database),
	}
}
