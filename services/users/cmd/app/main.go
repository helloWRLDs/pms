package main

import (
	"context"
	"flag"
	"os"

	"pms.pkg/datastore/sqlite"
	"pms.pkg/types/ctx"
	"pms.pkg/utils"
	"pms.users/internal/domain"
	userrepo "pms.users/internal/repository"
)

func main() {
	dsn := flag.String("dsn", "./users.db", "")
	flag.Parse()

	db, err := sqlite.Open(*dsn)
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
	repo := userrepo.New(db)
	user, err := domain.NewUser("john", "doe", "john@doe.com", "12345")
	if err != nil {
		println(utils.JSON(err))
		os.Exit(1)
	}
	id, err := repo.Create(ctx.New(context.Background()), *user)
	if err != nil {
		println(utils.JSON(err))
		os.Exit(1)
	}
	usr, err := repo.GetByID(ctx.New(context.Background()), id)
	if err != nil {
		println(utils.JSON(err))
		os.Exit(1)
	}
	println(utils.JSON(usr))
}
