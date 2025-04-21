package data

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"pms.auth/internal/entity"
	"pms.pkg/utils"
)

func Test_CreateParticipant(t *testing.T) {
	p := entity.Participant{
		UserId:    "fb11170c-8f61-4fe5-858f-a5b256f6c1bd",
		CompanyId: "8f557202-0853-4672-aafb-a0b6cae7067a",
		RoleId:    "admin",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := repo.Participant.Create(ctx, p)
	assert.NoError(t, err)
}

func Test_GetByUserID(t *testing.T) {
	userID := "be10a73c-0927-4e3d-afe5-b4bae2e84946"
	p, err := repo.Participant.GetByUserID(context.Background(), userID)
	assert.NoError(t, err)
	t.Log(utils.JSON(p))
}
