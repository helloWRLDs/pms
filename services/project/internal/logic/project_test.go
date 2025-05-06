package logic

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
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

	list, err := logic.ListProjects(ctx, list.Filters{
		Fields: map[string]string{
			"company_id": "1",
		},
	})
	assert.NoError(t, err)
	t.Log(utils.JSON(list))
}

func Test_GetProject(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	projectID := "2fd3661d-2340-4c17-aa8e-23856ffa7412"

	proj, err := logic.GetProjectByID(ctx, projectID)
	assert.NoError(t, err)
	t.Log(utils.JSON(proj))
}
