package app

import (
	"bank-service/configs"
	servergrpc "bank-service/internal/gRPC"
	"bank-service/internal/handlers"
	"bank-service/internal/rabbit"
	"bank-service/internal/repository/postgresql"
	"bank-service/internal/services"
)

type App struct {
	storage    *postgresql.BankStorage
	rabbit     *rabbit.Rabbit
	service    *services.BankService
	apiServer  *handlers.Server
	gRPCServer *servergrpc.BankServer
	appGRPC    *servergrpc.AppGRPC
}

func New(cfg *configs.Config) (*App, error) {

	store, err := postgresql.New(cfg.Database)
	if err != nil {
		return nil, err
	}
	myRabbit, err := rabbit.NewRabbit(cfg.RabbitMQ)
	if err != nil {
		return nil, err
	}

	service := services.NewBankService(store, cfg, myRabbit)

	apiServer := handlers.NewServer(service, cfg.App)

	gRPCServer := servergrpc.NewBankServer(*service)

	appGRPC := servergrpc.NewAppGRPC(gRPCServer, cfg.GRPC.GRPCPort)

	return &App{
		storage:    store,
		rabbit:     myRabbit,
		service:    service,
		apiServer:  apiServer,
		gRPCServer: gRPCServer,
		appGRPC:    appGRPC,
	}, nil
}

func (app *App) MustRun() error {

	err := app.apiServer.Start()
	if err != nil {
		return err
	}
	err = app.appGRPC.Run()
	if err != nil {
		return err
	}

	return nil
}

func (app *App) MustStop() {
	app.appGRPC.Stop()
}
