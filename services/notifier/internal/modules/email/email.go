package email

import (
	"fmt"
	"net/smtp"
)

type Email struct {
	Conf Config
	Auth smtp.Auth
}

func New(config Config) *Email {
	return &Email{
		Conf: config,
		Auth: smtp.PlainAuth("", config.Username, config.Password, config.Host),
	}
}

func (e *Email) Send(data []byte, to ...string) error {
	err := smtp.SendMail(fmt.Sprintf("%s:%s", e.Conf.Host, e.Conf.Port), e.Auth, e.Conf.Username, to, data)
	return err
}
