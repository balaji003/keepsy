package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Conn *sql.DB
}

func Connect(connString string) (*DB, error) {
	// Open checks if the arguments are valid, but doesn't connect
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	// Verify connection
	// Simple retry logic or just ping
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	// Set reasonable pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Connected to MySQL database successfully")
	return &DB{Conn: db}, nil
}

func (d *DB) Close() {
	if d.Conn != nil {
		d.Conn.Close()
	}
}
