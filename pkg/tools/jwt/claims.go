package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	Email     string `json:"email"`
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	jwt.RegisteredClaims
}

func (c *AccessTokenClaims) Expired() bool {
	return c.ExpiresAt.Time.Unix() < time.Now().Unix()
}
