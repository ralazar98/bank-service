package http

import "bank-service/internal/storage"

type AccountHandler struct {
	store storage.BankStorage
}

func NewAccountHandler(b storage.BankStorage) *AccountHandler {
	return &AccountHandler{
		store: b,
	}
}
