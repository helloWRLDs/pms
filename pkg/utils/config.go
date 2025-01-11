package utils

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Config parser
func LoadConfig[T any](path string) (T, error) {
	var t T
	err := godotenv.Load(path)
	if err != nil {
		return t, err
	}
	err = env.Parse(&t)
	if err != nil {
		return t, err
	}
	return t, nil
}
