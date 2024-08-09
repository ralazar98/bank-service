package handlers

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

func NewAccountHandler(bankService *services.BankService) *AccountHandler {
	return &AccountHandler{bankService}
}

func (a *AccountHandler) ApiRoute(r chi.Router) {
	r.Post("/create", a.CreateAccount)
	r.Post("/get", a.ShowBalance)
	r.Post("/update", a.Update)
}

func (a *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req *services.CreateAccount
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	res, err := a.bankService.Create(req)
	if err != nil {
		render.JSON(w, r, err)
	}
	render.JSON(w, r, res)
}

func (a *AccountHandler) ShowBalance(w http.ResponseWriter, r *http.Request) {
	var req *services.GetBalance
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.JSON(w, r, err)
	}
	res, err := a.bankService.Get(req)
	if err != nil {
		render.JSON(w, r, err)
	}
	render.JSON(w, r, res)

}

func (a *AccountHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req *services.UpdateBalance
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.JSON(w, r, err)
	}
	if operation(req.Operation) == TakeOperation {
		req.ChangingInBalance *= -1
		res, err := a.bankService.Update(req)
		if err != nil {
			render.JSON(w, r, err)

		}
		render.JSON(w, r, res)
	} else if operation(req.Operation) == AddOperation {
		res, err := a.bankService.Update(req)
		if err != nil {
			render.JSON(w, r, err)

		}
		render.JSON(w, r, res)
	}
}
