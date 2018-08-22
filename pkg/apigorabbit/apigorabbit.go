package apigorabbit

import (
	"github.com/streadway/amqp"
)

//Producer ...
var Producer producer

type producer struct {
	C        *amqp.Connection
	Chan     *amqp.Channel
	Q        amqp.Queue
	Exchange string
}

func (c producer) Dial(url string) error {
	var err error
	c.C, err = amqp.Dial(url)
	return err
}

func (c producer) Channel() error {
	var err error
	c.Chan, err = c.C.Channel()
	return err
}

func (c producer) QueueDeclare(
	queue string,
	durable bool,
	deleteWhenUnused bool,
	exclusive bool,
	noWait bool,
	args amqp.Table) error {
	var err error
	c.Q, err = c.Chan.QueueDeclare(queue, durable, deleteWhenUnused, exclusive, noWait, args)
	return err
}

func (c producer) Publish(message string) error {
	err := c.Chan.Publish(
		c.Exchange,
		c.Q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte(message),
		})
	return err
}

func (c producer) Close() error {
	return c.C.Close()
}
