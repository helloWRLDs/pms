package logic

import (
	userdata "pms.auth/internal/data/user"
)

type AuthLogic struct {
	repo *userdata.Repository
}

func New(repo *userdata.Repository) *AuthLogic {
	return &AuthLogic{
		repo: repo,
	}
}

// func (l *AuthLogic) RegsterUser() (domain.User, error) {

// }
