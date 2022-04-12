package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func newDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		return db, fmt.Errorf("there was an error establishing database connection: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return db, fmt.Errorf("ping: %w", err)
	}
	return db, nil
}

func main() {
	db, err := newDB()
	if err != nil {
		log.Fatal("A new connection of the DB could not be placed", err)
	}
	fmt.Println(db)
	//defer db.Close()
	/*
		repo := mysql.NewWallet(db)
		fmt.Println(repo)
	*/

}
