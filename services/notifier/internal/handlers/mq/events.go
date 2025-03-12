package mqhandler

import (
	"context"

	"go.uber.org/zap"
	notifiermq "pms.pkg/transport/mq/notifier"
)

func (m *MessageQueueHandler) HandleGreetEvent(ctx context.Context, event notifiermq.GreetMessage) error {
	log := m.log.With(
		zap.String("func", "mq.HandleGreetEvent"),
		zap.String("email", event.MetaData.ToEmail),
		zap.String("name", event.Name),
	)
	log.Debug("mq.HandleGreetEvent called")

	if err := m.service.GreetUser(ctx, event.Name, event.ToEmail); err != nil {
		log.Errorw("failed to send greet email", "err", err)
		return err
	}
	log.Debug("greet email sent")
	return nil
}
