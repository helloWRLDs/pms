package userrepo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"pms.pkg/utils"
	"pms.users/internal/domain"
)

type Repository struct {
	db  *sqlx.DB
	gen sq.StatementBuilderType
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db:  db,
		gen: sq.StatementBuilder.PlaceholderFormat(sq.Question),
	}
}

func (r *Repository) Create(ctx context.Context, user domain.User) (string, error) {

	tx := r.db.MustBeginTx(ctx, nil)

	q, args, _ := r.gen.
		Insert("Users").
		Columns("id", "first_name", "last_name", "email", "password").
		Values(user.ID, user.FirstName, user.LastName, user.Email, user.Password).
		ToSql()

	res, err := r.db.ExecContext(ctx, q, args...)
	println(utils.JSON(res))
	if err != nil {
		return "", err
	}
	return user.ID.String(), nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var usr domain.User
	q, args, _ := r.gen.
		Select("*").
		From("users").
		Where(sq.Eq{"id": id}).
		ToSql()
	err := r.db.QueryRowxContext(ctx, q, args...).StructScan(&usr)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}
