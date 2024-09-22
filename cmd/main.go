package main

import (
	"bank-service/internal/config"
	"bank-service/internal/services"
	"bank-service/pkg/infrastructure/memory_cache/postgresql"
	http2 "bank-service/pkg/presentation/http/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	cfg := config.InitConfig()

	conConfig := postgresql.NewConnConfig(cfg)
	conn, _ := postgresql.NewConnect(conConfig)

	store := postgresql.New(conn)

	service := services.NewBankService(store)
	accountHandler := http2.NewAccountHandler(service)

	r := chi.NewRouter()
	accountHandler.ApiRoute(r)

	http.ListenAndServe("localhost:8080", r)
}
