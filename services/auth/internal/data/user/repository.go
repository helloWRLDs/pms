package userdata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"pms.auth/internal/entity"
	"pms.pkg/errs"
)

var (
	_ Reader[entity.User] = &Repository{}
	_ Writer[entity.User] = &Repository{}
	_ UserRepository      = &Repository{}
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (entity.User, error)
}

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
		errctx: errs.RepositoryDetails{Object: "user", DBType: "SQLITE"},
		log:    log.Named("userdata"),
	}
}
