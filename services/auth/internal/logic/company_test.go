package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

func Test_CreateCompany(t *testing.T) {
	userID := "eb306dc5-52bb-4009-88af-347b4d040718"
	newCompany := &dto.NewCompany{
		Name:     "test",
		Codename: "test",
	}
	created, err := logic.CreateCompany(context.Background(), userID, newCompany)
	assert.NoError(t, err)
	t.Log(utils.JSON(created))
}

func Test_GetCompany(t *testing.T) {
	id := "60cde332-ad5a-4aab-932b-81b5f16a61d2"
	comp, err := logic.GetCompany(context.Background(), id)
	assert.NoError(t, err)
	t.Log(utils.JSON(comp))
}
