package data

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	participantdata "pms.auth/internal/data/participant"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

func Test_CreateParticipant(t *testing.T) {
	p := participantdata.Participant{
		UserID:    "fb11170c-8f61-4fe5-858f-a5b256f6c1bd",
		CompanyID: "8f557202-0853-4672-aafb-a0b6cae7067a",
		Role:      "admin",
	}

	err := repo.Participant.Create(context.Background(), p)
	assert.NoError(t, err)
}

func Test_ListParticipant(t *testing.T) {
	list, err := repo.Participant.List(context.Background(), &dto.ParticipantFilter{
		Page:    1,
		PerPage: 10,
	})
	assert.NoError(t, err)
	t.Log(utils.JSON(list))
}

func Test_GetByUserID(t *testing.T) {
	userID := "eb306dc5-52bb-4009-88af-347b4d040718"
	p, err := repo.Participant.GetByUserID(context.Background(), userID)
	assert.NoError(t, err)
	t.Log(utils.JSON(p))
}
