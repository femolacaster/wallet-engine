package service

import (
	"context"
	"fmt"
	"time"

	"github.com/femolacaster/wallet-engine/internal"
)

type Repository interface {
	CreateWallet(ctx context.Context, walletNumber string, isActive string, firstName string, lastName string, email string, secretkey string, bvn string, dob time.Time, currency string) (internal.Wallet, error)
	CreateTransaction(ctx context.Context, transactionRef string, transactionType string, transactionTimestamp time.Time, amount string, secretkey string, transactionStatus string, transactionDescription string, balance string, walletID int32) (internal.Transaction, error)
	Update(ctx context.Context, id int32, isActive string) error
}

type Wallet struct {
	repo Repository
}

type Transaction struct {
	repo Repository
}

// Instantiate new WalletRepo
func NewWallet(repo Repository) *Wallet {
	return &Wallet{
		repo: repo,
	}
}

// Instantiate new TransactionRepo
func NewTransaction(repo Repository) *Transaction {
	return &Transaction{
		repo: repo,
	}
}

func (w *Wallet) CreateWallet(ctx context.Context, walletNumber string, isActive string, firstName string, lastName string, email string, secretkey string, bvn string, dob time.Time, currency string) (internal.Wallet, error) {
	wallet, err := w.repo.CreateWallet(ctx, walletNumber, isActive, firstName, lastName, email, secretkey, bvn, dob, currency)
	if err != nil {
		return internal.Wallet{}, fmt.Errorf("wallet repository creation: %w", err)
	}

	return wallet, nil
}

func (t *Transaction) CreateTransaction(ctx context.Context, transactionRef string, transactionType string, transactionTimestamp time.Time, amount string, secretkey string, transactionStatus string, transactionDescription string, balance string, walletID int32) (internal.Transaction, error) {
	transaction, err := t.repo.CreateTransaction(ctx, transactionRef, transactionType, transactionTimestamp, amount, secretkey, transactionStatus, transactionDescription, balance, walletID)
	if err != nil {
		return internal.Transaction{}, fmt.Errorf("transaction repository creation: %w", err)
	}

	return transaction, nil
}

func (w *Wallet) Update(ctx context.Context, id int32, isActive string) error {
	// XXX: We will revisit the number of received arguments in future episodes.
	if err := w.repo.Update(ctx, id, isActive); err != nil {
		return fmt.Errorf("transaction repository update: %w", err)
	}

	return nil
}
