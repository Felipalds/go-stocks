package models

type StockPrice struct {
	ID       int
	Ticker   string
	Price    float64
	Currency Currency
}
