package google

import "time"

// Session represents a Google OAuth2 session
type Session struct {
	Provider     string            `json:"provider"`      // Always "google"
	State        string            `json:"state"`         // CSRF state token
	RawProfile   map[string]string `json:"raw_profile"`   // Raw profile data from Google
	AccessToken  string            `json:"access_token"`  // Google's access token
	RefreshToken string            `json:"refresh_token"` // Google's refresh token
	ExpiresAt    time.Time         `json:"expires_at"`    // Token expiration time
}

// IsValid checks if the session is still valid
func (s *Session) IsValid() bool {
	return s != nil && time.Now().Before(s.ExpiresAt)
}

// IsExpired checks if the session has expired
func (s *Session) IsExpired() bool {
	return s == nil || time.Now().After(s.ExpiresAt)
}
