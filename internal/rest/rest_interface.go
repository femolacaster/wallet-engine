package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/xid"

	"github.com/femolacaster/wallet-engine/internal"
)

const numberRegEx string = `[0-9]+`

type WalletService interface {
	Create(ctx context.Context, walletNumber string, isActive string, firstName string, lastName string, email string, secretkey string, bvn string, dob time.Time, currency string) (internal.Wallet, error)
	Update(ctx context.Context, id int64, isActive string) error
}

type TransactionService interface {
	Create(ctx context.Context, transactionRef string, transactionType string, transactionTimestamp time.Time, amount string, secretkey string, transactionStatus string, transactionDescription string, walletID int32) (internal.Transaction, error)
}

type WalletHandler struct {
	svc WalletService
}

type TransactionHandler struct {
	svc TransactionService
}

func NewWalletHandler(svc WalletService) *WalletHandler {
	return &WalletHandler{
		svc: svc,
	}
}

func NewTransactionHandler(svc TransactionService) *TransactionHandler {
	return &TransactionHandler{
		svc: svc,
	}
}

func (w *WalletHandler) Register(r *mux.Router) {
	r.HandleFunc("/wallet", w.create).Methods(http.MethodPost)
	r.HandleFunc(fmt.Sprintf("/wallets/{id:%s}", numberRegEx), w.update).Methods(http.MethodPut)
}

func (t *TransactionHandler) Register(r *mux.Router) {
	r.HandleFunc("/transaction", t.create).Methods(http.MethodPost)
}

type Wallet struct {
	ID           int64     `json:"id"`
	WalletNumber string    `json:"wallet_number"`
	IsActive     string    `json:"is_active"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Secretkey    string    `json:"secretkey"`
	Bvn          string    `json:"bvn"`
	Dob          time.Time `json:"dob"`
	Currency     string    `json:"currency"`
}

type Transaction struct {
	ID                     int64     `json:"id"`
	TransactionRef         string    `json:"transaction_ref"`
	TransactionType        string    `json:"transaction_type"`
	TransactionTimestamp   time.Time `json:"transaction_timestamp"`
	Amount                 string    `json:"amount"`
	Secretkey              string    `json:"secretkey"`
	TransactionStatus      string    `json:"transaction_status"`
	TransactionDescription string    `json:"transaction_description"`
	WalletID               int32     `json:"wallet_id"`
}

type CreateWalletRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Secretkey string    `json:"secretkey"`
	Bvn       string    `json:"bvn"`
	Dob       time.Time `json:"dob"`
	Currency  string    `json:"currency"`
}

type CreateTransactionRequest struct {
	TransactionType        string `json:"transaction_type"`
	Amount                 string `json:"amount"`
	Secretkey              string `json:"secretkey"`
	TransactionDescription string `json:"transaction_description"`
	WalletID               int32  `json:"wallet_id"`
}

type CreateWalletResponse struct {
	Wallet Wallet `json:"wallet"`
}

type CreateTransactionResponse struct {
	Transaction Transaction `json:"transaction"`
}

func (wa *WalletHandler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateWalletRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		renderErrorResponse(w, "invalid request", http.StatusBadRequest)
		return
	}
	fmt.Printf("The last request is%+v\n", req)
	defer r.Body.Close()

	walletNumber := xid.New().String()

	wallet, err := wa.svc.Create(r.Context(), walletNumber, "1", req.FirstName, req.LastName, req.Email, req.Secretkey, req.Bvn, req.Dob, req.Currency)
	if err != nil {
		renderErrorResponse(w, "create failed", http.StatusInternalServerError)
		return
	}

	renderResponse(w,
		&CreateWalletResponse{
			Wallet: Wallet{
				ID:           wallet.ID,
				WalletNumber: wallet.WalletNumber,
				IsActive:     wallet.IsActive,
				FirstName:    wallet.FirstName,
				LastName:     wallet.LastName,
				Email:        wallet.Email,
				Secretkey:    wallet.Secretkey,
				Bvn:          wallet.Bvn,
				Dob:          wallet.Dob,
				Currency:     wallet.Currency,
			},
		},
		http.StatusCreated)
}

type UpdateWalletRequest struct {
	IsActive string `json:"is_active"`
}

func (wa *WalletHandler) update(w http.ResponseWriter, r *http.Request) {
	var req UpdateWalletRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		renderErrorResponse(w, "invalid request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	id := mux.Vars(r)["id"]

	parsedID, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		fmt.Println("There was an error Parsing ID")
	}

	err = wa.svc.Update(r.Context(), parsedID, req.IsActive)
	if err != nil {
		renderErrorResponse(w, "The update failed", http.StatusInternalServerError)
		return
	}

	renderResponse(w, &struct{}{}, http.StatusOK)
}

func (t *TransactionHandler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		renderErrorResponse(w, "invalid request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	transactionRef := xid.New().String()
	transactionTimestamp := time.Now()

	transaction, err := t.svc.Create(r.Context(), transactionRef, req.TransactionType, transactionTimestamp, req.Amount, req.Secretkey, "success", req.TransactionDescription, req.WalletID)

	if err != nil {
		renderErrorResponse(w, "create failed", http.StatusInternalServerError)
		return
	}

	renderResponse(w,
		&CreateTransactionResponse{
			Transaction: Transaction{
				ID:                     transaction.ID,
				TransactionRef:         transaction.TransactionRef,
				TransactionType:        transaction.TransactionType,
				TransactionTimestamp:   transactionTimestamp,
				Amount:                 transaction.Amount,
				Secretkey:              transaction.Secretkey,
				TransactionStatus:      transaction.TransactionStatus,
				TransactionDescription: transaction.TransactionDescription,
				WalletID:               transaction.WalletID,
			},
		},
		http.StatusCreated)
}
