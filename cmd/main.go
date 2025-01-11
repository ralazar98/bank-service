package main

import (
	http2 "bank-service/internal/handlers"
	"bank-service/internal/rabbit"
	"bank-service/internal/repository/postgresql"
	"bank-service/internal/services"
	"github.com/go-chi/chi/v5"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	r := chi.NewRouter()

	r.Use(http2.RequestLogger)

	apiRouter := chi.NewRouter()
	apiRouter.Use(http2.MetricsMiddleware)

	store := postgresql.New()

	service := services.NewBankService(store)
	accountHandler := http2.NewAccountHandler(service)
	techRouterHandler := http2.NewTechRouteHandler()

	accountHandler.ApiRoute(apiRouter)
	techRouterHandler.TechRoute(r)

	r.Mount("/api", apiRouter)

	newRabbit, err := rabbit.NewRabbit(store)
	if err != nil {
		panic(err)
	}
	go newRabbit.Updater()
	defer newRabbit.Close()

	address := ":" + os.Getenv("PORT")
	http.ListenAndServe(address, r)

}
