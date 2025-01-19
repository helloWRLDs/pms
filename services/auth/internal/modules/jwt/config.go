package jwt

type Config struct {
	TTL    int64  `env:"TTL"`
	Secret string `env:"SECRET"`
}
