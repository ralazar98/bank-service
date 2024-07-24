package http

import (
	"bank-service/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type operation string

const (
	AddOperation  operation = "add"
	TakeOperation operation = "take"
)

type AccountHandler struct {
	bankService *services.BankService
}

func NewAccountHandler(service *services.BankService) *AccountHandler {
	return &AccountHandler{
		bankService: service,
	}
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
	err := a.bankService.Create(req.UserID, req.Balance)
	if err != nil {
		return
	}
	render.Status(r, http.StatusCreated)

}

func (a *AccountHandler) ShowBalance(w http.ResponseWriter, r *http.Request) {
	var req ShowBalance
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	balance, err := a.bankService.Get(req.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	render.JSON(w, r, balance)

}

func (a *AccountHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := a.bankService.List()
	if err != nil {
		return
	}
	render.JSON(w, r, list)
	return
}

func (a *AccountHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req UpdateBalanceRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if operation(req.Operation) == TakeOperation {
		err := a.bankService.Update(req.UserID, req.ChangingInBalance*-1)
		if err != nil {
			return
		}
	} else if operation(req.Operation) == AddOperation {
		err := a.bankService.Update(req.UserID, req.ChangingInBalance)
		if err != nil {
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)

}
