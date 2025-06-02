package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"pms.api-gateway/internal/config"
	configgrpc "pms.pkg/transport/grpc/config"
	"pms.pkg/transport/grpc/dto"
)

func TestIntegration_Project(t *testing.T) {

	cfg := &config.Config{
		Project: configgrpc.ClientConfig{
			Host: "localhost:50051",
		},
		Auth: configgrpc.ClientConfig{
			Host: "localhost:50052",
		},
	}

	log := zap.NewNop().Sugar()

	l := New(*cfg, log)

	t.Run("List Projects", func(t *testing.T) {
		filter := &dto.ProjectFilter{
			Page:    1,
			PerPage: 10,
		}

		projects, err := l.ListProjects(context.Background(), filter)
		require.NoError(t, err)
		assert.NotNil(t, projects)
		assert.NotEmpty(t, projects.Items)
		assert.Equal(t, filter.Page, projects.Page)
		assert.Equal(t, filter.PerPage, projects.PerPage)
	})

	t.Run("Get Project", func(t *testing.T) {
		creation := &dto.ProjectCreation{
			Title:       "Test Project",
			Description: "Test Description",
			CompanyId:   "test-company-1",
			CodeName:    "TEST",
		}

		err := l.CreateProject(context.Background(), creation)
		require.NoError(t, err)

		filter := &dto.ProjectFilter{
			Page:    1,
			PerPage: 1,
		}
		projects, err := l.ListProjects(context.Background(), filter)
		require.NoError(t, err)
		require.NotEmpty(t, projects.Items)

		projectID := projects.Items[0].Id

		project, err := l.GetProjectByID(context.Background(), projectID)
		require.NoError(t, err)
		assert.NotNil(t, project)
		assert.Equal(t, projectID, project.Id)
		assert.Equal(t, creation.Title, project.Title)
		assert.Equal(t, creation.Description, project.Description)
		assert.Equal(t, creation.CompanyId, project.CompanyId)
		assert.Equal(t, creation.CodeName, project.CodeName)
	})

	t.Run("Create Project", func(t *testing.T) {
		creation := &dto.ProjectCreation{
			Title:       "New Test Project",
			Description: "New Test Description",
			CompanyId:   "test-company-1",
			CodeName:    "NEW",
		}

		err := l.CreateProject(context.Background(), creation)
		require.NoError(t, err)

		filter := &dto.ProjectFilter{
			Page:    1,
			PerPage: 10,
		}
		projects, err := l.ListProjects(context.Background(), filter)
		require.NoError(t, err)

		found := false
		for _, p := range projects.Items {
			if p.Title == creation.Title {
				found = true
				assert.Equal(t, creation.Description, p.Description)
				assert.Equal(t, creation.CompanyId, p.CompanyId)
				assert.Equal(t, creation.CodeName, p.CodeName)
				break
			}
		}
		assert.True(t, found, "Created project not found in list")
	})

	t.Run("Error Cases", func(t *testing.T) {
		_, err := l.GetProjectByID(context.Background(), "non-existent-id")
		assert.Error(t, err)

		invalidCreation := &dto.ProjectCreation{
			Title: "",
		}
		err = l.CreateProject(context.Background(), invalidCreation)
		assert.Error(t, err)

		invalidFilter := &dto.ProjectFilter{
			Page:    -1,
			PerPage: 0,
		}
		_, err = l.ListProjects(context.Background(), invalidFilter)
		assert.Error(t, err)
	})
}
