CREATE TABLE IF NOT EXISTS trades (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  ticker TEXT,
  price REAL,
  quantity REAL,
  date TEXT,
  operation TEXT,
  tax REAL,
  currency TEXT
);

CREATE TABLE IF NOT EXISTS stock_prices (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  ticker TEXT UNIQUE,
  price REAL,
  currency TEXT
)
