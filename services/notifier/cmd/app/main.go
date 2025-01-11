package main

import (
	"flag"

	"pms.notifier/internal/config"
	"pms.pkg/utils"
)

func main() {
	envPath := flag.String("env path", "./services/notifier/.env", "")
	flag.Parse()

	conf, err := utils.LoadConfig[config.Config](*envPath)
	if err != nil {
		panic(err)
	}
	print(utils.JSON(conf))
}
