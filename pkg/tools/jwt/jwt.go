package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"pms.pkg/protobuf/dto"
)

func (j *Config) GenerateAccessToken(session_id string, userData *dto.User) (string, error) {
	claims := &AccessTokenClaims{
		Email:     userData.Email,
		UserID:    userData.Id,
		SessionID: session_id,
	}
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.TTL))),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    claims.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *Config) DecodeToken(token string) (*AccessTokenClaims, error) {
	decoded, err := jwt.ParseWithClaims(token, &AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := decoded.Claims.(*AccessTokenClaims)
	if !ok {
		return nil, errors.New("invalid token or claims")
	}
	return claims, nil
}
