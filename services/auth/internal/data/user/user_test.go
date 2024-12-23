package userdata

import (
	"context"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"pms.pkg/datastore/sqlite"
	"pms.pkg/types/ctx"
	"pms.pkg/utils"
)

const (
	dsn = "../../../data/users.db"
)

var (
	repo Repository
)

func TestMain(m *testing.M) {
	setupDB()
	code := m.Run()
	os.Exit(code)
}

func setupDB() {
	db, err := sqlite.Open(dsn)
	if err != nil {
		logrus.WithError(err).Fatal("failed to connect to db")
	}
	repo = *New(db)
}

func Test_GetByEmail(t *testing.T) {
	email := "john@doe.ru"
	ctx := ctx.New(context.Background())
	user, err := repo.GetByEmail(ctx, email)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(utils.JSON(user))
}

func Test_GetByID(t *testing.T) {
	id := "b417760e-de9b-4b1e-8b63-dfe32e5888c6"
	ctx := ctx.New(context.Background())
	user, err := repo.Get(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(utils.JSON(user))
}
