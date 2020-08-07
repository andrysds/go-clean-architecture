package connection

import "database/sql"

// NewDB creates a new SQLite3 connection
func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./go-clean-architecture.db")
	if err != nil {
		panic(err)
	}
	return db
}
