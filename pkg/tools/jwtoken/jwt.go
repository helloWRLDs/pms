package jwtoken

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken[T jwt.Claims](claims T, conf *Config) (string, error) {
	if _, ok := any(claims).(jwt.Claims); !ok {
		return "", errors.New("invalid claims type")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func DecodeToken(token string, claims jwt.Claims, conf *Config) (jwt.Claims, error) {
	decoded, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(conf.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if decoded.Valid {
		return decoded.Claims, nil
	}
	return nil, errors.New("invalid token or claims")
}
