package mqhandler

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	mqtp "pms.pkg/transport/mq"
	notifiermq "pms.pkg/transport/mq/notifier"
)

func (m *MessageQueueHandler) HandleMessage(ctx context.Context, msg *amqp.Delivery) error {
	log := m.log.With(
		zap.String("func", "mq.HandleMessage"),
		zap.String("routing_key", msg.RoutingKey),
	)
	log.Debug("mq.HandleMessage called")

	switch mqtp.QueueRoute(msg.RoutingKey) {
	case (notifiermq.GreetMessage{}).RoutingKey():
		var event notifiermq.GreetMessage
		if err := json.Unmarshal(msg.Body, &event); err != nil {
			log.Errorw("failed to unmarshal message", "err", err)
			return err
		}
		return m.HandleGreetEvent(ctx, event)
	case (notifiermq.TaskAssignmentMessage{}).RoutingKey():
		var event notifiermq.TaskAssignmentMessage
		if err := json.Unmarshal(msg.Body, &event); err != nil {
			log.Errorw("failed to unmarshal message", "err", err)
			return err
		}
		return m.HandleTaskAssignmentEvent(ctx, event)
	default:
		log.Warnw("unhandled routing key", "key", msg.RoutingKey)
		return nil
	}
}
