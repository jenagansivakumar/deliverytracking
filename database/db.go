package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath string = "./deliveries.db"

func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)

	}
	fmt.Println("you are now connected")
	return db, nil
}

func InitializeDatabase(conn *sql.DB) error {
	createTableQuery := `
        CREATE TABLE IF NOT EXISTS deliveries (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            address TEXT NOT NULL,
            status TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );
    `
	_, err := conn.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}
