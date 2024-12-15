package services

import (
	"encoding/json"
	"fmt"
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

func Rabbit(service *BankService) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"queue_of_payment", // Имя очереди
		false,              // Не является долговечной (не сохраняется при перезапуске брокера)
		false,              // Не удаляется автоматически, если не используется
		false,              // Не является эксклюзивной (доступна для других подключений)
		false,              // Без ожидания подтверждения от брокера
		nil,                // Дополнительные аргументы отсутствуют
	)
	FailOnError(err, "Failed to declare a queue")

	go func() {
		mess, err := ch.Consume(
			queue.Name, // queue
			"",         // consumer
			false,      // auto-ack
			false,      // exclusive
			false,      // no-local
			false,      // no-wait
			nil,        // args
		)
		if err != nil {
			log.Fatalf("Failed to register a consumer: %v", err)
		}
		fmt.Print(mess)
		log.Printf("Waiting for messages from queue: %s", queue.Name)
		//updater(service, mess)
	}()

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
