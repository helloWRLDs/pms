package userdata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.pkg/utils"
)

// func Test_CreateUser(t *testing.T) {
// 	user := entity.User{
// 		Name:     "Bob",
// 		Email:    "bob@gmail.com",
// 		Password: "password",
// 	}
// 	err := repo.Create(context.Background(), user)
// 	assert.NoError(t, err)
// 	fetched, err := repo.GetByEmail(context.Background(), user.Email)
// 	assert.NoError(t, err)
// 	assert.Equal(t, user, fetched)
// }

func Test_GetUser(t *testing.T) {
	email := "bob@gmail.ru"
	user, err := repo.GetByEmail(context.Background(), email)
	assert.NoError(t, err)
	t.Log(utils.JSON(user))
}

func Test_Exists(t *testing.T) {
	email := "bob@gmail.com"
	t.Log(repo.Exists(context.Background(), email))
}
