package data

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.auth/internal/entity"
	"pms.pkg/consts"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

func Test_CreateRole(t *testing.T) {
	role := entity.Role{
		Name: "admin",
		Persmissions: consts.PermissionSet{
			consts.ORG_READ_PERMISSION,
			consts.ORG_WRITE_PERMISSION,
			consts.USER_DELETE_PERMISSION,
			consts.USER_READ_PERMISSION,
			consts.USER_WRITE_PERMISSION,
		},
		CompanyID: utils.Ptr("8f557202-0853-4672-aafb-a0b6cae7067a"),
	}
	err := repo.Role.Create(context.Background(), role)
	assert.NoError(t, err)
}

func Test_UpdateRole(t *testing.T) {
	role := entity.Role{
		Name: "admin-1",
		Persmissions: consts.PermissionSet{
			consts.ORG_READ_PERMISSION,
			consts.ORG_WRITE_PERMISSION,
			consts.USER_DELETE_PERMISSION,
			consts.USER_READ_PERMISSION,
			consts.USER_WRITE_PERMISSION,
		},
	}
	err := repo.Role.Update(context.Background(), "admin", role)
	assert.NoError(t, err)
	updated, err := repo.Role.GetByName(context.Background(), "admin-1")
	assert.NoError(t, err)
	t.Log(utils.JSON(updated))
}

func Test_GetRole(t *testing.T) {
	role, err := repo.Role.GetByName(context.Background(), "admin")
	assert.NoError(t, err)
	t.Log(utils.JSON(role))
}

func Test_Count(t *testing.T) {
	count := repo.Role.Count(context.Background(), list.Filters{})
	t.Log(count)
}

func Test_ListRoles(t *testing.T) {
	roles, err := repo.Role.List(context.Background(), list.Filters{
		Pagination: list.Pagination{
			Page:    1,
			PerPage: 10,
		},
		Fields: map[string]string{
			"company_id": "8f557202-0853-4672-aafb-a0b6cae7067a",
		},
	})
	assert.NoError(t, err)
	t.Log(utils.JSON(roles))
}
