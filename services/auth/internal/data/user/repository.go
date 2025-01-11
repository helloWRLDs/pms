package userdata

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"pms.auth/internal/data"
	"pms.pkg/errs"
)

type Repository struct {
	data.Reader

	DB     *sqlx.DB
	gen    sq.StatementBuilderType
	errctx errs.RepositoryDetails
}

func New(db *sqlx.DB) *Repository {
	repo := &Repository{
		DB:     db,
		gen:    sq.StatementBuilder.PlaceholderFormat(sq.Question),
		errctx: errs.RepositoryDetails{Object: "user"},
	}
	repo.ApplyMigrations()

	return repo
}
