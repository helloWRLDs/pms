package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Dev  bool
	Path string
}

func (c Config) Init() {
	if c.Dev || c.Path == "" {
		log.SetFormatter(&log.TextFormatter{
			DisableColors: false,
			FullTimestamp: false,
		})
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.ErrorLevel)
		f, err := os.OpenFile(c.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			panic(err)
		}

		log.SetOutput(f)
	}
}

func (c Config) Close() {
	if !c.Dev && c.Path != "" {
		if f, ok := log.StandardLogger().Out.(*os.File); ok {
			f.Close()
		}
	}
}
