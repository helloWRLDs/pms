package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddParticipant(t *testing.T) {
	tests := []struct {
		name      string
		companyID string
		userID    string
		roleID    string
		wantErr   bool
	}{
		{
			name:      "add valid participant",
			companyID: "60cde332-ad5a-4aab-932b-81b5f16a61d2",
			userID:    "eb306dc5-52bb-4009-88af-347b4d040718",
			roleID:    "admin",
			wantErr:   false,
		},
		{
			name:      "add participant with invalid user",
			companyID: "60cde332-ad5a-4aab-932b-81b5f16a61d2",
			userID:    "non-existent-user",
			roleID:    "1",
			wantErr:   true,
		},
		{
			name:      "add participant with invalid company",
			companyID: "non-existent-company",
			userID:    "eb306dc5-52bb-4009-88af-347b4d040718",
			roleID:    "1",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := logic.AddParticipant(context.Background(), tt.companyID, tt.userID, tt.roleID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			company, err := logic.GetCompany(context.Background(), tt.companyID)
			assert.NoError(t, err)
			assert.NotNil(t, company)
			assert.Contains(t, company, tt.userID)
		})
	}
}

func TestDeleteParticipant(t *testing.T) {
	tests := []struct {
		name      string
		companyID string
		userID    string
		wantErr   bool
	}{
		{
			name:      "delete existing participant",
			companyID: "60cde332-ad5a-4aab-932b-81b5f16a61d2",
			userID:    "eb306dc5-52bb-4009-88af-347b4d040718",
			wantErr:   false,
		},
		{
			name:      "delete non-existent participant",
			companyID: "60cde332-ad5a-4aab-932b-81b5f16a61d2",
			userID:    "non-existent-user",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logic.DeleteParticipant(context.Background(), tt.companyID, tt.userID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			company, err := logic.GetCompany(context.Background(), tt.companyID)
			assert.NoError(t, err)
			assert.NotNil(t, company)
			assert.NotContains(t, company, tt.userID)
		})
	}
}
