package userdata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
	userdomain "pms.auth/internal/domain/user"
	"pms.pkg/datastore/utils"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/type/list"
)

func (r *Repository) CreateUser(ctx context.Context, newUser userdomain.User) (err error) {
	log := logrus.WithFields(logrus.Fields{
		"func":      "CreateUser",
		"email":     newUser.Email,
		"full_name": newUser.FullName,
	})
	log.Debug("CreateUser called")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithOperation("create"), errs.WithField("email", newUser.Email))
	}()

	tx, err := transaction.RetrieveCTX(ctx)
	if err != nil {
		tx, err = transaction.Start(r.DB)
		if err != nil {
			return err
		}
		defer func() {
			transaction.End(tx, err)
		}()
	}

	mapper := utils.
		NewEnityMapper(newUser).
		Ignore("id", "created_at", "updated_at")

	query, args, _ := r.gen.
		Insert("user").
		Columns(mapper.Columns()...).
		Values(mapper.Values()...).
		ToSql()

	if _, err := tx.Exec(query, args...); err != nil {
		log.WithError(err).Error("failed to create user")
		return err
	}
	log.Debug("user created")
	return nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (user userdomain.User, err error) {
	log := logrus.
		WithField("email", email).
		WithField("func", "GetByEmail")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithField("email", email), errs.WithOperation("retrieve"))
	}()

	tx, err := transaction.RetrieveCTX(ctx)
	if err != nil {
		tx, err = transaction.Start(r.DB)
		if err != nil {
			return
		}
		defer func() {
			transaction.End(tx, err)
		}()
	}

	query, args, _ := r.gen.
		Select("*").
		From("user").
		Where(sq.Eq{"email": email}).
		ToSql()

	if err := tx.QueryRowx(query, args...).StructScan(&user); err != nil {
		log.WithError(err).Warn("failed to fetch user by email")
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

func (r *Repository) GetByID(ctx context.Context, id string) (user userdomain.User, err error) {
	log := logrus.
		WithField("id", id).
		WithField("func", "Get")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithField("id", id), errs.WithOperation("retrieve"))
	}()

	tx, err := transaction.RetrieveCTX(ctx)
	if err != nil {
		tx, err = transaction.Start(r.DB)
		if err != nil {
			return
		}
		defer func() {
			transaction.End(tx, err)
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

func (r *Repository) List(ctx context.Context, filter list.Filters) (list.List[userdomain.User], error) {
	return list.List[userdomain.User]{}, nil
}
