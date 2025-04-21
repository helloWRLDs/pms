package data

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"pms.auth/internal/entity"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

func Test_GetByEmail(t *testing.T) {
	email := "john@example.com"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user, err := repo.User.GetByEmail(ctx, email)
	assert.NoError(t, err)
	t.Log(utils.JSON(user))
}

func Test_SetAvatar(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user, err := repo.User.GetByID(ctx, "fb11170c-8f61-4fe5-858f-a5b256f6c1bd")
	if err != nil {
		t.Fatal(err)
	}
	avatar, err := os.ReadFile("./user.png")
	if err != nil {
		t.Fatal(err)
	}
	user.AvatarIMG = avatar
	err = repo.User.Update(ctx, user.ID.String(), user)
	assert.NoError(t, err)
}

func Test_ListUsers(t *testing.T) {
	list, err := repo.User.List(context.Background(), list.Filters{
		Pagination: list.Pagination{
			Page:    1,
			PerPage: 10,
		},
	})
	assert.NoError(t, err)
	t.Log(utils.JSON(list))
}

func Test_RegisterUser(t *testing.T) {
	user := entity.User{
		Name:     "John",
		Email:    "john@example.com",
		Password: "admin",
	}
	err := repo.User.Create(context.Background(), user)
	assert.NoError(t, err)
}
