package rest_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/mux"

	"github.com/femolacaster/wallet-engine/internal"
	"github.com/femolacaster/wallet-engine/internal/rest"
	"github.com/femolacaster/wallet-engine/internal/rest/resttest"
)

func TestWallet_Post(t *testing.T) {

	t.Parallel()

	type output struct {
		expectedStatus int
		expected       interface{}
		target         interface{}
	}

	tests := []struct {
		name   string
		setup  func(*resttest.FakeWalletService)
		input  []byte
		output output
	}{
		{
			"OK: 201",
			func(s *resttest.FakeWalletService) {
				s.CreateReturns(
					internal.Wallet{
						ID:           3,
						WalletNumber: "cfyuee",
						IsActive:     "3",
						FirstName:    "Jango",
						LastName:     "jumbo",
						Email:        "kin",
						Secretkey:    "kon",
						Bvn:          "xxcdd",
						Currency:     "naira",
					},
					nil)
			},
			func() []byte {
				b, _ := json.Marshal(&rest.CreateWalletRequest{
					FirstName: "Jango",
					LastName:  "jumbo",
					Email:     "kin",
					Secretkey: "kon",
					Bvn:       "xxcdd",
					Currency:  "naira",
				})

				return b
			}(),
			output{
				http.StatusCreated,
				&rest.CreateWalletResponse{
					Wallet: rest.Wallet{
						ID:           3,
						WalletNumber: "cfyuee",
						IsActive:     "3",
						FirstName:    "Jango",
						LastName:     "jumbo",
						Email:        "kin",
						Secretkey:    "kon",
						Bvn:          "xxcdd",
						Currency:     "naira",
					},
				},
				&rest.CreateWalletResponse{},
			},
		},
		{
			"ERR: 400",
			func(*resttest.FakeWalletService) {},
			[]byte(`{"invalid":"json`),
			output{
				http.StatusBadRequest,
				&rest.ErrorResponse{
					Error: "invalid request",
				},
				&rest.ErrorResponse{},
			},
		},
		{
			"ERR: 500",
			func(s *resttest.FakeWalletService) {
				s.CreateReturns(internal.Wallet{},
					errors.New("service error"))
			},
			[]byte(`{}`),
			output{
				http.StatusInternalServerError,
				&rest.ErrorResponse{
					Error: "create failed",
				},
				&rest.ErrorResponse{},
			},
		},
	}

	//-

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			router := mux.NewRouter()
			svc := &resttest.FakeWalletService{}
			tt.setup(svc)

			rest.NewWalletHandler(svc).Register(router)

			//-

			res := doRequest(router,
				httptest.NewRequest(http.MethodPost, "/wallet", bytes.NewReader(tt.input)))

			//-

			assertResponse(t, res, test{tt.output.expected, tt.output.target})

			if tt.output.expectedStatus != res.StatusCode {
				t.Fatalf("expected code %d, actual %d", tt.output.expectedStatus, res.StatusCode)
			}
		})
	}
}

