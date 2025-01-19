package jwt

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"pms.pkg/utils"
)

var (
	Conf Config
)

func TestMain(m *testing.M) {
	setupConfig()
	code := m.Run()
	os.Exit(code)
}

func setupConfig() {
	Conf = Config{
		TTL:    1440,
		Secret: "secret",
	}
}

func Test_GenerateToken(t *testing.T) {
	claims := AccessTokenClaims{
		Email: "danil.li24x@gmail.com",
	}
	token, err := Conf.GenerateAccessToken(claims)
	assert.NoError(t, err)
	t.Log(token)
	decoded, err := Conf.DecodeToken(token)
	assert.NoError(t, err)
	assert.NotNil(t, decoded)
	t.Log(utils.JSON(decoded))
}
