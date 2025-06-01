package logic

import (
	"go.uber.org/zap"
	"pms.auth/internal/config"
	"pms.auth/internal/data"
	"pms.pkg/api/github"
	"pms.pkg/api/google"
)

type Logic struct {
	Repo *data.Repository
	conf *config.Config

	log *zap.SugaredLogger

	googleClient *google.Client
	githubClient *github.Client
}

func New(repo *data.Repository, conf *config.Config, log *zap.SugaredLogger) *Logic {
	conf.GoogleConfig.Scopes = []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	}
	return &Logic{
		Repo:         repo,
		conf:         conf,
		log:          log,
		googleClient: google.New(conf.GoogleConfig, log),
		githubClient: github.New(conf.GitHubConfig, log),
	}
}
