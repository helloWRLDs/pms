package google

import "time"

type Session struct {
	Provider     string            `json:"provider"`      // Always "google"
	State        string            `json:"state"`         // CSRF state token
	RawProfile   map[string]string `json:"raw_profile"`   // Raw profile data from Google
	AccessToken  string            `json:"access_token"`  // Google's access token
	RefreshToken string            `json:"refresh_token"` // Google's refresh token
	ExpiresAt    time.Time         `json:"expires_at"`    // Token expiration time
}

func (s *Session) IsValid() bool {
	return s != nil && time.Now().Before(s.ExpiresAt)
}

func (s *Session) IsExpired() bool {
	return s == nil || time.Now().After(s.ExpiresAt)
}
