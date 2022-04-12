package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/femolacaster/wallet-engine/internal"
)

type Transaction struct {
	q *Queries
}

type Wallet struct {
	q *Queries
}

// Instantiate new Wallet
func NewWallet(db *sql.DB) *Wallet {
	return &Wallet{
		q: New(db),
	}
}

// Instantiate new Transaction
func NewTransaction(db *sql.DB) *Transaction {
	return &Transaction{
		q: New(db),
	}
}

// Generate Wallet
func (w *Wallet) Create(ctx context.Context, walletNumber string, isActive string, firstName string, lastName string, email string, secretkey string, bvn string, dob time.Time, currency string) (internal.Wallet, error) {

	sqlResult, err := w.q.GenerateWallet(ctx, GenerateWalletParams{
		IsActive:  isActive,
		FirstName: newNullString(firstName),
		LastName:  newNullString(lastName),
		Email:     email,
		Secretkey: secretkey,
		Bvn:       bvn,
		Dob:       newNullTime(dob),
		Currency:  currency,
	})
	if err != nil {
		return internal.Wallet{}, fmt.Errorf("generating a new wallet failed: %w", err)
	}
	id, err := sqlResult.LastInsertId()

	return internal.Wallet{
		ID:        id,
		IsActive:  isActive,
		FirstName: newNullString(firstName),
		LastName:  newNullString(lastName),
		Email:     email,
		Secretkey: secretkey,
		Bvn:       bvn,
		Dob:       newNullTime(dob),
		Currency:  currency,
	}, nil
}

// Generate Transactions
func (t *Transaction) Create(ctx context.Context, transactionRef string, transactionType string, transactionTimestamp time.Time, amount string, secretkey string, transactionStatus string, transactionDescription string, balance string, walletID int32) (internal.Transaction, error) {
	sqlResult, err := t.q.InsertTransaction(ctx, InsertTransactionParams{
		TransactionRef:         transactionRef,
		TransactionType:        transactionType,
		TransactionTimestamp:   transactionTimestamp,
		Amount:                 amount,
		Secretkey:              secretkey,
		TransactionStatus:      transactionStatus,
		TransactionDescription: transactionDescription,
		Balance:                balance,
		WalletID:               newNullInt32(walletID),
	})
	if err != nil {
		return internal.Transaction{}, fmt.Errorf("Transaction failed: %w", err)
	}
	id, err := sqlResult.LastInsertId()

	return internal.Transaction{
		ID:                     id,
		TransactionRef:         transactionRef,
		TransactionType:        transactionType,
		TransactionTimestamp:   transactionTimestamp,
		Amount:                 amount,
		Secretkey:              secretkey,
		TransactionStatus:      transactionStatus,
		TransactionDescription: transactionDescription,
		Balance:                balance,
		WalletID:               newNullInt32(walletID),
	}, nil
}

// Deactivate or Activate a Wallet
func (w *Wallet) Update(ctx context.Context, isActive string) error {
	if err := w.q.ChangeWalletStatus(ctx, ChangeWalletStatusParams{
		IsActive: isActive,
	}); err != nil {
		return fmt.Errorf("could not update wallet: %w", err)
	}

	return nil
}
