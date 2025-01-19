package userdata

import (
	"github.com/jmoiron/sqlx"
	"pms.auth/internal/data"
	"pms.pkg/errs"
	"pms.pkg/tools/sqlbuilder"
)

var (
	_ data.Reader = &Repository{}
)

type Repository struct {
	DB     *sqlx.DB
	gen    *sqlbuilder.Builder
	errctx errs.RepositoryDetails
}

func New(db *sqlx.DB) *Repository {
	repo := &Repository{
		DB:     db,
		gen:    sqlbuilder.New(sqlbuilder.SQLITE),
		errctx: errs.RepositoryDetails{Object: "user"},
	}
	repo.gen.IgnoreOnInsert(
		"id",
		"created_at",
		"updated_at",
	)
	repo.ApplyMigrations()

	return repo
}
