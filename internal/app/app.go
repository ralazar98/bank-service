package app

import (
	servergrpc "bank-service/internal/gRPC"
	"bank-service/internal/handlers"
	"bank-service/internal/rabbit"
	"bank-service/internal/repository/postgresql"
	"bank-service/internal/services"
	"log"
)

type AppRunner interface {
	Run() error
}

type App struct {
	storage    *postgresql.BankStorage
	rabbit     *rabbit.Rabbit
	service    *services.BankService
	apiServer  *handlers.Server
	gRPCServer *servergrpc.BankServer
	appGRPC    *servergrpc.AppGRPC
	options    []AppRunner
}

func New(options ...AppRunner) (*App, error) {

	return &App{
		options: options,
	}, nil
}

func (app *App) Runner() {
	for _, option := range app.options {
		go func() {
			err := option.Run()
			if err != nil {
				log.Println("Runner Error:", err)
			}
		}()
	}

}

func (app *App) MustRun() error {

	err := app.apiServer.Run()
	if err != nil {
		return err
	}
	log.Println(app.appGRPC)
	err = app.appGRPC.Run()
	if err != nil {
		return err
	}

	return nil
}

func (app *App) MustStop() {
	app.appGRPC.Stop()
}
