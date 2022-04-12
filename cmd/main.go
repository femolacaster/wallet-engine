package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func newDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return db, fmt.Errorf("There was an error establishing database connection: %w", err)
	}
	if err := db.Ping(); err != nil {
		return db, fmt.Errorf("The database server could not be pinged: %w", err)
	}
	return db, nil
}

func main() {
	db, err := newDB()
	if err != nil {
		log.Fatal("A new connection of the DB could not be placed", err)
	}
	defer db.Close()

	repo := postgresql.NewTask(db)
}
