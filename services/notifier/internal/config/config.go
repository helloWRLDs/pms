package config

import "pms.notifier/internal/modules/email"

type Config struct {
	Host  string       `env:"HOST"`
	Gmail email.Config `envPrefix:"GMAIL_"`
}
