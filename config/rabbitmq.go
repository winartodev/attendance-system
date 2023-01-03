package config

import (
	"fmt"
	"os"
	"time"

	"github.com/labstack/gommon/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQ store message queue
type RabbitMQ struct {
	URL        string
	Exchange   string
	Conn       *amqp.Connection
	Chann      *amqp.Channel
	Queue      amqp.Queue
	closeChann chan *amqp.Error
	quitChann  chan bool
}

// NewRabbitMQ create connection with rabbitmq
func NewRabbitMQ(cfg Config) (*RabbitMQ, error) {
	addr := fmt.Sprintf("amqp://%s:%s@%s:%s/", os.Getenv("AMQP_USERNAME"), os.Getenv("AMQP_PASSWORD"), cfg.AMQP.Host, cfg.AMQP.Port)
	rmq := &RabbitMQ{
		URL:      addr,
		Exchange: "reminder.event.publish",
	}

	err := rmq.load()
	if err != nil {
		return nil, err
	}

	rmq.quitChann = make(chan bool)

	go rmq.handleDisconnect()

	return rmq, err
}

func (rmq *RabbitMQ) load() error {
	var err error

	rmq.Conn, err = amqp.Dial(rmq.URL)
	if err != nil {
		return err
	}

	rmq.Chann, err = rmq.Conn.Channel()
	if err != nil {
		return err
	}

	rmq.closeChann = make(chan *amqp.Error)
	rmq.Conn.NotifyClose(rmq.closeChann)

	// declare exchange if not exist
	err = rmq.Chann.ExchangeDeclare(rmq.Exchange, "direct", true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("%v declaring exchange %v", err, rmq.Exchange)
	}

	args := make(amqp.Table)
	args["x-delayed-type"] = "direct"
	err = rmq.Chann.ExchangeDeclare("delayed", "x-delayed-message", true, false, false, false, args)
	if err != nil {
		return fmt.Errorf("%v declaring exchange %v", err, "delayed")
	}

	return nil
}

// Shutdown will will close connection
func (rmq *RabbitMQ) Shutdown() {
	rmq.quitChann <- true

	log.Info("shutting down rabbitMQ's connection...")

	<-rmq.quitChann
}

// handleDisconnect will handle disconnect from server and try every 5 second
func (rmq *RabbitMQ) handleDisconnect() {
	for {
		select {
		case errChann := <-rmq.closeChann:
			if errChann != nil {
				log.Errorf("rabbitMQ disconnection: %v", errChann)
			}
		case <-rmq.quitChann:
			rmq.Conn.Close()
			log.Info("...rabbitMQ has been shut down")
			rmq.quitChann <- true
			return
		}

		log.Info("...trying to reconnect to rabbitMQ...")

		time.Sleep(5 * time.Second)

		if err := rmq.load(); err != nil {
			log.Errorf("rabbitMQ error: %v", err)
		}
	}
}
