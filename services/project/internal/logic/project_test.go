package logic

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

func TestCreateProject(t *testing.T) {
	tests := []struct {
		name     string
		creation *dto.ProjectCreation
		wantErr  bool
	}{
		{
			name: "create valid project",
			creation: &dto.ProjectCreation{
				Title:       "Test Project",
				Description: "Test Description",
				CompanyId:   "60cde332-ad5a-4aab-932b-81b5f16a61d2",
				CodeName:    "TEST",
				CodePrefix:  "TST",
			},
			wantErr: false,
		},
		{
			name: "create project with empty title",
			creation: &dto.ProjectCreation{
				Title:       "",
				Description: "Test Description",
				CompanyId:   "60cde332-ad5a-4aab-932b-81b5f16a61d2",
				CodeName:    "TEST",
				CodePrefix:  "TST",
			},
			wantErr: true,
		},
		{
			name: "create project with invalid company",
			creation: &dto.ProjectCreation{
				Title:       "Test Project",
				Description: "Test Description",
				CompanyId:   "invalid-company-id",
				CodeName:    "TEST",
				CodePrefix:  "TST",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logic.CreateProject(context.Background(), tt.creation)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestListProjects(t *testing.T) {
	tests := []struct {
		name      string
		filter    *dto.ProjectFilter
		wantCount int
		wantErr   bool
	}{
		{
			name: "list all projects",
			filter: &dto.ProjectFilter{
				Page:    1,
				PerPage: 10,
			},
			wantCount: 10,
			wantErr:   false,
		},
		{
			name: "list projects by company",
			filter: &dto.ProjectFilter{
				Page:      1,
				PerPage:   10,
				CompanyId: "60cde332-ad5a-4aab-932b-81b5f16a61d2",
			},
			wantCount: 5,
			wantErr:   false,
		},
		{
			name: "list projects with invalid company",
			filter: &dto.ProjectFilter{
				Page:      1,
				PerPage:   10,
				CompanyId: "invalid-company-id",
			},
			wantCount: 0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			projects, err := logic.ListProjects(context.Background(), tt.filter)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Len(t, projects.Items, tt.wantCount)
			assert.Equal(t, tt.filter.Page, projects.Page)
			assert.Equal(t, tt.filter.PerPage, projects.PerPage)
		})
	}
}

func TestGetProjectByID(t *testing.T) {
	tests := []struct {
		name      string
		projectID string
		wantErr   bool
	}{
		{
			name:      "get existing project",
			projectID: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
			wantErr:   false,
		},
		{
			name:      "get non-existent project",
			projectID: "non-existent-id",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			project, err := logic.GetProjectByID(context.Background(), tt.projectID)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, project)
			assert.Equal(t, tt.projectID, project.Id)
		})
	}
}

func Test_getProject(t *testing.T) {
	projectID := "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1"
	project, err := logic.getProject(context.Background(), projectID)
	assert.NoError(t, err)
	t.Log(utils.JSON(project))
}

func Test_GetProject(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	projectID := "2fd3661d-2340-4c17-aa8e-23856ffa7412"

	proj, err := logic.GetProjectByID(ctx, projectID)
	assert.NoError(t, err)
	t.Log(utils.JSON(proj))
}
