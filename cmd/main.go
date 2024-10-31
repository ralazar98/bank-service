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
	//configs.GetEnv()
	//Создает роутер
	r := chi.NewRouter()

	r.Use(http2.RequestLogger)
	r.Use(http2.MetricsMiddleware)
	r.Mount("/debug/pprof/", http.StripPrefix("/debug/pprof", http.HandlerFunc(pprof.Index)))

	store := postgresql.New()
	service := services.NewBankService(store)
	accountHandler := http2.NewAccountHandler(service)

	accountHandler.ApiRoute(r)
	accountHandler.TechRoute(r)
	address := ":" + os.Getenv("PORT")
	http.ListenAndServe(address, r)

}
