package cachemodels

import (
	"time"
)

var _ Cachable = &Session{}

type Session struct {
	ID     string    `json:"session_id"`
	UserID string    `json:"user_id"`
	Token  string    `json:"access_token"`
	Exp    time.Time `json:"exp"`

	db int
}

func NewSession(id, userID, token string, exp time.Time) Session {
	return Session{
		ID:     id,
		UserID: userID,
		Token:  token,
		Exp:    exp,
		db:     0,
	}
}

func (s Session) GetDB() int {
	return s.db
}
