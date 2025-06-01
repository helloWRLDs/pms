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

func (m *MessageQueueHandler) HandleTaskAssignmentEvent(ctx context.Context, event notifiermq.TaskAssignmentMessage) error {
	log := m.log.With(
		zap.String("func", "mq.HandleTaskAssignmentEvent"),
		zap.String("email", event.MetaData.ToEmail),
		zap.String("task_id", event.TaskId),
	)
	log.Debug("mq.HandleTaskAssignmentEvent called")

	if err := m.service.NotifyTaskAssignment(ctx, event.AssigneeName, event.ToEmail, event.TaskName, event.TaskId, event.ProjectName); err != nil {
		log.Errorw("failed to send task assignment email", "err", err)
		return err
	}
	log.Debug("task assignment email sent")
	return nil
}
