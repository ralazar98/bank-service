package server

import (
	"bank-service/configs"
	"bank-service/internal/app"
	servergrpc "bank-service/internal/gRPC"
	"bank-service/internal/handlers"
	"bank-service/internal/rabbit"
	"bank-service/internal/repository/postgresql"
	"bank-service/internal/services"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() error {
	cfg, err := configs.LoadConfig()
	if err != nil {
		return err
	}

	store, err := postgresql.New(cfg.Database)
	if err != nil {

	}
	myRabbit, err := rabbit.NewRabbit(cfg.RabbitMQ)
	if err != nil {

	}

	service := services.NewBankService(store, cfg, myRabbit)

	apiServer := handlers.NewServer(service, cfg.App)

	gRPCServer := servergrpc.NewBankServer(*service)

	appGRPC := servergrpc.NewAppGRPC(gRPCServer, cfg.GRPC.GRPCPort)

	myApp, err := app.New(apiServer, appGRPC)
	if err != nil {
		return err
	}

	myApp.Runner()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	time.Sleep(1 * time.Second)

	myApp.MustStop()

	return nil
}
