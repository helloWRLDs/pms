package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

func Test_GetProfile(t *testing.T) {
	userID := "be10a73c-0927-4e3d-afe5-b4bae2e84946"
	profile, err := logic.GetProfile(context.Background(), userID)
	assert.NoError(t, err)
	t.Log(utils.JSON(profile))
}

func Test_ListUsers(t *testing.T) {
	users, err := logic.ListUsers(context.Background(), &dto.UserFilter{
		Page:      1,
		PerPage:   10,
		CompanyId: "60cde332-ad5a-4aab-932b-81b5f16a61d2",
	})
	assert.NoError(t, err)
	t.Log(utils.JSON(users))
}
