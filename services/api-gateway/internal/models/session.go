package models

import (
	"time"

	"pms.pkg/datastore/redis"
	ctxutils "pms.pkg/utils/ctx"
)

var _ redis.Cachable = &Session{}
var _ ctxutils.ContextKeyHolder = &Session{}

type Session struct {
	ID                string    `json:"session_id"`
	UserID            string    `json:"user_id"`
	AccessToken       string    `json:"access_token"`
	RefreshToken      string    `json:"refresh_token"`
	Expires           time.Time `json:"exp"`
	SelectedCompanyID string    `json:"selected_company_id"`
}

func (s Session) GetDB() int {
	return 0
}

func (s Session) ContextKey() ctxutils.ContextKey {
	return ctxutils.ContextKey("session")
}
