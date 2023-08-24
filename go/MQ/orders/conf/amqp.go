package conf

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

const (
	EXCHANGE_NAME = "orders"
)

func DialAmqp() (*amqp091.Channel, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_PASS"),
		os.Getenv("RABBIT_HOSTNAME"), os.Getenv("RABBIT_PORT"))
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}

	return conn.Channel()
}
