package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.pkg/transport/grpc/dto"
)

func TestGetProfile(t *testing.T) {
	tests := []struct {
		name    string
		userID  string
		wantErr bool
	}{
		{
			name:    "get existing user profile",
			userID:  "be10a73c-0927-4e3d-afe5-b4bae2e84946",
			wantErr: false,
		},
		{
			name:    "get non-existent user profile",
			userID:  "non-existent-id",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			profile, err := logic.GetProfile(context.Background(), tt.userID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, profile)
			assert.Equal(t, tt.userID, profile.Id)
		})
	}
}

func TestListUsers(t *testing.T) {
	tests := []struct {
		name      string
		filter    *dto.UserFilter
		wantCount int
		wantErr   bool
	}{
		{
			name: "list all users",
			filter: &dto.UserFilter{
				Page:    1,
				PerPage: 10,
			},
			wantCount: 10,
			wantErr:   false,
		},
		{
			name: "list users with company filter",
			filter: &dto.UserFilter{
				Page:      1,
				PerPage:   10,
				CompanyId: "60cde332-ad5a-4aab-932b-81b5f16a61d2",
			},
			wantCount: 5,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users, err := logic.ListUsers(context.Background(), tt.filter)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Len(t, users.Items, tt.wantCount)
			assert.Equal(t, tt.filter.Page, users.Page)
			assert.Equal(t, tt.filter.PerPage, users.PerPage)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	userID := "be10a73c-0927-4e3d-afe5-b4bae2e84946"

	initialUser, err := logic.GetProfile(context.Background(), userID)
	assert.NoError(t, err)

	tests := []struct {
		name    string
		updates *dto.User
		wantErr bool
	}{
		{
			name: "update first name",
			updates: &dto.User{
				Id:        userID,
				FirstName: "Updated",
				LastName:  initialUser.LastName,
				Email:     initialUser.Email,
			},
			wantErr: false,
		},
		{
			name: "update last name",
			updates: &dto.User{
				Id:        userID,
				FirstName: initialUser.FirstName,
				LastName:  "Updated",
				Email:     initialUser.Email,
			},
			wantErr: false,
		},
		{
			name: "update email",
			updates: &dto.User{
				Id:        userID,
				FirstName: initialUser.FirstName,
				LastName:  initialUser.LastName,
				Email:     "updated@example.com",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updated, err := logic.UpdateUser(context.Background(), userID, tt.updates)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.updates.FirstName, updated.FirstName)
			assert.Equal(t, tt.updates.LastName, updated.LastName)
			assert.Equal(t, tt.updates.Email, updated.Email)
		})
	}
}
