package data

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	userdata "pms.auth/internal/data/user"
	"pms.pkg/transport/grpc/dto"
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

func Test_CountUsers(t *testing.T) {
	count := repo.Company.Count(context.Background(), &dto.CompanyFilter{
		CompanyId: "60cde332-ad5a-4aab-932b-81b5f16a61d2",
	})
	t.Log(count)
}

func Test_CreateUser(t *testing.T) {
	hashed, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	id := uuid.NewString()
	t.Log("id=", id)
	user := userdata.User{
		ID:        id,
		FirstName: "admin",
		LastName:  utils.Ptr("admin"),
		Email:     "admin@example.com",
		Password:  utils.Ptr(string(hashed)),
	}
	if err := repo.User.Create(context.Background(), user); err != nil {
		t.Fatal(err)
	}
	t.Log("user created")
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
	user.AvatarIMG = &avatar
	err = repo.User.Update(ctx, user.ID, user)
	assert.NoError(t, err)
}

func Test_UserExists(t *testing.T) {
	exists := repo.User.Exists(context.Background(), "email", "admin@example.com")
	t.Log(exists)
}

func Test_ListUsers(t *testing.T) {
	list, err := repo.User.List(context.Background(), &dto.UserFilter{
		Page:    1,
		PerPage: 10,
	})
	assert.NoError(t, err)
	t.Log(utils.JSON(list))
}
