package github

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"pms.pkg/tools/httpclient"
)

func (c *Client) AuthURL(state string) string {
	scopes := strings.Join(c.Conf.Scopes, " ")
	return fmt.Sprintf(
		"%s/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=%s&state=%s",
		c.Conf.HOST,
		url.QueryEscape(c.Conf.ClientID),
		url.QueryEscape(c.Conf.RedirectURL),
		url.QueryEscape(scopes),
		url.QueryEscape(state),
	)
}

// exchange the authorization code for an access token
func (c *Client) SetToken(code string) error {
	log := logrus.WithField("func", "SetToken")

	data := url.Values{}

	data.Set("client_id", c.Conf.ClientID)
	data.Set("client_secret", c.Conf.ClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", c.Conf.RedirectURL)

	res, err := httpclient.New().
		Method("POST").
		URL("https://github.com/login/oauth/access_token").
		Body(data.Encode()).
		Headers(c.headers()...).
		Do()

	if err != nil {
		log.WithError(err).Error("failed to make request")
		return err
	}

	if res.Status >= 400 {
		log.WithField("res", res.Status).Error("GitHub OAuth request failed")
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
