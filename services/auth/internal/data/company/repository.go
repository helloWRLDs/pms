package companydata

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	comp "pms.auth/internal/entity/company"
	"pms.pkg/errs"
)

var (
	_ Reader[comp.Company] = &Repository{}
	_ Writer[comp.Company] = &Repository{}
)

type Repository struct {
	DB     *sqlx.DB
	gen    sq.StatementBuilderType
	errctx errs.RepositoryDetails

	log *zap.SugaredLogger
}

func New(db *sqlx.DB, log *zap.SugaredLogger) *Repository {
	return &Repository{
		DB:     db,
		gen:    sq.StatementBuilder.PlaceholderFormat(sq.Question),
		errctx: errs.RepositoryDetails{Object: "company", DBType: "SQLITE"},
		log:    log.Named("companydata"),
	}
}
