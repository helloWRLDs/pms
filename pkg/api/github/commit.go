package github

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"pms.pkg/utils/request"
)

type CommitDetails struct {
	SHA     string  `json:"sha"`
	Message string  `json:"commit"`
	URL     string  `json:"html_url"`
	Author  Author  `json:"author"`
	Files   []Files `json:"files"`
}

type Files struct {
	Name  string `json:"filename"`
	Patch string `json:"patch"`
}

type Author struct {
	Login string `json:"login"`
}

func (c *Client) GetCommitDetails(owner, repo, commitSHA string) (CommitDetails, error) {
	log := logrus.WithField("func", "GetCommitDetails")

	details := request.Details{
		Method: "GET",
		URL:    fmt.Sprintf("%s/repos/%s/%s/commits/%s", c.Conf.HOST, owner, repo, commitSHA),
		Headers: c.setHeaders(
			"Accept", "application/vnd.github.v3+json",
		),
	}

	err := details.Build()
	if err != nil {
		log.WithError(err).Error("failed to build request")
		return CommitDetails{}, err
	}
	status, res, err := details.Make()
	if err != nil {
		log.WithError(err).Error("failed to make request")
		return CommitDetails{}, err
	}
	if status >= 400 {
		log.WithField("res", string(res)).Error("failed to make request")
		return CommitDetails{}, err
	}
	var commit CommitDetails
	if err = json.Unmarshal(res, &commit); err != nil {
		log.WithError(err).Error("failed to unmarshal response")
		return CommitDetails{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	log.WithField("commit", commit).Debug("fetched commit")
	return CommitDetails{}, nil
}
