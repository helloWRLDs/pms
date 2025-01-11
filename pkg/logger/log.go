package logger

import (
	"os"
	"path"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Dev  bool   `env:"DEV"`
	Path string `env:"PATH"`
}

func Init(cfg Config) {
	if cfg.Dev || cfg.Path == "" {
		log.SetFormatter(&log.TextFormatter{
			DisableColors: false,
			FullTimestamp: false,
			CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
				s := strings.Split(f.Function, ".")
				funcname := s[len(s)-1]
				_, filename := path.Split(f.File)
				return funcname, filename
			},
		})
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.ErrorLevel)
		f, err := os.OpenFile(cfg.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			panic(err)
		}
		log.SetOutput(f)
	}
}

func Close(conf Config) {
	if !conf.Dev && conf.Path != "" {
		if f, ok := log.StandardLogger().Out.(*os.File); ok {
			f.Close()
		}
	}
}
