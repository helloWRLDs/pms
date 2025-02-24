package roledata

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	roledomain "pms.auth/internal/domain/role"
	"pms.pkg/type/permission"
	"pms.pkg/utils"
)

func Test_CreateRole(t *testing.T) {
	role := roledomain.Role{
		Name:        "admin",
		Permissions: permission.PermissionSet{permission.ORG_MANAGE, permission.ORG_READ, permission.ORG_UPDATE},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.Create(ctx, role)
	assert.NoError(t, err)
}

func Test_GetRole(t *testing.T) {
	id := 1
	role, err := repo.GetByID(context.Background(), id)
	assert.NoError(t, err)
	t.Log(utils.JSON(role))
}
