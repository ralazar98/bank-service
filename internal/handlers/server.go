package handlers

import (
	"bank-service/configs"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func setupRoutes(bankService BankServiceI) *chi.Mux {
	r := chi.NewRouter()
	r.Use(RequestLogger)

	apiRouter := chi.NewRouter()
	apiRouter.Use(MetricsMiddleware)

	accountHandler := NewAccountHandler(bankService)
	techRouteHandler := NewTechRouteHandler()

	accountHandler.ApiRoute(apiRouter)
	techRouteHandler.TechRoute(r)

	r.Mount("/api", apiRouter)

	return r

}

func NewServer(bankService BankServiceI, cfg configs.AppConfig) *Server {
	handlers := setupRoutes(bankService)

	server := &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: handlers,
		},
	}

	return server
}

func (server *Server) Start() {
	err := server.httpServer.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %v", err)
	}
}

func (server *Server) Restart() {
	//TODO:Подумать о перезапуске серва

}
