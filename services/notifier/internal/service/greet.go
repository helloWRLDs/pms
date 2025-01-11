package service

import (
	"github.com/sirupsen/logrus"
	"pms.notifier/internal/modules/email/render"
)

func (s *NotifierService) GreetUser(name, email string) error {
	log := logrus.WithField("func", "GreetUser")

	greeting := render.GreetContent{
		Name:        name,
		CompanyName: "<company-name>",
	}

	data, err := render.Render(greeting)
	if err != nil {
		log.WithError(err).Error("failed to render email")
		return err
	}

	if err := s.Email.Send(data, email); err != nil {
		log.WithError(err).Error("failed to send email")
		return err
	}
	return nil
}
