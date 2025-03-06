package jwt

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"pms.pkg/protobuf/dto"
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
	user := &dto.User{
		Id:    "1",
		Name:  "Bob",
		Email: "bob@gmail.com",
	}
	token, err := Conf.GenerateAccessToken(uuid.New().String(), user)
	assert.NoError(t, err)
	t.Log(token)
	decoded, err := Conf.DecodeToken(token)
	assert.NoError(t, err)
	assert.NotNil(t, decoded)
	t.Log(utils.JSON(decoded))
}

func Test_WithOptions(t *testing.T) {
	user := &dto.User{
		Id:    "1",
		Name:  "Bob",
		Email: "bob@gmail.com",
	}
	token, err := WithConfig(
		WithTTL(24),
		WithSecret("secret")).
		GenerateAccessToken(uuid.NewString(), user)
	assert.NoError(t, err)
	t.Log("access token: ", token)
}
