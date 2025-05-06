package taskdata

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"pms.pkg/errs"
)

type Repository struct {
	tableName string
	DB        *sqlx.DB
	gen       sq.StatementBuilderType
	errctx    errs.RepositoryDetails

	log *zap.SugaredLogger
}

func New(db *sqlx.DB, log *zap.SugaredLogger) *Repository {
	return &Repository{
		tableName: "Task",
		DB:        db,
		gen:       sq.StatementBuilder.PlaceholderFormat(sq.Question),
		errctx:    errs.RepositoryDetails{Object: "task", DBType: "SQLITE"},
		log:       log.Named("taskdata"),
	}
}