package data

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

func Test_ListCompanies(t *testing.T) {
	list, err := repo.Company.List(context.Background(), list.Filters{
		Pagination: list.Pagination{
			Page:    1,
			PerPage: 10,
		},
		Fields: map[string]string{
			"p.user_id": "be10a73c-0927-4e3d-afe5-b4bae2e84946",
		},
	})
	assert.NoError(t, err)
	t.Log(utils.JSON(list))
}
