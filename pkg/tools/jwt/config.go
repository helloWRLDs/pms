package jwt

type Config struct {
	TTL    int64  `env:"TTL"` // in hours
	Secret string `env:"SECRET"`
}

func WithTTL(ttl int64) func(*Config) {
	return func(c *Config) {
		c.TTL = ttl
	}
}

func WithSecret(secret string) func(*Config) {
	return func(c *Config) {
		c.Secret = secret
	}
}

func WithConfig(opts ...func(*Config)) *Config {
	var conf Config
	for _, fn := range opts {
		fn(&conf)
	}
	return &conf
}
