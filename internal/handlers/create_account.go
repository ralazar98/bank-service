package handlers

import (
	"bank-service/internal/entity"
	"github.com/go-chi/render"
	"net/http"
)

func (a *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {

	var req *entity.CreateAccount

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
