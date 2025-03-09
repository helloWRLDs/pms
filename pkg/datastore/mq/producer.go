package mq

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	mqtp "pms.pkg/transport/mq"
	"pms.pkg/utils"
)

type Publisher struct {
	Queue string

	Ch   *amqp.Channel
	Conn *amqp.Connection
	log  *zap.SugaredLogger
}

type PublisherOpts struct {
	Queue  mqtp.QueueName
	Config Config
	Logger *zap.SugaredLogger
}

func NewPublisher(ctx context.Context, opts PublisherOpts) (*Publisher, error) {
	log := opts.Logger.With(
		zap.String("func", "mq.NewPublisher"),
		zap.String("queue", string(opts.Queue)),
	)
	log.Debug("mq.NewPublisher called")

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

	go func() {
		for {
			select {
			case cancel := <-ch.NotifyCancel(make(chan string)):
				log.Errorw("received cancel", "msg", cancel)
				return
			case <-ctx.Done():
				log.Debug("context done")
				return
			case <-conn.NotifyClose(make(chan *amqp.Error)):
				log.Debug("connection closed")
				return
			}
		}
	}()

	return &Publisher{
		Ch:    ch,
		Conn:  conn,
		log:   log,
		Queue: string(opts.Queue),
	}, nil
}

func (p *Publisher) Publish(ctx context.Context, msg mqtp.Queueable) error {
	log := p.log.With(
		zap.String("func", "mq.Publish"),
		zap.String("routing key", string(msg.RoutingKey())),
		zap.String("queue", utils.JSON(msg)),
	)
	log.Debug("mq.Publish called")

	body, err := json.Marshal(msg)
	if err != nil {
		log.Errorw("failed to marshal message", "err", err)
		return err
	}
	err = p.Ch.PublishWithContext(
		ctx,
		"direct-exchange",
		string(msg.RoutingKey()),
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Errorw("failed to publish message", "err", err)
		return err
	}
	return nil
}

func (p *Publisher) ConnState() string {
	if p.Conn.IsClosed() || p.Ch.IsClosed() {
		return "CLOSED"
	}
	return "READY"
}

func (p *Publisher) Close() error {
	log := p.log.With(
		zap.String("func", "mq.Close"),
	)
	log.Debug("mq.Close called")
	if err := p.Ch.Close(); err != nil {
		log.Errorw("failed to close channel", "err", err)
		return err
	}
	if err := p.Conn.Close(); err != nil {
		log.Errorw("failed to close connection", "err", err)
		return err
	}
	return nil
}
