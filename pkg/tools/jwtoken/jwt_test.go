package jwtoken

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"pms.pkg/utils"
)

// CustomClaims is a generic struct for JWT claims
type CustomClaims struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	SessionID string `json:"session_id"`
	jwt.RegisteredClaims
}

var Conf Config

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
	claims := CustomClaims{
		UserID:    uuid.NewString(),
		Email:     "bob@example.com",
		SessionID: uuid.NewString(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token, err := GenerateAccessToken(claims, &Conf)
	assert.NoError(t, err)
	t.Log(token)

	decoded, err := DecodeToken(token, &CustomClaims{}, &Conf)
	assert.NoError(t, err)
	assert.NotNil(t, decoded)
	t.Log(utils.JSON(decoded))
}
