package producermq

import (
	"github.com/streadway/amqp"
)

type RabbitMQProducer interface {
	ProducerMessage(queueName string, message []byte) error
	Close() error
}

type RabbitMQProducerImpl struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQProducer(amqpURI string) (*RabbitMQProducerImpl, error) {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQProducerImpl{
		conn:    conn,
		channel: channel,
	}, nil
}

func (r *RabbitMQProducerImpl) ProducerMessage(queueName string, message []byte) error {
	err := r.channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQProducerImpl) Close() error {
	return r.conn.Close()
}
