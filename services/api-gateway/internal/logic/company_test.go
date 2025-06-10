package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"pms.api-gateway/internal/config"
	"pms.pkg/logger"
	configgrpc "pms.pkg/transport/grpc/config"
	"pms.pkg/transport/grpc/dto"
)

func TestIntegration_Company(t *testing.T) {

	cfg := config.Config{
		Auth: configgrpc.ClientConfig{
			Host: "localhost:50051",
		},
		Project: configgrpc.ClientConfig{
			Host: "localhost:50052",
		},
	}
	log := logger.Log

	l := &Logic{
		Config: cfg,
		log:    log,
	}

	t.Run("List Companies", func(t *testing.T) {
		filter := &dto.CompanyFilter{
			Page:    1,
			PerPage: 10,
		}
		companies, err := l.ListCompanies(context.Background(), filter)
		require.NoError(t, err)
		assert.NotNil(t, companies)
		assert.GreaterOrEqual(t, companies.TotalItems, int32(0))
		assert.GreaterOrEqual(t, companies.TotalPages, int32(0))

		for _, company := range companies.Items {
			assert.NotEmpty(t, company.Id)
			assert.NotEmpty(t, company.Name)
			assert.NotEmpty(t, company.Codename)
			assert.NotNil(t, company.Projects)
		}
	})

	t.Run("Get Company", func(t *testing.T) {
		creation := &dto.NewCompany{
			Name:        "Test Company",
			Codename:    "TEST",
			Bin:         "123456789",
			Address:     "Test Address",
			Description: "Test Description",
		}
		err := l.CreateCompany(context.Background(), creation)
		require.NoError(t, err)

		company, err := l.GetCompany(context.Background(), "test-company-1")
		require.NoError(t, err)
		assert.NotNil(t, company)
		assert.Equal(t, "Test Company", company.Name)
		assert.Equal(t, "TEST", company.Codename)
		assert.Equal(t, "123456789", company.Bin)
		assert.Equal(t, "Test Address", company.Address)
		assert.Equal(t, "Test Description", company.Description)
		assert.NotNil(t, company.Projects)
	})

	t.Run("Create Company", func(t *testing.T) {
		creation := &dto.NewCompany{
			Name:        "New Test Company",
			Codename:    "NEWTEST",
			Bin:         "987654321",
			Address:     "New Test Address",
			Description: "New Test Description",
		}
		err := l.CreateCompany(context.Background(), creation)
		require.NoError(t, err)
	})

	t.Run("Add and Remove Participant", func(t *testing.T) {
		companyID := "test-company-1" // Replace with actual company ID
		userID := "test-user-1"       // Replace with actual user ID

		err := l.CompanyAddParticipant(context.Background(), companyID, userID, "admin")
		require.NoError(t, err)

		err = l.CompanyRemoveParticipant(context.Background(), companyID, userID)
		require.NoError(t, err)
	})

	t.Run("Get Company with Invalid ID", func(t *testing.T) {
		company, err := l.GetCompany(context.Background(), "non-existent-company")
		assert.Error(t, err)
		assert.Nil(t, company)
	})

	t.Run("Create Company with Invalid Data", func(t *testing.T) {
		creation := &dto.NewCompany{
			Name: "", // Empty name should fail
		}
		err := l.CreateCompany(context.Background(), creation)
		assert.Error(t, err)
	})

	t.Run("Add Participant with Invalid IDs", func(t *testing.T) {
		err := l.CompanyAddParticipant(context.Background(), "invalid-company", "invalid-user", "admin")
		assert.Error(t, err)
	})

	t.Run("Remove Participant with Invalid IDs", func(t *testing.T) {
		err := l.CompanyRemoveParticipant(context.Background(), "invalid-company", "invalid-user")
		assert.Error(t, err)
	})
}
