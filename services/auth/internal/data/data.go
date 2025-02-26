package data

import (
	"github.com/jmoiron/sqlx"
	userdata "pms.auth/internal/data/user"
)

type Repository struct {
	User *userdata.Repository
}

func New(db *sqlx.DB) *Repository {
	repo := Repository{
		User: userdata.New(db),
	}

	return &repo
}
