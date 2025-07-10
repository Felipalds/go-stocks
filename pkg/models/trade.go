package models

import "time"

type Operation string

const (
	Buy  Operation = "buy"
	Sell Operation = "sell"
)

type Trade struct {
	ID        int
	Ticker    string
	Price     float64
	Quantity  float64
	Date      time.Time
	Operation Operation
	Tax       float64
	Currency  Currency
}
