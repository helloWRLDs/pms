package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Dev        bool   `env:"DEV"`
	Level      string `env:"LEVEL"`
	FileOutput bool   `env:"FILE_OUTPUT"`
	FilePath   string `env:"FILE_PATH"`
	StackTrace bool   `env:"STACKTRACE"`
	Caller     bool   `env:"CALLER"`
}

var Log *zap.SugaredLogger

func init() {
	WithConfig(
		WithDev(true),
		WithCaller(true),
		WithLevel("debug"),
	).Init()
}

// Set from env or config
func (conf *Config) Init() {
	var (
		encoding       string
		encConfig      zapcore.EncoderConfig
		outputChannels = []string{"stderr"}
	)
	if conf.Level == "" {
		conf.Level = "INFO"
	}
	if conf.FileOutput {
		outputChannels = append(outputChannels, conf.FilePath)
	}
	if conf.Dev {
		encoding = "console"
		encConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoding = "json"
		encConfig = zap.NewProductionEncoderConfig()
	}

	zapConf := zap.Config{
		Level: func(level string) zap.AtomicLevel {
			lvl, err := zap.ParseAtomicLevel(level)
			if err != nil {
				lvl = zap.NewAtomicLevel()
			}
			return lvl
		}(conf.Level),
		DisableCaller:     !conf.Caller,
		DisableStacktrace: !conf.StackTrace,
		OutputPaths:       outputChannels,
		ErrorOutputPaths:  outputChannels,
		Encoding:          encoding,
		EncoderConfig:     encConfig,
	}
	logger := zap.Must(zapConf.Build())
	defer logger.Sync()
	Log = logger.Sugar()
}
