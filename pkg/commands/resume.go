package commands

import (
	"fmt"

	"github.com/Felipalds/go-stocks/pkg/db"
	"github.com/Felipalds/go-stocks/pkg/helpers"
	"github.com/Felipalds/go-stocks/pkg/models"
	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

// sum of all multiplication of Trade.price by Trade.quantity
func getAllAmountOfMoney(trades []models.Trade) float64 {
	// TODO: add a separation by CURRENCY. It will need to have a table of relations of currencies
	// today it does by BRL only
	total := 0.0
	dollarPrice := helpers.GetDollarInBRL()

	for _, trade := range trades {

		stockPrice, err := db.GetStockPrice(trade.Ticker)
		if err != nil {
			log.Fatal("Error on function getAllAmountOfMoney: %w", err)
		}

		if trade.Currency == models.BRL {
			total += stockPrice.Price * trade.Quantity
		} else if trade.Currency == models.USD {
			total += stockPrice.Price * trade.Quantity * dollarPrice
		} else {
			log.Warn("Invalid currency found")
			continue
		}
	}

	return total

}

func resume() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		trades, err := db.GetAllTrades()
		if err != nil {
			return err
		}

		fmt.Println("Amount of money: R$", getAllAmountOfMoney(trades))
		return nil

	}
}

func ResumeCommand() *cli.Command {
	return &cli.Command{
		Name:  "resume",
		Usage: "Shows a resume of your wallet.",
		Flags: []cli.Flag{
			// currently unused
			&cli.StringFlag{
				Name:    "ticker",
				Aliases: []string{"t"},
				Usage:   "Filters the resume by a ticker",
			},
		},
		Action: resume(),
	}
}
