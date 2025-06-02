package logic

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
)

func TestCreateSprint(t *testing.T) {
	tests := []struct {
		name     string
		creation *dto.SprintCreation
		wantErr  bool
	}{
		{
			name: "create valid sprint",
			creation: &dto.SprintCreation{
				Title:       "Sprint 1",
				Description: "First sprint of the project",
				ProjectId:   "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
				StartDate:   timestamppb.New(time.Now()),
				EndDate:     timestamppb.New(time.Now().Add(14 * 24 * time.Hour)),
			},
			wantErr: false,
		},
		{
			name: "create sprint with empty title",
			creation: &dto.SprintCreation{
				Title:       "",
				Description: "First sprint of the project",
				ProjectId:   "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
				StartDate:   timestamppb.New(time.Now()),
				EndDate:     timestamppb.New(time.Now().Add(14 * 24 * time.Hour)),
			},
			wantErr: true,
		},
		{
			name: "create sprint with invalid project",
			creation: &dto.SprintCreation{
				Title:       "Sprint 1",
				Description: "First sprint of the project",
				ProjectId:   "invalid-project-id",
				StartDate:   timestamppb.New(time.Now()),
				EndDate:     timestamppb.New(time.Now().Add(14 * 24 * time.Hour)),
			},
			wantErr: true,
		},
		{
			name: "create sprint with end date before start date",
			creation: &dto.SprintCreation{
				Title:       "Sprint 1",
				Description: "First sprint of the project",
				ProjectId:   "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
				StartDate:   timestamppb.New(time.Now().Add(14 * 24 * time.Hour)),
				EndDate:     timestamppb.New(time.Now()),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sprint, err := logic.CreateSprint(context.Background(), tt.creation)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, sprint)
			assert.Equal(t, tt.creation.Title, sprint.Title)
			assert.Equal(t, tt.creation.Description, sprint.Description)
			assert.Equal(t, tt.creation.ProjectId, sprint.ProjectId)
		})
	}
}

func TestListSprints(t *testing.T) {
	tests := []struct {
		name      string
		filter    *dto.SprintFilter
		wantCount int
		wantErr   bool
	}{
		{
			name: "list all sprints",
			filter: &dto.SprintFilter{
				Page:    1,
				PerPage: 10,
			},
			wantCount: 10,
			wantErr:   false,
		},
		{
			name: "list sprints by project",
			filter: &dto.SprintFilter{
				Page:      1,
				PerPage:   10,
				ProjectId: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
			},
			wantCount: 5,
			wantErr:   false,
		},
		{
			name: "list sprints with invalid project",
			filter: &dto.SprintFilter{
				Page:      1,
				PerPage:   10,
				ProjectId: "invalid-project-id",
			},
			wantCount: 0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sprints, err := logic.ListSprints(context.Background(), tt.filter)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Len(t, sprints.Items, tt.wantCount)
			assert.Equal(t, tt.filter.Page, sprints.Page)
			assert.Equal(t, tt.filter.PerPage, sprints.PerPage)
		})
	}
}

func TestGetSprint(t *testing.T) {
	tests := []struct {
		name     string
		sprintID string
		wantErr  bool
	}{
		{
			name:     "get existing sprint",
			sprintID: "712b8a41-2351-4286-ad03-086eaee4c417",
			wantErr:  false,
		},
		{
			name:     "get non-existent sprint",
			sprintID: "non-existent-id",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sprint, err := logic.GetSprint(context.Background(), tt.sprintID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, sprint)
			assert.Equal(t, tt.sprintID, sprint.Id)
		})
	}
}

func TestUpdateSprint(t *testing.T) {
	tests := []struct {
		name     string
		sprintID string
		updates  *dto.Sprint
		wantErr  bool
	}{
		{
			name:     "update sprint title",
			sprintID: "712b8a41-2351-4286-ad03-086eaee4c417",
			updates: &dto.Sprint{
				Id:          "712b8a41-2351-4286-ad03-086eaee4c417",
				Title:       "Updated Sprint Title",
				Description: "Original Description",
				ProjectId:   "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
				StartDate:   timestamppb.New(time.Now()),
				EndDate:     timestamppb.New(time.Now().Add(14 * 24 * time.Hour)),
			},
			wantErr: false,
		},
		{
			name:     "update non-existent sprint",
			sprintID: "non-existent-id",
			updates: &dto.Sprint{
				Id:          "non-existent-id",
				Title:       "Updated Sprint Title",
				Description: "Original Description",
				ProjectId:   "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
				StartDate:   timestamppb.New(time.Now()),
				EndDate:     timestamppb.New(time.Now().Add(14 * 24 * time.Hour)),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updated, err := logic.UpdateSprint(context.Background(), tt.sprintID, tt.updates)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, updated)
			assert.Equal(t, tt.updates.Title, updated.Title)
			assert.Equal(t, tt.updates.Description, updated.Description)
		})
	}
}

func TestListSprintsWithFilters(t *testing.T) {

	tests := []struct {
		name      string
		filter    *dto.SprintFilter
		wantCount int
		wantErr   bool
	}{
		{
			name: "list sprints by date range",
			filter: &dto.SprintFilter{
				Page:      1,
				PerPage:   10,
				ProjectId: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
			},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "list sprints by status",
			filter: &dto.SprintFilter{
				Page:      1,
				PerPage:   10,
				ProjectId: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
			},
			wantCount: 1,
			wantErr:   false,
		},
		{
			name: "list sprints with invalid date range",
			filter: &dto.SprintFilter{
				Page:      1,
				PerPage:   10,
				ProjectId: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
			},
			wantCount: 0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sprints, err := logic.ListSprints(context.Background(), tt.filter)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Len(t, sprints.Items, tt.wantCount)
			assert.Equal(t, tt.filter.Page, sprints.Page)
			assert.Equal(t, tt.filter.PerPage, sprints.PerPage)
		})
	}
}
