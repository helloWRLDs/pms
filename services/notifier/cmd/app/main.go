package main

import (
	"context"
	"flag"

	"pms.notifier/internal/config"
	"pms.notifier/internal/service"
	"pms.pkg/utils"
)

func main() {
	envPath := flag.String("env path", "./services/notifier/.env", "")
	flag.Parse()

	conf, err := utils.LoadConfig[config.Config](*envPath)
	if err != nil {
		panic(err)
	}
	notifer := service.New(conf.Gmail)
	notifer.GreetUser(context.Background(), "John", "danil.li24x@gmail.com")
}
