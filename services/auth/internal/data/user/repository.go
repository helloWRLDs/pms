package userdata

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"pms.auth/internal/data"
	"pms.auth/internal/domain"
	"pms.pkg/errs"
	"pms.pkg/types/ctx"
	"pms.pkg/types/list"
)

type Repository struct {
	data.Reader

	DB     *sqlx.DB
	gen    sq.StatementBuilderType
	errctx errs.RepositoryDetails
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		DB:     db,
		gen:    sq.StatementBuilder.PlaceholderFormat(sq.Question),
		errctx: errs.RepositoryDetails{Object: "user"},
	}
}

func (r *Repository) GetByEmail(ctx ctx.Context, email string) (user domain.User, err error) {
	log := logrus.
		WithField("email", email).
		WithField("func", "GetByEmail")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithField("email", email), errs.WithOperation("retrieve"))
	}()

	tx, err := ctx.StartTx(r.DB)
	if err != nil {
		log.WithError(err).Error("failed to start tx")
		return user, err
	}

	query, args, _ := r.gen.
		Select("*").
		From("users").
		Where(sq.Eq{"email": email}).
		ToSql()

	if err := tx.QueryRowx(query, args...).StructScan(&user); err != nil {
		log.WithError(err).Warn("failed to fetch user by email")
		return user, err
	}
	return user, nil
}

func (r *Repository) Count(ctx ctx.Context, filter list.Filters) int {
	return 1
}

func (r *Repository) Exists(ctx ctx.Context, email string) bool {
	return true
}

func (r *Repository) Get(ctx ctx.Context, id string) (user domain.User, err error) {
	log := logrus.
		WithField("id", id).
		WithField("func", "Get")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithField("id", id), errs.WithOperation("retrieve"))
	}()

	tx, err := ctx.StartTx(r.DB)
	if err != nil {
		log.WithError(err).Error("failed to start tx")
		return user, err
	}

	query, args, _ := r.gen.
		Select("*").
		From("users").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err = tx.QueryRowx(query, args...).StructScan(&user); err != nil {
		log.WithError(err).Warn("failed to fetch user by email")
		return user, err
	}
	return user, nil
}

func (r *Repository) List(ctx ctx.Context, filter list.Filters) (list.List[domain.User], error) {
	return list.List[domain.User]{}, nil
}
