package models

import (
	"time"

	"pms.pkg/api/google"
	"pms.pkg/consts"
	"pms.pkg/datastore/redis"
	ctxutils "pms.pkg/utils/ctx"
)

var _ redis.Cachable = &Session{}
var _ ctxutils.ContextKeyHolder = &Session{}

type Session struct {
	ID                string                         `json:"session_id"`
	UserID            string                         `json:"user_id"`
	AccessToken       string                         `json:"access_token"`
	RefreshToken      string                         `json:"refresh_token"`
	Expires           time.Time                      `json:"exp"`
	SelectedCompanyID string                         `json:"selected_company_id"`
	Companies         []string                       `json:"companies"`
	Projects          []string                       `json:"projects"`
	Permissions       map[string][]consts.Permission `json:"permissions"`
	// OAuth2 specific fields
	OAuth2 *google.Session `json:"oauth2,omitempty"`
}

func (s Session) GetDB() int {
	return 0
}

func (s Session) ContextKey() ctxutils.ContextKey {
	return ctxutils.ContextKey("session")
}
