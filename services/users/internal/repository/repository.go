package userrepo

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"pms.pkg/interfaces"
	"pms.pkg/types/ctx"
	"pms.pkg/utils"
	"pms.users/internal/domain"
)

type Repository struct {
	interfaces.Reader[domain.User]
	interfaces.Writer[domain.User]

	DB  *sqlx.DB
	gen sq.StatementBuilderType
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		DB:  db,
		gen: sq.StatementBuilder.PlaceholderFormat(sq.Question),
	}
}

func (r *Repository) Get(ctx ctx.Context, field, value string) (domain.User, error) {
	tx, _ := ctx.StartTx(r.DB)

	var usr domain.User

	query, args, _ := r.gen.
		Select("*").
		From("users").
		Where(sq.Eq{field: value}).
		ToSql()

	if err := tx.QueryRowx(query, args...).Scan(&usr); err != nil {
		return usr, err
	}

	return usr, nil
}

// func (r *Repository) List(ctx ctx.Context, filters types.Filters) (types.List[domain.User], error) {
// 	query, args, _ := r.gen.
// 		Select("").
// 		OrderBy(filters.OrderBy).
// 		ToSql()
// }

func (r *Repository) Create(ctx ctx.Context, user domain.User) (string, error) {
	tx, _ := ctx.StartTx(r.DB)
	q, args, _ := r.gen.
		Insert("Users").
		Columns("id", "first_name", "last_name", "email", "password").
		Values(user.ID, user.FirstName, user.LastName, user.Email, user.Password).
		ToSql()

	res, err := tx.ExecContext(&ctx, q, args...)
	println(utils.JSON(res))
	if err != nil {
		return "", err
	}
	return user.ID.String(), nil
}

func (r Repository) GetByID(ctx ctx.Context, id string) (*domain.User, error) {
	tx, _ := ctx.StartTx(r.DB)
	var usr domain.User
	q, args, _ := r.gen.
		Select("*").
		From("users").
		Where(sq.Eq{"id": id}).
		ToSql()
	err := tx.QueryRowxContext(&ctx, q, args...).Scan(&usr)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}
