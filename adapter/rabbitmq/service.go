package rabbitmq

import (
	"github.com/chien-dd/library/clock"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
	"time"
)

type (
	Rabbit struct {
		config     Config
		connection *Connection
		channel    *Channel
	}
)

func New(conf Config) (Service, error) {
	rb := &Rabbit{config: conf}
	con, err := rb.Connect()
	if err != nil {
		log.Errorf("[Rabbit] Create connection failed. Reason: %v", err)
		return nil, err
	}
	ch, err := con.Channel()
	if err != nil {
		log.Errorf("[Rabbit] Create channel failed. Reason: %v", err)
	}
	rb.connection = con
	rb.channel = ch
	// Success
	return rb, nil
}

func (rb *Rabbit) QueueDeclare(name string, durable bool, priority int, ttl time.Duration) error {
	table := amqp.Table{}
	if priority > 0 {
		table["x-max-priority"] = priority
	}
	if ttl > 0 {
		table["message-ttl"] = ttl.Milliseconds()
	}
	_, err := rb.channel.QueueDeclare(name, durable, false, false, false, table)
	if err != nil {
		return err
	}
	// Success
	return nil
}

func (rb *Rabbit) Publish(exchange, queue string, body []byte, mode, priority uint8) error {
	err := rb.channel.Publish(
		exchange,
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType:  echo.MIMETextPlain,
			DeliveryMode: mode,
			Priority:     priority,
			Body:         body,
		},
	)
	if err != nil {
		clock.Sleep(delayPublish)
		return rb.Publish(exchange, queue, body, mode, priority)
	}
	// Success
	return nil
}

func (rb *Rabbit) Consume(queue string, auto bool, prefetchCount int, callback Consumer) error {
	err := rb.channel.Qos(prefetchCount, 0, false)
	if err != nil {
		log.Errorf("[Rabbit] Qos failed. Reason: %v", err)
		clock.Sleep(delayConsume)
		return rb.Consume(queue, auto, prefetchCount, callback)
	}
	deliveries := make(chan amqp.Delivery)
	go func() {
		for {
			d, err := rb.channel.Consume(queue, "", auto, false, false, false, amqp.Table{})
			if err != nil {
				clock.Sleep(delayConsume)
				continue
			}
			for msg := range d {
				deliveries <- msg
			}
		}
	}()
	// Process
	for msg := range deliveries {
		err := callback(msg.Body)
		if err == nil {
			if !auto {
				msg.Ack(false)
			}
		} else {
			if !auto {
				msg.Nack(false, true)
			}
		}
	}
	return err
}
