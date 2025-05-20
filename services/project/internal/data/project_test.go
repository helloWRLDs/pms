package data

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
	projectdata "pms.project/internal/data/project"
)

func Test_CreateProject(t *testing.T) {
	project := projectdata.Project{
		ID:          uuid.NewString(),
		Title:       "Frontend",
		Description: "Writing frontend on React",
		Status:      consts.PROJECT_STATUS_ACTIVE,
		CompanyID:   "60cde332-ad5a-4aab-932b-81b5f16a61d2",
	}

	err := repo.Project.Create(context.Background(), project)
	assert.NoError(t, err)
}

func Test_GetProject(t *testing.T) {
	projectID := "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1"
	p, err := repo.Project.GetByID(context.Background(), projectID)
	assert.NoError(t, err)
	t.Log(utils.JSON(p))
}

func Test_ListProject(t *testing.T) {
	l, err := repo.Project.List(context.Background(), &dto.ProjectFilter{
		Page:      1,
		PerPage:   10,
		CompanyId: "60cde332-ad5a-4aab-932b-81b5f16a61d2",
	})
	assert.NoError(t, err)
	t.Log(utils.JSON(l))
}

func Test_GetProjectCode(t *testing.T) {
	projectID := "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1"
	code := repo.Project.GetCode(context.Background(), projectID)
	t.Log(code)
}
