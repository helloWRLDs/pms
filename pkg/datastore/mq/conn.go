package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewChannel(conf *Config) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(conf.DSN)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}

type QueueRoute struct {
	Name string
}
