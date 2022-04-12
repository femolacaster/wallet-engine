// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: wallet.sql

package postgresql

import (
	"context"
	"database/sql"
	"time"
)

const generateWallet = `-- name: GenerateWallet :one
INSERT INTO wallet (
  wallet_number,
  is_active,
  first_name,
  last_name,
  email,
  secretkey,
  bvn,
  dob,
  currency
)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9
)
RETURNING id
`

type GenerateWalletParams struct {
	WalletNumber string
	IsActive     string
	FirstName    sql.NullString
	LastName     sql.NullString
	Email        string
	Secretkey    string
	Bvn          string
	Dob          sql.NullTime
	Currency     string
}

func (q *Queries) GenerateWallet(ctx context.Context, arg GenerateWalletParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, generateWallet,
		arg.WalletNumber,
		arg.IsActive,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Secretkey,
		arg.Bvn,
		arg.Dob,
		arg.Currency,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertTransaction = `-- name: InsertTransaction :one
INSERT INTO transactions (
  id,
  transaction_ref,
  transaction_type,
  transaction_timestamp,
  amount,
  secretkey,
  transaction_status,
  transaction_description,
  balance,
  wallet_id
)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10
)
RETURNING id
`

type InsertTransactionParams struct {
	ID                     int32
	TransactionRef         string
	TransactionType        string
	TransactionTimestamp   time.Time
	Amount                 string
	Secretkey              string
	TransactionStatus      string
	TransactionDescription string
	Balance                string
	WalletID               sql.NullInt32
}

func (q *Queries) InsertTransaction(ctx context.Context, arg InsertTransactionParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertTransaction,
		arg.ID,
		arg.TransactionRef,
		arg.TransactionType,
		arg.TransactionTimestamp,
		arg.Amount,
		arg.Secretkey,
		arg.TransactionStatus,
		arg.TransactionDescription,
		arg.Balance,
		arg.WalletID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateWallet = `-- name: UpdateWallet :exec
UPDATE wallet SET
  wallet_number = $1,
  is_active = $2,
  first_name = $3,
  last_name = $4,
  email = $5,
  secretkey = $6,
  bvn= $7,
  dob = $8,
  currency = $9,
  modified_time = $10
WHERE id = $11
`

type UpdateWalletParams struct {
	WalletNumber string
	IsActive     string
	FirstName    sql.NullString
	LastName     sql.NullString
	Email        string
	Secretkey    string
	Bvn          string
	Dob          sql.NullTime
	Currency     string
	ModifiedTime sql.NullTime
	ID           int32
}

func (q *Queries) UpdateWallet(ctx context.Context, arg UpdateWalletParams) error {
	_, err := q.db.ExecContext(ctx, updateWallet,
		arg.WalletNumber,
		arg.IsActive,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Secretkey,
		arg.Bvn,
		arg.Dob,
		arg.Currency,
		arg.ModifiedTime,
		arg.ID,
	)
	return err
}
