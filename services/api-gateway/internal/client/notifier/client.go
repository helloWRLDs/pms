package notifierclient

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/datastore/mq"
	notifiermq "pms.pkg/transport/mq/notifier"
)

func New(conf mq.Config, logger *zap.SugaredLogger) (*mq.Publisher, error) {
	log := logger.With(
		zap.String("func", "notifierclient.New"),
	)
	log.Debug("notifierclient.New called")

	pub, err := mq.NewPublisher(context.Background(), mq.PublisherOpts{
		Config: conf,
		Logger: log,
		Queue:  notifiermq.Queue,
	})
	if err != nil {
		log.Debugw("failed to create notifier pub", "err", err)
		return nil, err
	}
	log.Debug("notification pub created")
	return pub, nil
}
