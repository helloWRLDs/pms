package companydata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.pkg/type/list"
	"pms.pkg/utils"

	"pms.auth/internal/entity"
)

func Test_CreateCompany(t *testing.T) {
	company := entity.Company{
		Name:     "Tech Test 3",
		Codename: "TT123",
	}

	err := repo.Create(context.Background(), company)
	assert.NoError(t, err, "expected no error when creating company")
	t.Log("Company created successfully")
}

func Test_Exist(t *testing.T) {
	t.Log("Company exists: ", repo.Exists(context.Background(), "8f557202-0853-4672-aafb-a0b6cae7067a"))
}

func Test_ListComapanies(t *testing.T) {
	filter := list.Filters{
		Pagination: list.Pagination{
			Page:    1,
			PerPage: 10,
		},
		Fields: map[string]string{
			"name": "Tech Test 3",
		},
	}
	companies, err := repo.List(context.Background(), filter)
	assert.NoError(t, err, "expected no error when listing companies")
	t.Log("Companies list:", utils.JSON(companies))
}
