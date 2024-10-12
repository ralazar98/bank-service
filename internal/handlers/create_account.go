package handlers

import (
	"bank-service/internal/services"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

func (a *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	RequestsCounter.Inc()
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
	duration := time.Since(start).Seconds()
	ResponseDuration.Observe(duration)
}
