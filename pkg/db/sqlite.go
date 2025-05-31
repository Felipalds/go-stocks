package db

import (
	"database/sql"
	"os"
	"time"

	"github.com/Felipalds/go-stocks/pkg/helpers"
	"github.com/Felipalds/go-stocks/pkg/models"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(path string) (*sql.DB, error) {
	log.Info("Initializing database...")
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}

func CreateTables(db *sql.DB, path string) error {
	log.Info("Initializing tables...")
	content, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	_, err = db.Exec(string(content))
	return err
}

func InsertTrade(db *sql.DB, trade models.Trade) error {
	query := `
		INSERT INTO trades (ticker, price, quantity, date, operation, tax, currency)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query, trade.Ticker, trade.Price, trade.Quantity, trade.Date.Format(helpers.GetLayoutDate()), trade.Operation, trade.Tax, trade.Currency)
	return err
}

func GetAllTrades(db *sql.DB) ([]models.Trade, error) {
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
