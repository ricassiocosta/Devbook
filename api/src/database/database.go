package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/lib/pq" // Postgres Driver
)

// Connect open the database connection and returns it
func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DBConnectionString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
