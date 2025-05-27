package google

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
	"pms.pkg/tools/httpclient"
)

const (
	UserInfoEndpoint = "https://www.googleapis.com/oauth2/v2/userinfo"
)

type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func (c *Client) GetUserData() (user User, err error) {
	log := c.log.With("func", "GetUserData")
	log.Debug("GetUserData called")

	res, err := httpclient.New().
		Method("GET").
		URL(UserInfoEndpoint).
		Headers(c.headers()...).
		Do()

	if err != nil {
		log.Error("failed to make request", zap.Error(err))
		return user, err
	}

	if res.Status >= 400 {
		log.Error("failed to get user data", zap.Int("status", res.Status))
		return user, errors.New("failed to get user data")
	}

	if err = res.ScanJSON(&user); err != nil {
		log.Error("failed to unmarshal response", zap.Error(err))
		return user, fmt.Errorf("failed to parse user data: %v", err)
	}

	return user, nil
}
