package github

import (
	"os"
	"testing"
)

var (
	client = Client{
		Conf: Config{
			ClientID:     "",
			ClientSecret: "",
			RedirectURL:  "",
			HOST:         "https://api.github.co",
			Scopes:       []string{"user"},
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
