package participantdata

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
		tableName: "\"Participant\"",
		DB:        db,
		gen:       sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		errctx:    errs.RepositoryDetails{Object: "participant", DBType: "POSTGRES"},
		log:       log.Named("participantdata"),
	}
}
