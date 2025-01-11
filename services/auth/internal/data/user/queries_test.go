package userdata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.auth/internal/domain"
	"pms.auth/internal/domain/password"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
)

func Test_CreateUser(t *testing.T) {
	user := domain.User{
		FullName: "Bob",
		Email:    "bob@gmail.com",
	}
	pass, err := password.New("12345")
	if err != nil {
		t.Fatal(err)
	}
	user.Password = pass

	err = repo.CreateUser(context.Background(), user)
	assert.NoError(t, err)
}

func Test_GetByEmail(t *testing.T) {
	email := "bob@gmail.com"
	user, err := repo.GetByEmail(context.Background(), email)
	assert.NoError(t, err)
	t.Log(utils.JSON(user))
}

func Test_GetByID(t *testing.T) {
	id := "beb23917-e15e-48a2-8f7b-ebf85624aeb4"
	user, err := repo.GetByID(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(utils.JSON(user))
}

func Test_Transaction(t *testing.T) {
	user := domain.User{
		FullName: "Bob",
		Email:    "bob@gmail.com",
	}
	ctx, err := transaction.StartCTX(context.Background(), repo.DB)
	defer func() {
		transaction.EndCTX(ctx, err)
	}()
	assert.NoError(t, err)
	err = repo.CreateUser(ctx, user)
	assert.NoError(t, err)
}

func Test_Exists(t *testing.T) {
	t.Logf("is user exist: %v", repo.Exists(context.Background(), "bob@gmail.com"))
}
