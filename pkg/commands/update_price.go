package commands

import (
	"database/sql"

	"github.com/urfave/cli/v2"
)

func updatePrice(database *sql.DB) func(c *cli.Context) error {

	return func(c *cli.Context) error {
		// ticker := c.String("ticker")
		// price := c.Float64("price")

		return nil
	}

}

func UpdatePriceCommand(database *sql.DB) *cli.Command {

	return &cli.Command{
		Name:  "update-price",
		Usage: "Update the price of the stocks/cryptos in the stock market.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "ticker",
				Aliases:  []string{"t"},
				Usage:    "The ticker of the stock/crypto you want to update",
				Required: true,
			},
			&cli.Float64Flag{
				Name:     "price",
				Aliases:  []string{"p"},
				Usage:    "The price of the stock/crypto you want to update",
				Required: true,
			},
		},
		Action: updatePrice(database),
	}

}
