package config

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewAmqp() (*amqp.Connection, *amqp.Channel) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
		os.Getenv("BROKER_USERNAME"),
		os.Getenv("BROKER_PASSWORD"),
		os.Getenv("BROKER_HOST"),
		os.Getenv("BROKER_PORT"),
		os.Getenv("BROKER_VHOST"),
	)
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return conn, ch
}
