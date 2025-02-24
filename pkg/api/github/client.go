package github

import "fmt"

type Config struct {
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
	RedirectURL  string `env:"REDICRECT_URL"`
	HOST         string `env:"HOST"`
	Scopes       []string
}

// Default headers:
//  - Content-Type: application/json
//  - Authorization: Bearer <access-token>
func (c *Client) setHeaders(headers ...string) map[string]string {
	h := make(map[string]string, 0)
	if c.accessToken != "" {
		h["Authorization"] = fmt.Sprintf("Bearer %s", c.accessToken)
	}
	h["Content-Type"] = "application/json"
	for i := 0; i < len(headers); i += 2 {
		if i+2 > len(headers) {
			continue
		}
		h[headers[i]] = headers[i+1]
	}
	return h
}

type Client struct {
	Conf        Config
	accessToken string
}

func New(conf Config) *Client {
	return &Client{
		Conf: conf,
	}
}
