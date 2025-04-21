package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.pkg/utils"
)

func Test_GetProfile(t *testing.T) {
	userID := "be10a73c-0927-4e3d-afe5-b4bae2e84946"
	profile, err := logic.GetProfile(context.Background(), userID)
	assert.NoError(t, err)
	t.Log(utils.JSON(profile))
}
