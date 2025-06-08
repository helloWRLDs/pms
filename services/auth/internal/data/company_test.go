package data

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	companydata "pms.auth/internal/data/company"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

func Test_ListCompanies(t *testing.T) {
	list, err := repo.Company.List(context.Background(), &dto.CompanyFilter{
		Page:        1,
		PerPage:     10,
		CompanyName: "TEst",
	})
	assert.NoError(t, err)
	t.Log(utils.JSON(list))
}

func Test_CreateCompany(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newCompany := companydata.Company{
		Name:     "AITU",
		Codename: "aitu",
	}

	err := repo.Company.Create(ctx, newCompany)
	assert.NoError(t, err)
}

func Test_ExistCompany(t *testing.T) {
	exist := repo.Company.Exists(context.Background(), "id", "60cde332-ad5a-4aab-932b-81b5f16a61d2")
	t.Log(exist)
}
