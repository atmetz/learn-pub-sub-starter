package pubsub

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SimpleQueueType int

const (
	Durable SimpleQueueType = iota
	Transient
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType, // an enum to represent "durable" or "transient"
) (*amqp.Channel, amqp.Queue, error) {

	newCh, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, fmt.Errorf("could not create channel: %v", err)
	}

	newQueue, err := newCh.QueueDeclare(
		queueName,            // name
		queueType == Durable, // durable
		queueType != Durable, // delete when unused
		queueType != Durable, // exclusive
		false,                // no-wait
		nil,                  // args
	)
	if err != nil {
		return nil, amqp.Queue{}, fmt.Errorf("could not declare queue: %v", err)
	}

	err = newCh.QueueBind(newQueue.Name, key, exchange, false, nil)
	if err != nil {
		return nil, amqp.Queue{}, fmt.Errorf("could not bind queue: %v", err)
	}

	return newCh, newQueue, nil
}
