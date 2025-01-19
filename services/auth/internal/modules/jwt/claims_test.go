package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Test_Expired(t *testing.T) {
	claims := AccessTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}
	t.Log(claims.GetExpirationTime())
}
