package mq

type Config struct {
	DSN        string `env:"DSN"`
	Exchange   string `env:"EXCHANGE"`
	DisableLog bool   `env:"DISABLE_LOG"`
}
