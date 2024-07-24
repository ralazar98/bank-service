package main

import (
	"bank-service/internal/services"
	"bank-service/pkg/infrastructure/memory_cache"
	http2 "bank-service/pkg/presentation/http"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	r := chi.NewRouter()
	store := memory_cache.New()
	service := services.NewBankService(store)
	accountHandler := http2.NewAccountHandler(service)
	accountHandler.ApiRoute(r)
	http.ListenAndServe(":8080", r)

}
