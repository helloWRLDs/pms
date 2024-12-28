package config

type Config struct {
	Host  string `env:"HOST"`
	Email string `envPrefix:"EMAIL_"`
}
