package config

import (
	"pms.notifier/internal/modules/email"
	"pms.pkg/datastore/mq"
	"pms.pkg/logger"
)

type Config struct {
	Host  string       `env:"HOST"`
	Gmail email.Config `envPrefix:"GMAIL_"`
	AMQP  mq.Config    `envPrefix:"MQ_"`

	Log logger.Config `envPrefix:"LOG_"`
}
