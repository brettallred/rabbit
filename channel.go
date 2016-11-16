package rabbit

import (
	"github.com/streadway/amqp"
	"log"
)

func createChannel(conn *amqp.Connection, autoCloseConnection bool) (*amqp.Channel, error) {
	if conn == nil {
		log.Printf("Failed to open a channel: no connection")
		return nil, nil //TODO Remove this guard clause
	}

	channel, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	if autoCloseConnection {
		errorChannel := make(chan *amqp.Error, 1)
		errorHandler := func() {
			select {
			case <-errorChannel:
				lock.Lock()
				defer lock.Unlock()
				_connection.Close()
				return
			}
		}
		channel.NotifyClose(errorChannel)
		go errorHandler()
	}

	return channel, nil
}
