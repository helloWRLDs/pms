package github

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"go.uber.org/zap"
	"pms.pkg/tools/httpclient"
)

func (c *Client) AuthURL(state string) string {
	scopes := strings.Join(c.conf.Scopes, " ")
	return fmt.Sprintf(
		"%s/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=%s&state=%s",
		c.conf.HOST,
		url.QueryEscape(c.conf.ClientID),
		url.QueryEscape(c.conf.RedirectURL),
		url.QueryEscape(scopes),
		url.QueryEscape(state),
	)
}

func (c *Client) SetToken(code string) error {
	log := c.log.With("func", "SetToken")

	data := url.Values{}

	data.Set("client_id", c.conf.ClientID)
	data.Set("client_secret", c.conf.ClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", c.conf.RedirectURL)

	res, err := httpclient.New().
		Method("POST").
		URL("https://github.com/login/oauth/access_token").
		Body(data.Encode()).
		Headers(c.headers()...).
		Do()

	if err != nil {
		log.Error("failed to make request", zap.Error(err))
		return err
	}

	if res.Status >= 400 {
		log.Error("GitHub OAuth request failed", zap.Int("status", res.Status))
		return errors.New("GitHub OAuth request failed")
	}

	var authResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	if err = res.ScanJSON(&authResponse); err != nil {
		return err
	}
	c.accessToken = authResponse.AccessToken
	return nil
}
