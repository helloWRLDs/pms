package userdata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"pms.auth/internal/entity"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/type/list"
)

func (r *Repository) Create(ctx context.Context, newUser entity.User) (err error) {
	log := r.log.With(
		zap.String("func", "CreateUser"),
		zap.String("email", newUser.Email),
		zap.String("name", newUser.Name),
	)
	log.Debug("CreateUser called")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithOperation("create"), errs.WithField("email", newUser.Email))
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		log.Debug("tx not found")
		ctx, err = transaction.Start(ctx, r.DB)
		if err != nil {
			return err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	query, args, _ := r.gen.
		Insert("User").
		Columns("name", "email", "password").
		Values(newUser.Name, newUser.Email, newUser.Password).
		ToSql()

	if _, err := tx.Exec(query, args...); err != nil {
		log.Errorw("failed to create user", "err", err)
		return err
	}
	log.Debug("user created")
	return nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (user entity.User, err error) {
	log := r.log.With(
		zap.String("func", "GetByEmail"),
		zap.String("email", email),
	)
	log.Debug("CreateUser called")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithField("email", email), errs.WithOperation("retrieve"))
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err = transaction.Start(ctx, r.DB)
		if err != nil {
			return
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	query, args, _ := r.gen.
		Select("*").
		From("User").
		Where(sq.Eq{"email": email}).
		ToSql()

	if err := tx.QueryRowx(query, args...).StructScan(&user); err != nil {
		log.Warnw("failed to fetch user by email", "err", err)
		return user, err
	}
	return user, nil
}

func (r *Repository) Count(ctx context.Context, filter list.Filters) int {

	return 1
}

func (r *Repository) Exists(ctx context.Context, email string) bool {
	var exists bool
	query := `
		SELECT EXISTS(
			SELECT * FROM User WHERE email = ?
		);
	`
	if err := r.DB.QueryRowx(query, email).Scan(&exists); err != nil {
		return false
	}
	return exists
}

func (r *Repository) GetByID(ctx context.Context, id string) (user entity.User, err error) {
	log := logrus.
		WithField("id", id).
		WithField("func", "Get")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithField("id", id), errs.WithOperation("retrieve"))
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err = transaction.Start(ctx, r.DB)
		if err != nil {
			return
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	query, args, _ := r.gen.
		Select("*").
		From("user").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err = tx.QueryRowx(query, args...).StructScan(&user); err != nil {
		log.WithError(err).Warn("failed to fetch user by email")
		return user, err
	}
	return user, nil
}

func (r *Repository) List(ctx context.Context, filter list.Filters) (list.List[entity.User], error) {
	return list.List[entity.User]{}, nil
}
