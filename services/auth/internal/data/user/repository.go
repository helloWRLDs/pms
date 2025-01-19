package userdata

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"pms.pkg/errs"
)

var (
	_ Reader = &Repository{}
)

type Repository struct {
	DB     *sqlx.DB
	gen    sq.StatementBuilderType
	errctx errs.RepositoryDetails
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		DB:     db,
		gen:    sq.StatementBuilder.PlaceholderFormat(sq.Question),
		errctx: errs.RepositoryDetails{Object: "user", DBType: "SQLITE"},
	}
}
