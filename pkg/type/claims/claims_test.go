package claims

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestAccessTokenClaims_Expired(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		claims   AccessTokenClaims
		expected bool
	}{
		{
			name: "expired token",
			claims: AccessTokenClaims{
				Email:     "test@example.com",
				UserID:    "123",
				SessionID: "session123",
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(now.Add(-1 * time.Hour)),
				},
			},
			expected: true,
		},
		{
			name: "valid token",
			claims: AccessTokenClaims{
				Email:     "test@example.com",
				UserID:    "123",
				SessionID: "session123",
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(now.Add(1 * time.Hour)),
				},
			},
			expected: false,
		},
		{
			name: "token expiring now",
			claims: AccessTokenClaims{
				Email:     "test@example.com",
				UserID:    "123",
				SessionID: "session123",
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(now),
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.claims.Expired(); got != tt.expected {
				t.Errorf("AccessTokenClaims.Expired() = %v, want %v", got, tt.expected)
			}
		})
	}
}
