package internal

import (
	"database/sql"
	"time"
)

type Transaction struct {
	ID                     int64
	TransactionRef         string
	TransactionType        string
	TransactionTimestamp   time.Time
	Amount                 string
	Secretkey              string
	TransactionStatus      string
	TransactionDescription string
	CreatedTime            sql.NullTime
	ModifiedTime           sql.NullTime
	WalletID               int32
}

// TODO:Validate each of the transaction types
/*
func (t Transaction) Validate() error {
	if t.TransactionRef == "" {
		return errors.New("transaction reference is required")
	}

	return nil
}
*/

type Wallet struct {
	ID           int64
	WalletNumber string
	IsActive     string
	FirstName    string
	LastName     string
	Email        string
	Secretkey    string
	Bvn          string
	Dob          time.Time
	Currency     string
	CreatedTime  sql.NullTime
	ModifiedTime sql.NullTime
}

// TODO:Validate each of the wallet types
/*
func (w Wallet) Validate() error {
	if w.WalletNumber == "" {
		return errors.New("Wallet Number is required")
	}

	return nil
}
*/
