package data

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	userdata "pms.auth/internal/data/user"
)

type Repository struct {
	db *sqlx.DB

	User *userdata.Repository
}

func New(db *sqlx.DB, log *zap.SugaredLogger) *Repository {
	return &Repository{
		db:   db,
		User: userdata.New(db, log),
	}
}
