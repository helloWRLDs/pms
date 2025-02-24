package data

import (
	"github.com/jmoiron/sqlx"
	roledata "pms.auth/internal/data/roles"
	userdata "pms.auth/internal/data/user"
)

type Repository struct {
	User  *userdata.Repository
	Roles *roledata.Repository
}

func New(db *sqlx.DB) *Repository {
	repo := Repository{
		User:  userdata.New(db),
		Roles: roledata.New(db),
	}
	repo.RevertMigrations(db)
	repo.ApplyMigrations(db)

	return &repo
}
