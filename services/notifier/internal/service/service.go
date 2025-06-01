package service

import (
	"context"

	"go.uber.org/zap"
	"pms.notifier/internal/modules/email"
	"pms.notifier/internal/modules/email/render"
	"pms.pkg/logger"
)

type NotifierService struct {
	Email *email.Email
}

func New(conf email.Config) *NotifierService {
	return &NotifierService{
		Email: email.New(conf),
	}
}

func (s *NotifierService) NotifyTaskAssignment(ctx context.Context, assigneeName, assigneeEmail, taskName, taskId, projectName string) error {
	log := logger.Log.With(
		zap.String("func", "service.NotifyTaskAssignment"),
		zap.String("assignee", assigneeName),
		zap.String("email", assigneeEmail),
		zap.String("task", taskName),
	)
	log.Debug("service.NotifyTaskAssignment called")

	content := render.TaskAssignmentContent{
		AssigneeName: assigneeName,
		TaskName:     taskName,
		TaskId:       taskId,
		ProjectName:  projectName,
		CompanyName:  "TaskFlow",
	}

	data, err := render.Render(content)
	if err != nil {
		log.Errorw("failed to render email", "err", err)
		return err
	}

	if err := s.Email.Send(data, assigneeEmail); err != nil {
		log.Errorw("failed to send email", "err", err)
		return err
	}
	log.Debug("email sent successfully")
	return nil
}
