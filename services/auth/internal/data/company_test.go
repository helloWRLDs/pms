package data

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"pms.auth/internal/entity"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

func Test_ListCompanies(t *testing.T) {
	list, err := repo.Company.List(context.Background(), list.Filters{
		Pagination: list.Pagination{
			Page:    1,
			PerPage: 10,
		},
	})
	assert.NoError(t, err)
	t.Log(utils.JSON(list))
}

func Test_CreateCompany(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newCompany := entity.Company{
		Name:     "AITU",
		Codename: "aitu",
	}

	err := repo.Company.Create(ctx, newCompany)
	assert.NoError(t, err)
}
