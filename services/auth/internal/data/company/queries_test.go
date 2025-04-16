package companydata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	comp "pms.auth/internal/entity/company"
)

func Test_CreateCompany(t *testing.T) {
	company := comp.Company{
		Name:     "Tech Test",
		Codename: "TT123",
	}

	err := repo.Create(context.Background(), company)
	assert.NoError(t, err, "expected no error when creating company")
	t.Log("Company created successfully")
}
