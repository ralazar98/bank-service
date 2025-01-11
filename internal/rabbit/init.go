package rabbit

import (
	"bank-service/internal/entity"
	"bank-service/internal/services"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type rabbit struct {
	con         *amqp.Connection
	channel     *amqp.Channel
	nameOfQueue string
	service     BankI
}

type BankI interface {
	UpdateBalance(user *services.UpdateBalance) (*entity.User, error)
}

func NewRabbit(bankRep BankI) (*rabbit, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	services.FailOnError(err, "Failed to connect to RabbitMQ in check")

	log.Println("RabbitMQ connected")

	ch, err := conn.Channel()
	services.FailOnError(err, "Failed to open a channel in check")

	queue, err := ch.QueueDeclare(
		"queue_of_payment",
		false,
		false,
		false,
		false,
		nil,
	)
	services.FailOnError(err, "Failed to declare a queue")
	return &rabbit{
		con:         conn,
		channel:     ch,
		service:     bankRep,
		nameOfQueue: queue.Name,
	}, nil
}

func (r *rabbit) Updater() {
	msgs, err := r.channel.Consume(
		r.nameOfQueue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if r.channel.IsClosed() {
		log.Println("Канал закрыт после consume")
		return
	}
	FailOnError(err, "Не удалось зарегистрировать потребителя")
	for message := range msgs {
		var user *services.UpdateBalance
		err := json.Unmarshal(message.Body, &user)
		if err != nil {
			FailOnError(err, "Ошибка в дешифровке")
		}
		updateBalance, err := r.service.UpdateBalance(user)
		if err != nil {
			FailOnError(err, "ошибка при изменении баланса")
		}
		log.Println(updateBalance)
	}
}

func (r *rabbit) Close() {
	connectError := r.con.Close()
	channelError := r.channel.Close()
	if connectError != nil {
		FailOnError(channelError, "connectError")
	}
	if channelError != nil {
		FailOnError(channelError, "channelError")
	}

}
