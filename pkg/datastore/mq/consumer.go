package mq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	mqtp "pms.pkg/transport/mq"
)

type SubscriberOpts struct {
	Queue  string
	Routes []mqtp.QueueRoute
	Config Config
	Log    *zap.SugaredLogger
}

type Subscriber struct {
	Queue string

	Ch   *amqp.Channel
	Conn *amqp.Connection
	log  *zap.SugaredLogger
}

func NewSubscriber(ctx context.Context, opts SubscriberOpts) (*Subscriber, error) {
	log := opts.Log.With(
		zap.String("func", "mq.NewSubscriber"),
		zap.String("queue", opts.Queue),
		zap.String("dsn", opts.Config.DSN),
		zap.String("exchange", opts.Config.Exchange),
	)
	log.Debug("mq.NewSubscriber called")

	conn, err := amqp.Dial(opts.Config.DSN)
	if err != nil {
		log.Errorw("failed connecting to rabbitmq", "err", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Errorw("failed to open channel", "err", err)
		return nil, err
	}

	if err = ch.ExchangeDeclare(
		opts.Config.Exchange, // Exchange name
		"direct",             // Type
		true,                 // Durable
		false,                // Auto-delete
		false,                // Internal
		false,                // No-wait
		nil,                  // Arguments
	); err != nil {
		log.Errorw("failed to declare exchange", "err", err)
		return nil, err
	}

	if _, err = ch.QueueDeclare(
		opts.Queue,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		log.Errorw("failed to declare queue", "err", err)
		return nil, err
	}

	for _, route := range opts.Routes {
		err = ch.QueueBind(
			opts.Queue,
			string(route),
			opts.Config.Exchange,
			false,
			nil,
		)
		if err != nil {
			log.Errorw("failed to bind queue", "bind", route, "err", err)
			return nil, err
		}
	}

	sub := &Subscriber{
		Ch:    ch,
		Conn:  conn,
		log:   log,
		Queue: opts.Queue,
	}

	return sub, nil
}

func (s *Subscriber) Consume(ctx context.Context) (<-chan amqp.Delivery, error) {
	log := s.log.With(zap.String("func", "mq.Subscriber.Consume"))
	log.Debug("Starting Consumer...")

	msgs, err := s.Ch.Consume(
		s.Queue,
		"",
		true,  // Auto-Ack (set to false if you want manual Ack)
		false, // Exclusive
		false, // No-local
		false, // No-wait
		nil,   // Args
	)
	if err != nil {
		log.Errorw("failed to start consumer", "err", err)
		return nil, err
	}

	return msgs, nil
}

func (s *Subscriber) Close() {
	log := s.log.With(
		zap.String("func", "mq.Close"),
	)
	log.Debug("mq.Close called")
	if err := s.Ch.Close(); err != nil {
		log.Errorw("failed to close channel", "err", err)
	}
	if err := s.Conn.Close(); err != nil {
		log.Errorw("failed to close connection", "err", err)
	}
}
