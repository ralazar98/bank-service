package handlers

import (
	"bank-service/internal/entity"
	"bank-service/internal/services"
	"bank-service/pkg/infrastructure/memory_cache"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccountHandler_test(t *testing.T) {
	bankRep := memory_cache.New()
	service := services.NewBankService(bankRep)
	handler := NewAccountHandler(service)

	testTable := []struct {
		name     string
		body     services.CreateAccount
		wantBody entity.User
		wantErr  error
	}{
		{
			name:     "Create Account Success",
			body:     services.CreateAccount{UserID: 1, Balance: 100},
			wantBody: entity.User{ID: 1, Balance: entity.Balance{Sum: 100}},
		},
		{
			name: "Create Account Failure, already exist",
			body: services.CreateAccount{UserID: 1, Balance: 100},

			wantErr: services.AccountAlreadyExistsErr,
		},
		{
			name: "Create Account Failure, invalid balance",
			body: services.CreateAccount{UserID: 1, Balance: -100},

			wantErr: services.MinusBalanceErr,
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			bodyJSON, errMarsh := json.Marshal(tt.body)
			if errMarsh != nil {
				log.Fatal(errMarsh)
			}
			req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewReader(bodyJSON))
			if err != nil {
				log.Println(err.Error())
			}

			rr := httptest.NewRecorder()
			handler.CreateAccount(rr, req)
			res := rr.Result()
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			if res.StatusCode != http.StatusOK {
				assert.Equal(t, tt.wantErr.Error(), string(body))
			} else {
				var test entity.User
				_ = json.Unmarshal(body, &test)
				assert.Equal(t, tt.wantBody, test)
			}

		})
	}

	testTable1 := []struct {
		name     string
		body     services.GetBalance
		wantBody entity.User
		wantErr  error
	}{
		{
			name:     "Get Balance Success",
			body:     services.GetBalance{UserID: 1},
			wantBody: entity.User{ID: 1, Balance: entity.Balance{Sum: 100}},
		},
		{
			name: "Get Balance Failure, not found",
			body: services.GetBalance{UserID: 2},

			wantErr: services.ChosenAccountNotFoundErr,
		},
		{
			name: "Get Balance Failure, invalid id",
			body: services.GetBalance{UserID: -1},

			wantErr: services.WrongIdErr,
		},
	}
	for _, tt := range testTable1 {
		t.Run(tt.name, func(t *testing.T) {

			bodyJSON, errMarsh := json.Marshal(tt.body)
			if errMarsh != nil {
				log.Fatal(errMarsh)
			}
			req, err := http.NewRequest(http.MethodPost, "/get", bytes.NewReader(bodyJSON))
			if err != nil {
				log.Println(err.Error())
			}

			rr := httptest.NewRecorder()
			handler.ShowBalance(rr, req)
			res := rr.Result()
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			if res.StatusCode != http.StatusOK {
				assert.Equal(t, tt.wantErr.Error(), string(body))
			} else {
				var test entity.User
				_ = json.Unmarshal(body, &test)
				assert.Equal(t, tt.wantBody, test)
			}

		})
	}

	testTable2 := []struct {
		name     string
		body     services.UpdateBalance
		wantBody entity.User
		wantErr  error
	}{
		{
			name:     "Update Balance Success,take",
			body:     services.UpdateBalance{UserID: 1, Operation: "take", ChangingInBalance: 50},
			wantBody: entity.User{ID: 1, Balance: entity.Balance{Sum: 50}},
		},
		{
			name:     "Update Balance Success,add",
			body:     services.UpdateBalance{UserID: 1, Operation: "add", ChangingInBalance: 100},
			wantBody: entity.User{ID: 1, Balance: entity.Balance{Sum: 150}},
		},
		{
			name:    "Update Balance Failure, not found",
			body:    services.UpdateBalance{UserID: 2, Operation: "take", ChangingInBalance: 50},
			wantErr: services.ChosenAccountNotFoundErr,
		},
		{
			name:    "Update Balance Failure, invalid id",
			body:    services.UpdateBalance{UserID: -1, Operation: "take", ChangingInBalance: 50},
			wantErr: services.WrongIdErr,
		},
		{
			name:    "Update Balance Failure, invalid operation",
			body:    services.UpdateBalance{UserID: 1, Operation: "tak", ChangingInBalance: 50},
			wantErr: services.InvalidOperationErr,
		},
		{
			name:    "Update Balance Failure, not enough balance",
			body:    services.UpdateBalance{UserID: 1, Operation: "take", ChangingInBalance: 1500},
			wantErr: services.NotEnoughBalanceErr,
		},
	}
	for _, tt := range testTable2 {
		t.Run(tt.name, func(t *testing.T) {
			bodyJSON, errMarsh := json.Marshal(tt.body)
			if errMarsh != nil {
				log.Fatal(errMarsh)
			}
			req, err := http.NewRequest(http.MethodPost, "/update", bytes.NewReader(bodyJSON))
			if err != nil {
				log.Println(err.Error())
			}

			rr := httptest.NewRecorder()
			handler.Update(rr, req)
			res := rr.Result()
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			if res.StatusCode != http.StatusOK {
				assert.Equal(t, tt.wantErr.Error(), string(body))
			} else {
				var test entity.User
				_ = json.Unmarshal(body, &test)
				assert.Equal(t, tt.wantBody, test)
			}

		})
	}

}
