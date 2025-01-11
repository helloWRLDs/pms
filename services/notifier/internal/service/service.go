package service

import (
	"pms.notifier/internal/modules/email"
)

type NotifierService struct {
	Email *email.Email
}

func New(conf email.Config) *NotifierService {
	return &NotifierService{
		Email: email.New(conf),
	}
}
