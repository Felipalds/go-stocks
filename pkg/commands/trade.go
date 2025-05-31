package commands

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Felipalds/go-stocks/pkg/db"
	"github.com/Felipalds/go-stocks/pkg/helpers"
	"github.com/Felipalds/go-stocks/pkg/models"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

func trade(database *sql.DB) func(c *cli.Context) error {
	return func(c *cli.Context) error {

		ticker := c.String("ticker")
		qty := c.Float64("qty")
		price := c.Float64("price")
		currency := c.String("currency")
		tax := c.Float64("tax")

		// guard function to handle errors
		op := ""

		if c.Bool("buy") {
			op = "buy"
		}
		if c.Bool("sell") {
			op = "sell"
		}
		if op == "" {
			return errors.New("No valid operation, expected BUY or SELL")
		}

		// parsing the date
		date := time.Now()
		inputDate := c.String("date")
		if inputDate != "" {
			var err error
			date, err = time.Parse(helpers.GetLayoutDate(), inputDate)
			if err != nil {
				return err
			}
		}

		trade := models.Trade{
			Ticker:    ticker,
			Price:     price,
			Quantity:  qty,
			Date:      date,
			Operation: models.Operation(op),
			Tax:       tax,
			Currency:  models.Currency(currency),
		}

		fmt.Println(trade)
		prompt := fmt.Sprintln("Wish you continue?")

		confirmation, err := helpers.ConfirmPrompt(prompt)
		if err != nil {
			return err
		}

		err = db.InsertTrade(database, trade)
		if err != nil {
			return err
		}

		if confirmation {
			log.Info("Operation inserted successfully!")
		} else {
			log.Warn("Operation couldn't proceed!")
		}

		return nil
	}
}

func TradeCommand(database *sql.DB) *cli.Command {

	return &cli.Command{
		Name:  "trade",
		Usage: "Register a trade operation (buy or sell)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "ticker",
				Aliases:  []string{"t"},
				Usage:    "Name of the stock, e.g.: AAPL, BTC",
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "buy",
				Usage:   "Indicates that the operation is buy",
				Aliases: []string{"b"},
			},
			&cli.BoolFlag{
				Name:    "sell",
				Usage:   "Indicates that the operation is sell",
				Aliases: []string{"s"},
			},
			&cli.Float64Flag{
				Name:     "quantity",
				Aliases:  []string{"q"},
				Usage:    "Quantity of stocks",
				Required: true,
			},
			&cli.Float64Flag{
				Name:     "price",
				Usage:    "Price of ticker",
				Aliases:  []string{"p"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "currency",
				Usage:    "Currency: BRL or USD",
				Aliases:  []string{"c"},
				Required: true,
			},
			&cli.Float64Flag{
				Name:  "tax",
				Usage: "Amount of tax paid in the operation",
			},
			&cli.StringFlag{
				Name:  "date",
				Usage: "Date of the operation in the format YYYY-MM-DD. If ommited, will be used today as date.",
			},
		},
		Action: trade(database),
	}
}
