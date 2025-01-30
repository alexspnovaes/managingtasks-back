package database

import (
	"database/sql"
	"log"
	"managingtasks/config"

	_ "modernc.org/sqlite"
)

func Connect() (*sql.DB, error) {
	config.LoadEnv()
	dbPath := config.GetEnv("DATABASE_URL", "tasks.db")
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Println("Error creating table:", err)
	}
	return err
}
