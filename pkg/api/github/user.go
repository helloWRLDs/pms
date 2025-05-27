package github

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"pms.pkg/tools/httpclient"
	"pms.pkg/utils"
)

type User struct {
	ID               int64     `json:"id"`
	Login            string    `json:"login"`
	NodeID           string    `json:"node_id"`
	AvatarURL        string    `json:"avatar_url"`
	ApiURL           string    `json:"url"`
	HtmlURL          string    `json:"html_url"`
	OrhagizationsURL string    `json:"organizations_url"`
	ReposURL         string    `json:"repos_url"`
	Type             string    `json:"type"` // User, Organization
	FullName         string    `json:"name"`
	Location         string    `json:"location"`
	Email            string    `json:"email"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (c *Client) GetUserData() (user User, err error) {
	log := c.log.With("func", "GetUserData")

	if !utils.ContainsInArray(c.conf.Scopes, "user") {
		return user, fmt.Errorf("missing scope for this action")
	}

	res, err := httpclient.New().
		Method("GET").
		Headers(c.headers()...).
		URL(fmt.Sprintf("%s/user", c.conf.HOST)).
		Do()

	if err != nil {
		log.Error("failed to make request", zap.Error(err))
		return user, err
	}
	if res.Status >= 400 {
		log.Error("failed to make request", zap.Int("status", res.Status))
		return user, err
	}
	if err = res.ScanJSON(&user); err != nil {
		return user, err
	}
	return user, nil
}
