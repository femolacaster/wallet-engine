package mysql

import (
	"context"
	"database/sql"
	"fmt"

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
func (w *Wallet) Create(ctx context.Context, walletNumber string, isActive string, firstName string, lastName string, email string, secretkey string, bvn string, currency string) (internal.Wallet, error) {
	sqlResult, err := w.q.GenerateWallet(ctx, GenerateWalletParams{
		WalletNumber: walletNumber,
		IsActive:     isActive,
		FirstName:    newNullString(firstName),
		LastName:     newNullString(lastName),
		Email:        email,
		Secretkey:    secretkey,
		Bvn:          bvn,
		Currency:     currency,
	})
	if err != nil {
		return internal.Wallet{}, fmt.Errorf("generating a new wallet failed: %w", err)
	}
	id, err := sqlResult.LastInsertId()
	fmt.Println("The last insert is", id)
	if err != nil {
		return internal.Wallet{}, fmt.Errorf("there are no last Wallet ids on insert: %w", err)
	}

	return internal.Wallet{
		ID:           id,
		WalletNumber: walletNumber,
		IsActive:     isActive,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		Secretkey:    secretkey,
		Bvn:          bvn,
		Currency:     currency,
	}, nil
}

// Generate Transactions
func (t *Transaction) Create(ctx context.Context, transactionRef string, transactionType string, amount string, secretkey string, transactionStatus string, transactionDescription string, walletID int32) (internal.Transaction, error) {
	sqlResult, err := t.q.InsertTransaction(ctx, InsertTransactionParams{
		TransactionRef:         transactionRef,
		TransactionType:        transactionType,
		Amount:                 amount,
		Secretkey:              secretkey,
		TransactionStatus:      transactionStatus,
		TransactionDescription: transactionDescription,
		WalletID:               newNullInt32(walletID),
	})
	if err != nil {
		return internal.Transaction{}, fmt.Errorf("Transaction failed: %w", err)
	}
	id, err := sqlResult.LastInsertId()
	fmt.Println("The last insert is", id)
	if err != nil {
		return internal.Transaction{}, fmt.Errorf("there are no last transaction ids on insert: %w", err)
	}

	return internal.Transaction{
		ID:                     id,
		TransactionRef:         transactionRef,
		TransactionType:        transactionType,
		Amount:                 amount,
		Secretkey:              secretkey,
		TransactionStatus:      transactionStatus,
		TransactionDescription: transactionDescription,
		WalletID:               walletID,
	}, nil
}

// Deactivate or Activate a Wallet
func (w *Wallet) Update(ctx context.Context, id int64, isActive string) error {
	if err := w.q.ChangeWalletStatus(ctx, ChangeWalletStatusParams{
		ID:       id,
		IsActive: isActive,
	}); err != nil {
		return fmt.Errorf("could not update wallet: %w", err)
	}

	return nil
}
