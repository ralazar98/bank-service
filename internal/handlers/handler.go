package handlers

import (
	"bank-service/internal/entity"
	"bank-service/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	r.Get("/health", HealthCheckHandler)
	r.Handle("/metrics", promhttp.Handler())

}
