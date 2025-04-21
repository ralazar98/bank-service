package handlers

import (
	"bank-service/internal/entity"
	"bank-service/internal/services"
	"github.com/go-chi/render"
	"net/http"
)

func (a *AccountHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req *entity.UpdateBalance
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	//TODO:Переписать в кейсы
	if operation(req.Operation) == TakeOperation {
		req.ChangingInBalance *= -1

		err := a.bankService.Update(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("Successfully updated"))
		}

	} else if operation(req.Operation) == AddOperation {

		err := a.bankService.Update(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("Successfully updated"))
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(services.InvalidOperationErr.Error()))
	}
}
