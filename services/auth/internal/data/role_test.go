package data

import (
	"context"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	roledata "pms.auth/internal/data/role"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

func Test_CreateRole(t *testing.T) {
	role := roledata.Role{
		Name: "admin-aitu",
		Permissions: pq.StringArray{
			string(consts.COMPANY_READ_PERMISSION),
			string(consts.COMPANY_WRITE_PERMISSION),
			string(consts.COMPANY_DELETE_PERMISSION),
			string(consts.COMPANY_INVITE_PERMISSION),
			string(consts.USER_DELETE_PERMISSION),
			string(consts.USER_READ_PERMISSION),
			string(consts.USER_WRITE_PERMISSION),
		},
		CompanyID: utils.Ptr("dee3b9c8-b6a4-4106-9304-525b3da7dc30"),
	}
	err := repo.Role.Create(context.Background(), role)
	assert.NoError(t, err)
}

func Test_UpdateRole(t *testing.T) {
	role := roledata.Role{
		Name: "admin-1",
		Permissions: pq.StringArray{
			string(consts.COMPANY_READ_PERMISSION),
			string(consts.COMPANY_WRITE_PERMISSION),
			string(consts.COMPANY_DELETE_PERMISSION),
			string(consts.COMPANY_INVITE_PERMISSION),
			string(consts.USER_DELETE_PERMISSION),
			string(consts.USER_READ_PERMISSION),
			string(consts.USER_WRITE_PERMISSION),
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
	roles, err := repo.Role.List(context.Background(), &dto.RoleFilter{
		Page:      1,
		PerPage:   10,
		CompanyId: "dee3b9c8-b6a4-4106-9304-525b3da7dc30",
		UserId:    "146763a5-06f7-43c0-af23-615fc120563d",
		// WithDefault: true,
	})
	assert.NoError(t, err)
	t.Log(utils.JSON(roles))
}
