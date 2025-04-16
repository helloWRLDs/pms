package userdata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.pkg/type/list"
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
func Test_Count(t *testing.T) {
	filter := list.Filters{
		Created: list.Created{
			CreatedFrom: "2025-01-01",
			CreatedTo:   "2025-12-31",
		},
	}

	count, err := repo.Count(context.Background(), filter)
	assert.NoError(t, err)
	t.Log("Filtered user count:", count)
}
func Test_ListUsers(t *testing.T) {
	filter := list.Filters{
		Created: list.Created{
			CreatedFrom: "2024-01-01",
			CreatedTo:   "2024-12-31",
		},
		OrderBy: "name ASC",
		Pagination: list.Pagination{
			Page:    1,
			PerPage: 10,
		},
	}

	users, err := repo.ListUsers(context.Background(), filter)
	assert.NoError(t, err)
	t.Log("Users list:", utils.JSON(users))
}
