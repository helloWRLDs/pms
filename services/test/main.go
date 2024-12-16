package main

import (
	"github.com/sirupsen/logrus"
	"pms.pkg/logger"
)

func main() {
	cfg := logger.Config{
		Dev:  true,
		Path: "./test.log",
	}
	cfg.Init()
	defer cfg.Close()

	// dsn := flag.String("dsn", "./test.db", "")
	// flag.Parse()

	// _, err := sqlite.Open(*dsn)
	// if err != nil {
	// 	print(err.Error())
	// }

	logrus.Info("server started")
}
