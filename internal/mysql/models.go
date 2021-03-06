// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package mysql

import (
	"database/sql"
	"time"
)

type Transactions struct {
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
	WalletID               sql.NullInt32
}

type Wallets struct {
	ID           int64
	WalletNumber string
	IsActive     string
	FirstName    sql.NullString
	LastName     sql.NullString
	Email        string
	Secretkey    string
	Bvn          string
	Currency     string
	CreatedTime  sql.NullTime
	ModifiedTime sql.NullTime
}
