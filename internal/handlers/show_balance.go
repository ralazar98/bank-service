package handlers

import (
	"bank-service/internal/services"
	"github.com/go-chi/render"
	"net/http"
)

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
