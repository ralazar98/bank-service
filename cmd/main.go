package main

import (
	"bank-service/configs"
	servergrpc "bank-service/internal/gRPC"
	http2 "bank-service/internal/handlers"
	"bank-service/internal/rabbit"
	"bank-service/internal/repository/postgresql"
	"bank-service/internal/services"
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Println(err)
	}

	store := postgresql.New(cfg.Database)

	newRabbit, err := rabbit.NewRabbit(cfg.RabbitMQ)
	if err != nil {
		log.Println("Can't connect to RabbitMQ:", err)
	}

	service := services.NewBankService(store, cfg, newRabbit)

	serv := http2.NewServer(service, cfg.App)
	go serv.Start()

	//TODO:Убрать порт в конфиг
	app := servergrpc.NewApp(*service, 50051)
	//TODO:Обработать ошибку
	go app.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	time.Sleep(1 * time.Second)

	log.Println("Shutting down gRPC...")
	app.Stop()
}
