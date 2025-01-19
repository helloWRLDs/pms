package service

import (
	"context"

	"github.com/sirupsen/logrus"
	"pms.notifier/internal/modules/email/render"
)

func (s *NotifierService) GreetUser(ctx context.Context, name, email string) error {
	log := logrus.WithFields(logrus.Fields{
		"func":  "GreetUser",
		"name":  name,
		"email": email,
	})
	log.Debug("GreetUser called")

	greeting := render.NewGreetContent(name, "<company-name>")

	data, err := render.Render(greeting)
	if err != nil {
		log.WithError(err).Error("failed to render email")
		return err
	}

	if err := s.Email.Send(data, email); err != nil {
		log.WithError(err).Error("failed to send email")
		return err
	}
	log.Debug("email sent successfuly")
	return nil
}
