package data

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
	sprintdata "pms.project/internal/data/sprint"
)

func Test_CreateSprint(t *testing.T) {
	sprint := sprintdata.Sprint{
		ID:          uuid.NewString(),
		Title:       "Sprint 1",
		Description: "Sprint 1 description",
		StartDate:   time.Now(),
		EndDate:     time.Now().AddDate(0, 1, 0),
		ProjectID:   "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
	}
	err := repo.Sprint.Create(context.Background(), sprint)
	assert.NoError(t, err)
	created, err := repo.Sprint.GetByID(context.Background(), sprint.ID)
	assert.NoError(t, err)
	t.Log(utils.JSON(created))
}

func Test_ListSprints(t *testing.T) {
	sprints, err := repo.Sprint.List(context.Background(), &dto.SprintFilter{
		ProjectId: "fc6ba40d-8d0a-48b1-a84d-e358f6438aa1",
	})
	assert.NoError(t, err)
	t.Log(utils.JSON(sprints))
}
