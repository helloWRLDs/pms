package userdata

import (
	"context"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"pms.auth/internal/domain/user/password"
	userentity "pms.auth/internal/entity/user"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
)

func Test_CreateUser(t *testing.T) {
	user := userentity.User{
		Name:  "Bob",
		Email: "bob@gmail.com",
	}
	pass, err := password.New("12345")
	if err != nil {
		t.Fatal(err)
	}
	user.Password = pass
	id := uuid.New()
	user.ID = id
	log.Info(id)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = repo.Create(
		ctx,
		user,
	)
	assert.NoError(t, err)
}

func Test_GetByEmail(t *testing.T) {
	email := "bob@gmail.com"
	user, err := repo.GetByEmail(context.Background(), email)
	assert.NoError(t, err)
	t.Log(utils.JSON(user))
	t.Log(user.CreatedAt.Format)
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
	user := userentity.User{
		Name:  "Bob",
		Email: "bob@gmail.com",
	}
	ctx, err := transaction.StartCTX(context.Background(), repo.DB)
	defer func() {
		transaction.EndCTX(ctx, err)
	}()
	assert.NoError(t, err)
	err = repo.Create(ctx, user)
	assert.NoError(t, err)
}

func Test_Exists(t *testing.T) {
	t.Logf("is user exist: %v", repo.Exists(context.Background(), "bob@gmail.com"))
}
