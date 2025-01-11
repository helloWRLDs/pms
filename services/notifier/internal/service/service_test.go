package service

import (
	"os"
	"testing"

	"pms.pkg/utils"

	"pms.notifier/internal/config"
	"pms.notifier/internal/modules/email"
)

var (
	Notifier *NotifierService
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	conf, err := utils.LoadConfig[config.Config]("../../.env")
	if err != nil {
		panic(err)
	}
	print("loaded conf: \n", utils.JSON(conf))
	Notifier = &NotifierService{
		Email: email.New(conf.Gmail),
	}
}
