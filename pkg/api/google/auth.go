package google

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"go.uber.org/zap"
	"pms.pkg/tools/httpclient"
)

const (
	AuthEndpoint  = "https://accounts.google.com/o/oauth2/v2/auth"
	TokenEndpoint = "https://oauth2.googleapis.com/token"
)

func (c *Client) AuthURL(state string) string {
	scopes := strings.Join(c.Conf.Scopes, " ")
	return fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s&access_type=offline",
		AuthEndpoint,
		url.QueryEscape(c.Conf.ClientID),
		url.QueryEscape(c.Conf.RedirectURL),
		url.QueryEscape(scopes),
		url.QueryEscape(state),
	)
}

// exchange the authorization code for an access token
func (c *Client) SetToken(code string) error {
	log := c.log.With("func", "SetToken")
	log.Debug("Set token called")

	data := url.Values{}
	data.Set("client_id", c.Conf.ClientID)
	data.Set("client_secret", c.Conf.ClientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", c.Conf.RedirectURL)
	data.Set("grant_type", "authorization_code")

	res, err := httpclient.New().
		Method("POST").
		URL(TokenEndpoint).
		Headers(
			"Content-Type", "application/x-www-form-urlencoded",
			"Accept", "application/json",
		).
		Body(data.Encode()).
		Do()

	if err != nil {
		log.Error("failed to make request", zap.Error(err))
		return err
	}

	if res.Status >= 400 {
		log.Error("Google OAuth request failed", zap.Int("status", res.Status))
		return errors.New("google OAuth request failed")
	}

	var authResponse struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		Scope        string `json:"scope"`
		ExpiresIn    int    `json:"expires_in"`
		IDToken      string `json:"id_token"`
		RefreshToken string `json:"refresh_token,omitempty"`
	}

	if err = res.ScanJSON(&authResponse); err != nil {
		log.Error("failed to unmarshal response", zap.Error(err))
		return err
	}

	c.accessToken = authResponse.AccessToken
	return nil
}
