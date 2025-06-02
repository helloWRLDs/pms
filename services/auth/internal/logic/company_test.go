package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.pkg/transport/grpc/dto"
)

func TestGetCompany(t *testing.T) {
	tests := []struct {
		name      string
		companyID string
		wantErr   bool
	}{
		{
			name:      "get existing company",
			companyID: "60cde332-ad5a-4aab-932b-81b5f16a61d2",
			wantErr:   false,
		},
		{
			name:      "get non-existent company",
			companyID: "non-existent-id",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			company, err := logic.GetCompany(context.Background(), tt.companyID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, company)
			assert.Equal(t, tt.companyID, company.Id)
		})
	}
}

func TestUpdateCompany(t *testing.T) {
	companyID := "60cde332-ad5a-4aab-932b-81b5f16a61d2"

	initialCompany, err := logic.GetCompany(context.Background(), companyID)
	assert.NoError(t, err)

	tests := []struct {
		name    string
		updates *dto.Company
		wantErr bool
	}{
		{
			name: "update company name",
			updates: &dto.Company{
				Id:          companyID,
				Name:        "Updated Company",
				Codename:    initialCompany.Codename,
				Bin:         initialCompany.Bin,
				Address:     initialCompany.Address,
				Description: initialCompany.Description,
			},
			wantErr: false,
		},
		{
			name: "update company description",
			updates: &dto.Company{
				Id:          companyID,
				Name:        initialCompany.Name,
				Codename:    initialCompany.Codename,
				Bin:         initialCompany.Bin,
				Address:     initialCompany.Address,
				Description: "Updated description",
			},
			wantErr: false,
		},
		{
			name: "update non-existent company",
			updates: &dto.Company{
				Id:          "non-existent-id",
				Name:        "Test Company",
				Codename:    "TEST",
				Bin:         "123456789012",
				Address:     "Test Address",
				Description: "Test Description",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logic.UpdateCompany(context.Background(), tt.updates.Id, tt.updates)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			updated, err := logic.GetCompany(context.Background(), tt.updates.Id)
			assert.NoError(t, err)
			assert.Equal(t, tt.updates.Name, updated.Name)
			assert.Equal(t, tt.updates.Description, updated.Description)
		})
	}
}
