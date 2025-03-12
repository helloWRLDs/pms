package mqhandler

import (
	"context"

	"go.uber.org/zap"
	"pms.notifier/internal/config"
	"pms.notifier/internal/service"
	"pms.pkg/datastore/mq"
	notifiermq "pms.pkg/transport/mq/notifier"
)

type MessageQueueHandler struct {
	Sub *mq.Subscriber

	service *service.NotifierService

	conf *config.Config
	log  *zap.SugaredLogger
}

func New(conf *config.Config, logger *zap.SugaredLogger) (*MessageQueueHandler, error) {
	log := logger.With(
		zap.String("func", "mq.New"),
		zap.String("dsn", conf.AMQP.DSN),
		zap.String("exchange", conf.AMQP.Exchange),
	)
	log.Debug("mq.New called")

	sub, err := mq.NewSubscriber(context.Background(), mq.SubscriberOpts{
		Queue:  "notifier",
		Routes: notifiermq.Routes,
		Config: conf.AMQP,
		Log:    logger,
	})
	if err != nil {
		log.Errorw("failed to create subscriber", "err", err)
		return nil, err
	}
	log.Debug("subscriber created")

	mqh := &MessageQueueHandler{
		conf:    conf,
		log:     logger,
		Sub:     sub,
		service: service.New(conf.Gmail),
	}

	return mqh, nil
}

func (m *MessageQueueHandler) Close() {
	m.Sub.Close()
}

func (m *MessageQueueHandler) Listen(ctx context.Context) error {
	log := m.log.With(
		zap.String("func", "mq.Listen"),
	)
	log.Debug("mq.Listen called")

	msgs, err := m.Sub.Consume(ctx)
	if err != nil {
		log.Errorw("failed to consume messages", "err", err)
		return err
	}

	go func() {
		log.Info("starting message queue consumption")
		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					log.Warn("msg channel is closed")
					return
				}
				if err := m.HandleMessage(context.Background(), &msg); err != nil {
					log.Errorw("failed to process message", "err", err)
					_ = msg.Nack(false, true) // Requeue message
				} else {
					log.Debugw("message processed successfully")
					_ = msg.Ack(false) // Acknowledge message
				}
			case <-ctx.Done():
				log.Warn("stopping message queue consumption")
				return
			}
		}
	}()

	return nil
}
