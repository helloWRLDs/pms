package logic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.pkg/utils"
)

func Test_GetCompany(t *testing.T) {
	id := "993e92af-e3ed-4b9f-9d51-c3ea30c40d08"
	comp, err := logic.GetCompany(context.Background(), id)
	assert.NoError(t, err)
	t.Log(utils.JSON(comp))
}
