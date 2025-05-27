package logic

import (
	"context"
	"testing"

	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

func Test_RegisterUser(t *testing.T) {
	newUser := &dto.NewUser{
		Email:     "admin@example.com",
		FirstName: "admin",
		LastName:  "admin",
		Password:  "admin",
	}
	created, err := logic.RegisterUser(context.Background(), newUser)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(utils.JSON(created))
}

func Test_GetUser(t *testing.T) {
	profile, err := logic.GetProfile(context.Background(), "eb306dc5-52bb-4009-88af-347b4d040718")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(utils.JSON(profile))
}

func Test_UpdateUser(t *testing.T) {
	userID := "eb306dc5-52bb-4009-88af-347b4d040718"
	user, err := logic.GetProfile(context.Background(), userID)
	if err != nil {
		t.Fatal(err)
	}
	user.FirstName = "admin2"

	updated, err := logic.UpdateUser(context.Background(), userID, user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(utils.JSON(updated))
}
