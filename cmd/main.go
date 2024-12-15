package main

import (
	http2 "bank-service/internal/handlers"
	"bank-service/internal/repository/postgresql"
	"bank-service/internal/services"
	"github.com/go-chi/chi/v5"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
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

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	services.FailOnError(err, "Failed to connect to RabbitMQ in check")
	log.Println("RabbitMQ connected")
	defer conn.Close()

	ch, err := conn.Channel()
	services.FailOnError(err, "Failed to open a channel in check")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"queue_of_payment",
		false,
		false,
		false,
		false,
		nil,
	)
	services.FailOnError(err, "Failed to declare a queue")

	//services.CheckUpdateData(ch, queue.Name)

	go service.Updater(ch, queue.Name)

	address := ":" + os.Getenv("PORT")
	http.ListenAndServe(address, r)

}
