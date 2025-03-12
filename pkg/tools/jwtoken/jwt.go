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

func DecodeToken[T jwt.Claims](token string, claims T, conf *Config) (T, error) {
	decoded, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(conf.Secret), nil
	})

	if err != nil {
		return claims, err
	}

	if claims, ok := decoded.Claims.(T); ok && decoded.Valid {
		return claims, nil
	}

	return claims, errors.New("invalid token or claims")
}
