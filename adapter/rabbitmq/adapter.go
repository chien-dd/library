package rabbitmq

import (
	"github.com/chien-dd/library/clock"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

type (
	Connection struct {
		*amqp.Connection
	}

	Channel struct {
		*amqp.Channel
	}
)

func (rb *Rabbit) Connect() (*Connection, error) {
	url := rb.config.String()
	con, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	connection := &Connection{Connection: con}
	// Reconnect
	go func() {
		for {
			if connection.Connection != nil {
				reason, ok := <-connection.NotifyClose(make(chan *amqp.Error))
				if !ok {
					log.Info("[Rabbit] Connection closed")
					break
				}
				if reason != nil {
					log.Infof("[Rabbit] Connection closed. Reason: %v", reason)
				}
			}
			for {
				clock.Sleep(delayReconnect)
				// Recreate channel
				if con, err := amqp.Dial(url); err == nil {
					log.Info("[Rabbit] Recreate connection success!")
					connection.Connection = con
					break
				} else {
					log.Errorf("[Rabbit] Recreate connection failed. Reason: %v", err)
				}
			}
		}
	}()
	// Success
	return connection, nil
}

func (con *Connection) Channel() (*Channel, error) {
	ch, err := con.Connection.Channel()
	if err != nil {
		return nil, err
	}
	channel := &Channel{Channel: ch}
	// Reconnect
	go func() {
		for {
			if channel.Channel != nil {
				reason, ok := <-channel.NotifyClose(make(chan *amqp.Error))
				if !ok {
					log.Info("[Rabbit] Channel closed")
					break
				}
				if reason != nil {
					log.Infof("[Rabbit] Channel closed. Reason: %v", reason)
				}
			}
			for {
				clock.Sleep(delayReconnect)
				// Recreate channel
				if ch, err := con.Connection.Channel(); err == nil {
					log.Info("[Rabbit] Recreate channel success!")
					channel.Channel = ch
					break
				} else {
					log.Errorf("[Rabbit] Recreate channel failed. Reason: %v", err)
				}
			}
		}
	}()
	// Success
	return channel, nil
}
