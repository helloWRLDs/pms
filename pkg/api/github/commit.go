package github

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"pms.pkg/tools/httpclient"
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

func (c *Client) GetCommitDetails(owner, repo, commitSHA string) (commit CommitDetails, err error) {
	log := logrus.WithField("func", "GetCommitDetails")

	res, err := httpclient.New().
		Method("GET").
		Headers(
			append([]string{"Accept", "application/vnd.github.v3+json"}, c.headers()...)...,
		).
		URL(fmt.Sprintf(
			"%s/repos/%s/%s/commits/%s",
			c.Conf.HOST, owner, repo, commitSHA,
		)).Do()
	if err != nil || res.Status >= 400 {
		log.WithField("res", string(res.Status)).Error("failed to make request")
		return
	}

	if err = res.ScanJSON(&commit); err != nil {
		log.WithError(err).Error("failed to unmarshal response")
		return
	}
	log.WithField("commit", commit).Debug("fetched commit")
	return commit, nil
}
