package apigorabbit

import (
	"github.com/streadway/amqp"
)

//Producer ...
var Producer producer

type producer struct {
	C *amqp.Connection
}
