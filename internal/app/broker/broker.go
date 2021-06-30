package broker

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Broker struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
	config     *BrokerConfig
	Logger     *logrus.Logger
}

func New(config *BrokerConfig) *Broker {
	return &Broker{
		config: config,
		Logger: logrus.New(),
	}
}

func (b *Broker) Open(queueName string) error {
	if err := b.configureLogger(); err != nil {
		return err
	}

	conn, err := amqp.Dial(b.config.URL)
	if err != nil {
		return err
	}
	b.connection = conn

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	b.channel = ch

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	b.queue = &q

	b.Logger.Info("amqp connection opened ", b.config.URL)
	return nil
}

func (b *Broker) Close() {
	b.connection.Close()
	b.channel.Close()
	b.Logger.Info("closed amqp connection ", b.config.URL)
}

func (b *Broker) configureLogger() error {
	level, err := logrus.ParseLevel(b.config.LogLevel)
	if err != nil {
		return err
	}
	b.Logger.SetLevel(level)

	b.Logger.Formatter = &logrus.TextFormatter{
		ForceColors: true,
	}
	return nil
}

func (b *Broker) Send(data []byte) error {
	if err := b.channel.Publish(
		"",
		b.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		}); err != nil {
		return err
	}
	return nil
}

func (b *Broker) Recieve() (<-chan amqp.Delivery, error) {
	msgs, err := b.channel.Consume(
		b.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}
