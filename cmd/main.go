package main

import (
	"bank-service/configs"
	"bank-service/internal/app"
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
		return
	}

	myApp, err := app.New(cfg)
	if err != nil {
		log.Println(err)
		return
	}
	//TODO: я думаю что имеет смысл исполизовать в структуре app именно поля реализации, будто лучше сделать интерфейс с методом Run и кидать туда все слои для старта
	//struct {
	//	Run []func() error
	//}
	err = myApp.MustRun()
	if err != nil {
		log.Println(err)
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	time.Sleep(1 * time.Second)

	myApp.MustStop()
}