func TestTransaction_Post(t *testing.T) {

	t.Parallel()

	type output struct {
		expectedStatus int
		expected       interface{}
		target         interface{}
	}

	tests := []struct {
		name   string
		setup  func(*resttest.FakeTransactionService)
		input  []byte
		output output
	}{
		{
			"OK: 201",
			func(s *resttest.FakeTransactionService) {
				s.CreateReturns(
					internal.Transaction{
						ID:                     3,
						TransactionRef:         "tyghj",
						TransactionType:        "credit",
						Amount:                 "45000",
						Secretkey:              "xxxxx",
						TransactionStatus:      "open",
						TransactionDescription: "I am paying for school fees",
						WalletID:               3,
					},
					nil)
			},
			func() []byte {
				b, _ := json.Marshal(&rest.CreateTransactionRequest{
					TransactionType:        "credit",
					Amount:                 "45000",
					Secretkey:              "xxxxx",
					TransactionDescription: "I am paying for school fees",
					WalletID:               3,
				})

				return b
			}(),
			output{
				http.StatusCreated,
				&rest.CreateTransactionResponse{
					Transaction: rest.Transaction{
						ID:                     3,
						TransactionRef:         "tyghj",
						TransactionType:        "credit",
						Amount:                 "45000",
						Secretkey:              "xxxxx",
						TransactionStatus:      "open",
						TransactionDescription: "I am paying for school fees",
						WalletID:               3,
					},
				},
				&rest.CreateTransactionResponse{},
			},
		},
		{
			"ERR: 400",
			func(*resttest.FakeTransactionService) {},
			[]byte(`{"invalid":"json`),
			output{
				http.StatusBadRequest,
				&rest.ErrorResponse{
					Error: "invalid request",
				},
				&rest.ErrorResponse{},
			},
		},
		{
			"ERR: 500",
			func(s *resttest.FakeTransactionService) {
				s.CreateReturns(internal.Transaction{},
					errors.New("service error"))
			},
			[]byte(`{}`),
			output{
				http.StatusInternalServerError,
				&rest.ErrorResponse{
					Error: "create failed",
				},
				&rest.ErrorResponse{},
			},
		},
	}

	//-

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			router := mux.NewRouter()
			svc := &resttest.FakeTransactionService{}
			tt.setup(svc)

			rest.NewTransactionHandler(svc).Register(router)

			//-

			res := doRequest(router,
				httptest.NewRequest(http.MethodPost, "/transaction", bytes.NewReader(tt.input)))

			//-

			assertResponse(t, res, test{tt.output.expected, tt.output.target})

			if tt.output.expectedStatus != res.StatusCode {
				t.Fatalf("expected code %d, actual %d", tt.output.expectedStatus, res.StatusCode)
			}
		})
	}
}

func TestWallet_Update(t *testing.T) {
	// XXX: Test "serviceArgs"

	t.Parallel()

	type output struct {
		expectedStatus int
		expected       interface{}
		target         interface{}
	}

	tests := []struct {
		name   string
		setup  func(*resttest.FakeWalletService)
		input  []byte
		output output
	}{
		{
			"OK: 200",
			func(s *resttest.FakeWalletService) {},
			func() []byte {
				b, _ := json.Marshal(&rest.UpdateWalletRequest{
					IsActive: "3",
				})

				return b
			}(),
			output{
				http.StatusOK,
				&struct{}{},
				&struct{}{},
			},
		},
		{
			"ERR: 400",
			func(*resttest.FakeWalletService) {},
			[]byte(`{"invalid":"json`),
			output{
				http.StatusBadRequest,
				&rest.ErrorResponse{
					Error: "invalid request",
				},
				&rest.ErrorResponse{},
			},
		},
		{
			"ERR: 500",
			func(s *resttest.FakeWalletService) {
				s.UpdateReturns(errors.New("service error"))
			},
			[]byte(`{}`),
			output{
				http.StatusInternalServerError,
				&rest.ErrorResponse{
					Error: "The update failed",
				},
				&rest.ErrorResponse{},
			},
		},
	}

	//-

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			router := mux.NewRouter()
			svc := &resttest.FakeWalletService{}
			tt.setup(svc)

			rest.NewWalletHandler(svc).Register(router)

			//-

			res := doRequest(router,
				httptest.NewRequest(http.MethodPut, "/wallets/3", bytes.NewReader(tt.input)))

			//-

			assertResponse(t, res, test{tt.output.expected, tt.output.target})

			if tt.output.expectedStatus != res.StatusCode {
				t.Fatalf("expected code %d, actual %d", tt.output.expectedStatus, res.StatusCode)
			}
		})
	}
}

type test struct {
	expected interface{}
	target   interface{}
}

func doRequest(router *mux.Router, req *http.Request) *http.Response {
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	return rr.Result()
}

func assertResponse(t *testing.T, res *http.Response, test test) {
	t.Helper()

	if err := json.NewDecoder(res.Body).Decode(test.target); err != nil {
		t.Fatalf("couldn't decode %s", err)
	}
	defer res.Body.Close()

	if !cmp.Equal(test.expected, test.target) {
		t.Fatalf("expected results don't match: %s", cmp.Diff(test.expected, test.target))
	}
}
