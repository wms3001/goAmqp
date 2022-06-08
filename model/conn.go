package model

import amqp "github.com/rabbitmq/amqp091-go"

type Conn struct {
	Base
	Connection *amqp.Connection
}

type Channel struct {
	Base
	Channel *amqp.Channel
}

type Queue struct {
	Base
	Queue *amqp.Queue
}
