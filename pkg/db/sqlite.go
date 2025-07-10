package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Felipalds/go-stocks/pkg/helpers"
	"github.com/Felipalds/go-stocks/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const INIT_SQL = "../../init.sql"

func InitDB() error {
	// log.Info("Initializing database...")
	db, err := sql.Open("sqlite3", INIT_SQL)
	if err != nil {
		return err
	}

	return db.Ping()
}

func InsertTrade(trade models.Trade) error {
	query := `
		INSERT INTO trades (ticker, price, quantity, date, operation, tax, currency)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query, trade.Ticker, trade.Price, trade.Quantity, trade.Date.Format(helpers.GetLayoutDate()), trade.Operation, trade.Tax, trade.Currency)
	return err
}

func GetAllTrades() ([]models.Trade, error) {
	rows, err := db.Query("SELECT id, ticker, price, quantity, date, operation, tax, currency FROM trades")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trades []models.Trade

	for rows.Next() {
		var t models.Trade
		var dateStr string
		if err := rows.Scan(&t.ID, &t.Ticker, &t.Price, &t.Quantity, &dateStr, &t.Operation, &t.Tax, &t.Currency); err != nil {
			return nil, err
		}
		t.Date, err = time.Parse(helpers.GetLayoutDate(), dateStr)
		if err != nil {
			return nil, err
		}
		trades = append(trades, t)
	}
	return trades, nil
}

func GetStockPrice(ticker string) (*models.StockPrice, error) {
	var stock_price models.StockPrice
	query := `SELECT id, ticker, price, currency FROM stock_prices WHERE ticker = ? LIMIT 1`
	err := db.QueryRow(query, ticker).Scan(&stock_price.ID, &stock_price.Ticker, &stock_price.Price, &stock_price.Currency)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("ticker %s not found", ticker)
	} else if err != nil {
		return nil, fmt.Errorf("error querying ticker %s: %w", ticker, err)
	}

	return &stock_price, nil
}

// https://sqlite.org/lang_upsert.html
func UpdateStockPrices(ticker string, price float64, currency models.Currency) error {

	upsertSQL := `
		INSERT INTO stock_prices (ticker, price, currency)
	  VALUES (?, ?, ?)
	  ON CONFLICT (ticker) DO UPDATE SET
	  price = excluded.price;
	`
	_, err := db.Exec(upsertSQL, ticker, price, currency)
	if err != nil {
		return fmt.Errorf("error upserting ticker %w", err)
	}

	return nil
}
