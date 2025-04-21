package data

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	projectdata "pms.project/internal/data/project"
)

type Repository struct {
	db *sqlx.DB

	Project *projectdata.Repository
}

func New(db *sqlx.DB, log *zap.SugaredLogger) *Repository {
	return &Repository{
		db:      db,
		Project: projectdata.New(db, log),
	}
}
