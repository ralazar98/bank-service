package main

import (
	"bank-service/pkg/infrastructure/memory_cache"
	http2 "bank-service/pkg/presentation/http"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	store := memory_cache.New()
	r := chi.NewRouter()

	accountHandler := http2.NewAccountHandler(store)
	http.ListenAndServe(":8080", r)
	r.Post("/", accountHandler.CreateAccount)

	//accountHandler.Route(r)

}
