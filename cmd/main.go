package main

import (
	http2 "bank-service/internal/handlers"
	"bank-service/internal/repository/postgresql"
	"bank-service/internal/services"
	"github.com/go-chi/chi/v5"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	//Создает роутер
	r := chi.NewRouter()
	http2.RegisMetrics()
	store := postgresql.New()
	service := services.NewBankService(store)
	accountHandler := http2.NewAccountHandler(service)

	accountHandler.ApiRoute(r)
	address := ":" + os.Getenv("PORT")
	http.ListenAndServe(address, r)

}
