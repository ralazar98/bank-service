package handlers

import (
	"bank-service/internal/services"
	"github.com/go-chi/render"
	"net/http"
)

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
