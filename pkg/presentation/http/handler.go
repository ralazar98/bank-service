package http

import (
	"bank-service/internal/services"
	"bank-service/pkg/infrastructure/memory_cache"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type bankServiceI interface {
	// some methods
	CreateAccount(userID int, balance float64) error
	ShowBalance(userID int) (float64, error)
}

type AccountHandler struct {
	store       *memory_cache.BankStorage
	bankService bankServiceI
}

func NewAccountHandler(b *memory_cache.BankStorage) *AccountHandler {
	return &AccountHandler{
		store: b,
	}
}

func (a *AccountHandler) Route(r chi.Router) {
	r.Route("/account", func(r chi.Router) {
		r.Post("/create", a.CreateAccount)
		r.Post("/show", a.ShowBalance)
	})

}

func (a *AccountHandler) techRoute(r chi.Router) {
	//todo: tech
}

func (a *AccountHandler) apiRoute(r chi.Router) {
	//todo: api
}

func (a *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req *CreateAccountRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := services.Cache.CreateAccount(a.store, req.UserID, req.Balance)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (a *AccountHandler) ShowBalance(w http.ResponseWriter, r *http.Request) {
	var req ShowBalance
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

}
