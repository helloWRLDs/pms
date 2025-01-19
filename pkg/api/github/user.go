package github

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"pms.pkg/utils"
	"pms.pkg/utils/request"
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

func (c *Client) GetUserData() (User, error) {
	log := logrus.WithField("func", "GetUserData")

	if !utils.ContainsInArray(c.Conf.Scopes, "user") {
		return User{}, fmt.Errorf("missing scope for this action")
	}

	details := request.Details{
		Method:  "GET",
		URL:     fmt.Sprintf("%s/user", c.Conf.HOST),
		Headers: c.setHeaders(),
	}
	err := details.Build()
	if err != nil {
		log.WithError(err).Error("failed to build request")
		return User{}, err
	}
	status, res, err := details.Make()
	if err != nil {
		log.WithError(err).Error("failed to make request")
		return User{}, err
	}
	if status >= 400 {
		log.WithField("res", string(res)).Error("failed to make request")
		return User{}, err
	}
	var user User
	if err = json.Unmarshal(res, &user); err != nil {
		return User{}, err
	}
	return user, nil
}
