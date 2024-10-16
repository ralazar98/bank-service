package main

import (
	http2 "bank-service/internal/handlers"
	"bank-service/internal/repository/postgresql"
	"bank-service/internal/services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"os"
)

func main() {
	//Создает роутер
	r := chi.NewRouter()
	r.Mount("/debug/pprof/", http.StripPrefix("/debug/pprof", http.HandlerFunc(pprof.Index)))

	http2.RegisMetrics()
	store := postgresql.New()
	service := services.NewBankService(store)
	accountHandler := http2.NewAccountHandler(service)

	accountHandler.ApiRoute(r)
	address := ":" + os.Getenv("PORT")
	http.ListenAndServe(address, r)

}
