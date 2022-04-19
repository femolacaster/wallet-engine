package service

import (
	"context"
	"fmt"

	"github.com/femolacaster/wallet-engine/internal"
)

type WalletRepository interface {
	Create(ctx context.Context, walletNumber string, isActive string, firstName string, lastName string, email string, secretkey string, bvn string, currency string) (internal.Wallet, error)
	Update(ctx context.Context, id int64, isActive string) error
}

type TransactionRepository interface {
	Create(ctx context.Context, transactionRef string, transactionType string, amount string, secretkey string, transactionStatus string, transactionDescription string, walletID int32) (internal.Transaction, error)
}

type Wallet struct {
	repo WalletRepository
}

type Transaction struct {
	repo TransactionRepository
}

// Instantiate new WalletRepo
func NewWallet(repo WalletRepository) *Wallet {
	return &Wallet{
		repo: repo,
	}
}

// Instantiate new TransactionRepo
func NewTransaction(repo TransactionRepository) *Transaction {
	return &Transaction{
		repo: repo,
	}
}

func (w *Wallet) Create(ctx context.Context, walletNumber string, isActive string, firstName string, lastName string, email string, secretkey string, bvn string, currency string) (internal.Wallet, error) {
	wallet, err := w.repo.Create(ctx, walletNumber, isActive, firstName, lastName, email, secretkey, bvn, currency)
	if err != nil {
		return internal.Wallet{}, fmt.Errorf("wallet repository creation: %w", err)
	}

	return wallet, nil
}

func (t *Transaction) Create(ctx context.Context, transactionRef string, transactionType string, amount string, secretkey string, transactionStatus string, transactionDescription string, walletID int32) (internal.Transaction, error) {
	transaction, err := t.repo.Create(ctx, transactionRef, transactionType, amount, secretkey, transactionStatus, transactionDescription, walletID)
	if err != nil {
		return internal.Transaction{}, fmt.Errorf("transaction repository creation: %w", err)
	}

	return transaction, nil
}

func (w *Wallet) Update(ctx context.Context, id int64, isActive string) error {
	if err := w.repo.Update(ctx, id, isActive); err != nil {
		return fmt.Errorf("transaction repository update: %w", err)
	}

	return nil
}
