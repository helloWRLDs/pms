package logger

func WithConfig(opts ...func(*Config)) *Config {
	var conf Config
	for _, fn := range opts {
		fn(&conf)
	}
	return &conf
}

func WithDev(dev bool) func(*Config) {
	return func(c *Config) {
		c.Dev = dev
	}
}

func WithLevel(level string) func(*Config) {
	return func(c *Config) {
		c.Level = level
	}
}

func WithFile(enabled bool, filpath string) func(*Config) {
	return func(c *Config) {
		c.FileOutput = enabled
		c.FilePath = filpath
	}
}

func WithStack(enabled bool) func(*Config) {
	return func(c *Config) {
		c.StackTrace = enabled
	}
}

func WithCaller(enabled bool) func(*Config) {
	return func(c *Config) {
		c.Caller = enabled
	}
}
