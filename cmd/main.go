package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/femolacaster/wallet-engine/internal/mysql"
	"github.com/femolacaster/wallet-engine/internal/service"
	_ "github.com/go-sql-driver/mysql"
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
	layout := "2006-01-02"
	dob, err := time.Parse(layout, time.Now().Format(layout))
	if err != nil {
		log.Fatal("DOB time could not be passed", err)
	}
	wallet, err := walletSvc.Create(context.Background(), "ASD123GTHOS", "0", "Olufemi", "Alabi", "femdomsteve@yahoo.com", "SCOFGRUIDS", "123453789", dob, "naira")
	if err != nil {
		log.Fatal("Could not insert wallet record", err)
	}
	fmt.Printf("NEW wallet %#v, err %s\n", wallet, err)

	if err := walletSvc.Update(context.Background(), wallet.ID, "2"); err != nil {
		log.Fatal("Could not update database", err)
	}
	transaction, err := transactionSvc.Create(context.Background(), "XDCVIONDFRTD", "debit", time.Now(), "600.90", "xxxxxxx", "success", "Payment for girls", "1000", 1)
	if err != nil {
		log.Fatal("Could not insert transaction record", err)
	}
	fmt.Printf("NEW transaction %#v, err %s\n", transaction, err)

}
