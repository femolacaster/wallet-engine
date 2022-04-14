package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/femolacaster/wallet-engine/internal/mysql"
	"github.com/femolacaster/wallet-engine/internal/rest"
	"github.com/femolacaster/wallet-engine/internal/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func newDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		return db, fmt.Errorf("there was an error establishing database connection: %w", err)
	}
	if err := db.Ping(); err != nil {
		return db, fmt.Errorf("could not ping database: %w", err)
	}
	return db, nil
}

func main() {
	db, err := newDB()
	if err != nil {
		log.Fatal("A new connection of the DB could not be placed", err)
	}

	defer db.Close()
	walletRepo := mysql.NewWallet(db)
	walletSvc := service.NewWallet(walletRepo)
	transactionRepo := mysql.NewTransaction(db)
	transactionSvc := service.NewTransaction(transactionRepo)

	r := mux.NewRouter()

	rest.NewWalletHandler(walletSvc).Register(r)
	rest.NewTransactionHandler(transactionSvc).Register(r)

	address := "0.0.0.0:9967"

	srv := &http.Server{
		Handler:           r,
		Addr:              address,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}
	log.Println("Starting server", address)

	log.Fatal(srv.ListenAndServe())
}
