package github

import (
	"fmt"

	"go.uber.org/zap"
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
	log := c.log.With("func", "GetCommitDetails")

	res, err := httpclient.New().
		Method("GET").
		Headers(
			append([]string{"Accept", "application/vnd.github.v3+json"}, c.headers()...)...,
		).
		URL(fmt.Sprintf(
			"%s/repos/%s/%s/commits/%s",
			c.conf.HOST, owner, repo, commitSHA,
		)).Do()
	if err != nil || res.Status >= 400 {
		log.Error("failed to make request", zap.Int("status", res.Status))
		return
	}

	if err = res.ScanJSON(&commit); err != nil {
		log.Error("failed to unmarshal response", zap.Error(err))
		return
	}
	log.Debug("fetched commit", zap.Any("commit", commit))
	return commit, nil
}
