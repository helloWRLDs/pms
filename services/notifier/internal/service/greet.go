package service

import (
	"context"

	"go.uber.org/zap"
	"pms.notifier/internal/modules/email/render"
	"pms.pkg/logger"
)

func (s *NotifierService) GreetUser(ctx context.Context, name, email string) error {
	log := logger.Log.With(
		zap.String("func", "service.GreetUser"),
		zap.String("name", name),
		zap.String("email", email),
	)
	log.Debug("service.GreetUser called")

	greeting := render.GreetContent{
		Name:        name,
		CompanyName: "<company-name>",
	}

	data, err := render.Render(greeting)
	if err != nil {
		log.Errorw("failed to render email", "err", err)
		return err
	}

	if err := s.Email.Send(data, email); err != nil {
		log.Errorw("failed to send email", "err", err)
		return err
	}
	log.Debug("email sent successfuly")
	return nil
}
