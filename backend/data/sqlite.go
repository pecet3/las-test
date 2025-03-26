package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_DIR = "./database/store.db"
)

func NewSQLite() *sql.DB {
	db, err := sql.Open("sqlite3", DB_DIR)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	if _, err := db.Exec(`
	PRAGMA journal_mode=wal;
	PRAGMA foreign_keys = ON;
	`); err != nil {
		log.Fatalf("Failed to setup db")
	}
	return db
}
