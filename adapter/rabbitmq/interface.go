package rabbitmq

import (
	"github.com/chien-dd/library/clock"
	"time"

	"github.com/streadway/amqp"
)

const (
	Transient      = amqp.Transient
	Persistent     = amqp.Persistent
	delayReconnect = 3 * clock.Second
	delayPublish   = clock.Second
	delayConsume   = clock.Second
)

type (
	Consumer func([]byte) error
	Service  interface {
		QueueDeclare(name string, durable bool, priority int, ttl time.Duration) error
		Publish(exchange, queue string, body []byte, mode, priority uint8) error
		Consume(queue string, auto bool, qos int, callback Consumer) error
	}
)
