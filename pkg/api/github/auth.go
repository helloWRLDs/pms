package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"pms.pkg/utils/request"
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

	details := request.Details{
		Method:  http.MethodPost,
		URL:     "https://github.com/login/oauth/access_token",
		Headers: c.setHeaders(),
		Body:    []byte(data.Encode()),
	}

	status, res, err := details.Make()
	if err != nil {
		log.WithError(err).Error("failed to make request")
		return err
	}

	if status >= 400 {
		log.WithField("res", string(res)).Error("GitHub OAuth request failed")
		return errors.New("GitHub OAuth request failed")
	}

	var authResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	if err = json.Unmarshal(res, &authResponse); err != nil {
		log.WithError(err).Error("failed to parse GitHub token response")
		return err
	}

	c.accessToken = authResponse.AccessToken

	return nil
}
