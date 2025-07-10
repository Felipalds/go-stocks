package commands

import (
	"fmt"

	"github.com/Felipalds/go-stocks/pkg/db"
	"github.com/Felipalds/go-stocks/pkg/helpers"
	"github.com/Felipalds/go-stocks/pkg/models"
	"github.com/urfave/cli/v2"
)

func list() func(c *cli.Context) error {
	return func(c *cli.Context) error {

		trades, err := db.GetAllTrades()
		if err != nil {
			return err
		}

		for key, trade := range trades {
			var stock_current_price *models.StockPrice
			stock_current_price, err = db.GetStockPrice(trade.Ticker)

			earns_or_losses := (stock_current_price.Price * trade.Quantity) - (trade.Price * trade.Quantity)
			fmt.Println(earns_or_losses)
			var background string
			if earns_or_losses > 0 {
				background = "\033[30m\u001b[42m"
			} else {
				background = "\033[30m\u001b[41m"
			}

			curr := "R$"
			if trade.Currency == "USD" {
				curr = "$"
			}

			fmt.Printf("(%d) %s   -   %s   -   %.2f   -   %s%.2f %s%s%.2f\u001b[0m\033[m\n", key, trade.Date.Format(helpers.GetLayoutDate()), trade.Ticker, trade.Quantity, curr, trade.Price, background, curr, earns_or_losses)

		}

		return nil
	}
}

func ListCommand() *cli.Command {

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
		Action: list(),
	}
}
