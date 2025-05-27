package google

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	client = Client{
		Conf: Config{
			ClientID:     "",
			ClientSecret: "",
			RedirectURL:  "",
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
		},
		accessToken: "token",
	}
)

func TestMain(m *testing.M) {
	checkClient()
	code := m.Run()
	os.Exit(code)
}

func checkClient() {
	if client.accessToken == "" {
		panic("access_token required")
	}
}

func TestAuthURL(t *testing.T) {
	url := client.AuthURL("test-state")
	assert.Contains(t, url, "https://accounts.google.com/o/oauth2/v2/auth")
	assert.Contains(t, url, "response_type=code")
	assert.Contains(t, url, "access_type=offline")
	assert.Contains(t, url, "state=test-state")
}
