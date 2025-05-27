package github

import (
	"fmt"

	"go.uber.org/zap"
)

type Client struct {
	conf        Config
	accessToken string

	log *zap.SugaredLogger
}

func New(conf Config, log *zap.SugaredLogger) *Client {
	return &Client{
		conf: conf,
		log:  log,
	}
}

func (c *Client) headers() []string {
	return []string{
		"Authorization", fmt.Sprintf("Bearer %s", c.accessToken),
	}
}

// Default headers:
//   - Content-Type: application/json
//   - Authorization: Bearer <access-token>
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
