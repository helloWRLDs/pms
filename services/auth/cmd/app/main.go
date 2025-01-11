package main

import (
	"flag"

	"github.com/sirupsen/logrus"
	"pms.auth/internal/config"
	"pms.pkg/utils"
)

func main() {
	path := flag.String("path", "./services/auth/.env", "path to .env file")
	flag.Parse()

	// db, err := sqlite.Open(*dsn)
	// if err != nil {
	// 	print(err.Error())
	// 	os.Exit(1)
	// }
	// repo := userrepo.New(db)
	// user, err := domain.NewUser("john", "doe", "john@doe.com", "12345")
	// if err != nil {
	// 	println(utils.JSON(err))
	// 	os.Exit(1)
	// }
	// id, err := repo.Create(ctx.New(context.Background()), *user)
	// if err != nil {
	// 	println(utils.JSON(err))
	// 	os.Exit(1)
	// }
	// usr, err := repo.GetByID(ctx.New(context.Background()), id)
	// if err != nil {
	// 	println(utils.JSON(err))
	// 	os.Exit(1)
	// }
	// println(utils.JSON(usr))

	cfg, err := utils.LoadConfig[config.Config](*path)
	if err != nil {
		logrus.Fatal("failed to load config", "err", err)
	}
	logrus.Info("config loaded", utils.JSON(cfg))
}
