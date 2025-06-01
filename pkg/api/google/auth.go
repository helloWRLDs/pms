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
	log := c.log.With("func", "SetToken").With(
		zap.String("code", code),
		zap.String("client_id", c.Conf.ClientID),
		zap.String("redirect_uri", c.Conf.RedirectURL),
		zap.String("scopes", strings.Join(c.Conf.Scopes, " ")),
		zap.String("token_endpoint", TokenEndpoint),
		zap.String("client_secret", c.Conf.ClientSecret),
	)
	log.Debug("Set token called")

	res, err := httpclient.New().
		Method("POST").
		URL(TokenEndpoint).
		Headers(
			"Content-Type", "application/x-www-form-urlencoded",
			"Accept", "application/json",
		).
		Query(
			"grant_type", "authorization_code",
			"client_id", c.Conf.ClientID,
			"client_secret", c.Conf.ClientSecret,
			"code", code,
			"redirect_uri", c.Conf.RedirectURL,
		).
		Do()

	if err != nil {
		log.Error("failed to make request", zap.Error(err))
		return err
	}

	if res.Status >= 400 {
		log.Errorw("Google OAuth request failed",
			"status", res.Status,
			"response_body", string(res.Data),
		)
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
