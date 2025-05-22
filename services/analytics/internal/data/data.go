package data

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	documentdata "pms.analytics/internal/data/document"
)

type Repository struct {
	db *sqlx.DB

	Document *documentdata.Repository
}

func New(db *sqlx.DB, log *zap.SugaredLogger) *Repository {
	return &Repository{
		db:       db,
		Document: documentdata.New(db, log),
	}
}
