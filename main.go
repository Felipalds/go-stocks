package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

var confirmStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("%FFD700")).Bold(true)

func ConfirmPrompt(message string) (bool, error) {
	fmt.Print(confirmStyle.Render(message + " [y/N]:"))

	var response string

	fmt.Scanln(&response)

	if response == "n" || response == "N" || response == "" {
		return false, nil
	} else {
		if response == "y" || response == "Y" {
			return true, nil
		} else {
			return false, errors.New("Unexpected answer")
		}
	}
}

func trade(c *cli.Context) error {
	ticker := c.String("ticker")
	qty := c.Float64("qty")
	price := c.Float64("price")
	currency := c.String("currency")

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

	prompt := fmt.Sprintf("Do you with to %s %f %s for  %f %s each?", op, qty, ticker, price, currency)

	confirmation, err := ConfirmPrompt(prompt)
	if err == nil {
		return err
	}

	if confirmation {
		log.Info("Done")
	} else {
		log.Warn("Not done!")
	}

	return nil
}

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "trade",
				Aliases: []string{"t"},
				Usage:   "Register a trade operation (buy or sell)",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "ticker",
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
						Name:     "qty",
						Usage:    "Quantity of stocks",
						Required: true,
					},
					&cli.Float64Flag{
						Name:     "price",
						Usage:    "Price of ticker",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "currency",
						Usage:    "Currency: BRL or USD",
						Required: true,
					},
				},
				Action: trade,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
