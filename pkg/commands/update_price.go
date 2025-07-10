package commands

import (
	"strings"

	"github.com/Felipalds/go-stocks/pkg/db"
	"github.com/Felipalds/go-stocks/pkg/models"
	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

func updatePrice() func(c *cli.Context) error {

	return func(c *cli.Context) error {
		ticker := c.String("ticker")
		price := c.Float64("price")
		currency := models.Currency(strings.ToUpper(c.String("currency")))

		err := db.UpdateStockPrices(ticker, price, currency)
		if err != nil {
			log.Fatal("Error updating price: %w", err)
		}

		return nil
	}

}

func UpdatePriceCommand() *cli.Command {

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
			&cli.StringFlag{
				Name:     "currency",
				Aliases:  []string{"c"},
				Usage:    "The currency of the stock/crypto you want to update",
				Required: true,
			},
		},
		Action: updatePrice(),
	}

}
