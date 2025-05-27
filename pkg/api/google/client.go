package google

import "go.uber.org/zap"

type Client struct {
	Conf        Config
	accessToken string

	log *zap.SugaredLogger
}

func New(conf Config, log *zap.SugaredLogger) *Client {
	return &Client{
		Conf: conf,
		log:  log,
	}
}

func (c *Client) headers() []string {
	return []string{
		"Authorization", "Bearer " + c.accessToken,
		"Accept", "application/json",
	}
}

// Default headers:
//   - Content-Type: application/json
//   - Authorization: Bearer <access-token>
//   - Accept: application/json
func (c *Client) setHeaders(headers ...string) map[string]string {
	h := make(map[string]string, 0)
	if c.accessToken != "" {
		h["Authorization"] = "Bearer " + c.accessToken
	}
	h["Content-Type"] = "application/json"
	h["Accept"] = "application/json"
	for i := 0; i < len(headers); i += 2 {
		if i+2 > len(headers) {
			continue
		}
		h[headers[i]] = headers[i+1]
	}
	return h
}
