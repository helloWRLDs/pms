package models

import (
	"time"

	"pms.api-gateway/internal/modules/cache"
	"pms.pkg/utils/ctx"
)

var _ cache.Cachable = &Session{}
var _ ctx.ContextKeyHolder = &Session{}

type Session struct {
	ID           string    `json:"session_id"`
	UserID       string    `json:"user_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Exp          time.Time `json:"exp"`
}

func (s Session) GetDB() int {
	return 0
}

func (s Session) ContextKey() ctx.ContextKey {
	return ctx.ContextKey("session")
}
