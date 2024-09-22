package handlers

import (
	"bank-service/internal/entity"
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

type BankServiceI interface {
	Create(user *services.CreateAccount) (*entity.User, error)
	Get(user *services.GetBalance) (*entity.User, error)
	Update(user *services.UpdateBalance) (*entity.User, error)
}

type AccountHandler struct {
	bankService BankServiceI
}

func NewAccountHandler(bankService BankServiceI) *AccountHandler {
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, res)
	}

}

func (a *AccountHandler) ShowBalance(w http.ResponseWriter, r *http.Request) {
	var req *services.GetBalance
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.JSON(w, r, err)
	}
	res, err := a.bankService.Get(req)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {

		render.JSON(w, r, res)

	}
}

func (a *AccountHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req *services.UpdateBalance
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	if operation(req.Operation) == TakeOperation {
		req.ChangingInBalance *= -1

		res, err := a.bankService.Update(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else {
			render.JSON(w, r, res)
		}

	} else if operation(req.Operation) == AddOperation {

		res, err := a.bankService.Update(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else {
			render.JSON(w, r, res)

		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(services.InvalidOperationErr.Error()))
	}
}
