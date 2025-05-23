package configgrpc

type ClientConfig struct {
	Host       string `env:"HOST"`
	DisableLog bool   `env:"DISABLE_LOG"`
}
