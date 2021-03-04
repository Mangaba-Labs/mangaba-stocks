package database

import (
	"database/sql"
	// Who knows
	_ "github.com/mattn/go-sqlite3"
)

// Instance database
var Instance *sql.DB

// SetupDatabase setting up database
func SetupDatabase() {
	Instance, _ = sql.Open("sqlite3", "./shares.db")
	statement, _ := Instance.Prepare("CREATE TABLE IF NOT EXISTS shares (id INTEGER PRIMARY KEY, company_name TEXT, name TEXT, buy_value FLOAT, now_value FLOAT)")
	statement.Exec()
}
