package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.pkg/utils"
)

func Test_getUserData(t *testing.T) {
	user, err := client.GetUserData()
	if assert.NoError(t, err) {
		t.Log(utils.JSON(user))
	}
}
