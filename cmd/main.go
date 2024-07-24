package main

import (
	"bank-service/internal/services"
	"bank-service/pkg/infrastructure/memory_cache"
	http2 "bank-service/pkg/presentation/http"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	store := memory_cache.New()
	service := services.NewBankService(store)
	r := chi.NewRouter()

	accountHandler := http2.NewAccountHandler(service)

	r.Post("/create", accountHandler.CreateAccount)
	r.Post("/get", accountHandler.ShowBalance)
	r.Get("/list", accountHandler.List)
	r.Post("/update", accountHandler.Update)
	http.ListenAndServe(":8080", r)

}
