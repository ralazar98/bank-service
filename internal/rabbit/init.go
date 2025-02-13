package rabbit

import (
	"bank-service/configs"
	"bank-service/internal/entity"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Rabbit struct {
	con     *amqp.Connection
	channel *amqp.Channel
	cfg     configs.RabbitMQConfig
}

func NewRabbit(cfg configs.RabbitMQConfig) (*Rabbit, error) {
	newRabbit := &Rabbit{
		cfg: cfg,
	}

	err := newRabbit.NewConnection()
	if err != nil {
		return nil, err
	}

	return newRabbit, nil
}

func (myRabbit *Rabbit) NewConnection() error {
	conn, err := amqp.Dial(myRabbit.cfg.RabbitURL)
	if err != nil {
		return err
	}
	log.Println("Connected to RabbitMQ")

	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	log.Println("Created channel")

	_, err = channel.QueueDeclare(
		myRabbit.cfg.NameOfQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	log.Println("Queue declare")

	myRabbit.con = conn
	myRabbit.channel = channel
	return nil
}

func (myRabbit *Rabbit) SendToPaymentService(user *entity.UpdateBalance) error {

	body, err := json.Marshal(user)
	if err != nil {
		log.Fatal("marshal err: ", err)
	}

	err = myRabbit.channel.Publish(
		"",
		myRabbit.cfg.NameOfQueue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (myRabbit *Rabbit) Close() error {
	connectError := myRabbit.con.Close()
	if connectError != nil {
		return connectError
	}
	channelError := myRabbit.channel.Close()
	if channelError != nil {
		return channelError
	}
	return nil
}
