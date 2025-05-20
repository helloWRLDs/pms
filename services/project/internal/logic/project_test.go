package logic

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

func Test_CreateProject(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newProject := &dto.ProjectCreation{
		Title:       "Frontend",
		Description: "Developing frontend using React",
		CompanyId:   "1",
	}

	err := logic.CreateProject(ctx, newProject)
	assert.NoError(t, err)
}

func Test_ListProjects(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	list, err := logic.ListProjects(ctx, &dto.ProjectFilter{
		Page:      1,
		PerPage:   10,
		CompanyId: "60cde332-ad5a-4aab-932b-81b5f16a61d2",
	})
	assert.NoError(t, err)
	t.Log(utils.JSON(list))
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
