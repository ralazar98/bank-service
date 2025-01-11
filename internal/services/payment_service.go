package services

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func sendToPaymentService(user *UpdateBalance) {

	body, err := json.Marshal(user)
	if err != nil {
		log.Fatal("marshal err: ", err)
	}

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	log.Println("RabbitMQ connected")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"queue_of_payment",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	FailOnError(err, "Failed to publish a message")

}

func (s *BankService) Updater(channel *amqp.Channel, nameOfQueue string) {
	msgs, err := channel.Consume(
		nameOfQueue, // Имя очереди
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if channel.IsClosed() {
		log.Println("Канал закрыт после")
		return
	}
	FailOnError(err, "Не удалось зарегистрировать потребителя")
	log.Println("UPDATER")
	for message := range msgs {
		var user *UpdateBalance
		err := json.Unmarshal(message.Body, &user)
		if err != nil {
			log.Fatal("unmarshal err: ", err)
		}
		updatedBalance, err := s.BankRep.UpdateBalance(user)
		if err != nil {
			log.Fatal("update err: ", err)
		}
		log.Printf("Updated balance: %+v", updatedBalance)
	}
}

func CheckUpdateData(channel *amqp.Channel, nameOfQueue string) {

	go func() {
		/*		if conn.IsClosed() {
				log.Println("Соединение закрыто")
			}*/
		if channel.IsClosed() {
			log.Println("Канал закрыт")
			return
		} else {
			log.Println("Канал открыт")
		}
		msgs, err := channel.Consume(
			nameOfQueue, // Имя очереди
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if channel.IsClosed() {
			log.Println("Канал закрыт после")
			return
		}
		FailOnError(err, "Не удалось зарегистрировать потребителя")

		for messages := range msgs {
			var user *UpdateBalance
			err := json.Unmarshal(messages.Body, &user)
			if err != nil {
				log.Fatal("unmarshal err: ", err)
			}
			log.Println("TEST")
			log.Printf("Updated balance: %+v", user.ChangingInBalance)
		}
	}()
}
