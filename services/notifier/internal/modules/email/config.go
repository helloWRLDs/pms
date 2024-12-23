package email

type Config struct {
	Host     string `env:"HOST"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Port     string `env:"PORT"`
}
