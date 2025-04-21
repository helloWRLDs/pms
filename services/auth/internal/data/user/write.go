package userdata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.auth/internal/entity"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
)

func (r *Repository) Update(ctx context.Context, id string, user entity.User) (err error) {
	log := r.log.With(
		zap.String("func", "Update"),
		zap.Any("updated_user", user),
		zap.String("id", id),
	)
	log.Debug("Update called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("update"),
			errs.WithObject("user"),
			errs.WithField("id", id),
		)
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err := transaction.Start(ctx, r.DB)
		if err != nil {
			return err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	cols := utils.GetColumns(user)
	args := utils.GetArguments(user)

	builder := r.gen.
		Update(r.tableName)

	for i, col := range cols {
		builder = builder.Set(col, args[i])
	}

	q, a, _ := builder.Where(sq.Eq{"id": id}).ToSql()

	if _, err = tx.Exec(q, a...); err != nil {
		log.Errorw("failed to update user", "err", err)
		return err
	}
	return nil
}

func (r *Repository) Create(ctx context.Context, user entity.User) (err error) {
	log := r.log.With(
		zap.String("func", "Create"),
		zap.Any("new_user", user),
	)
	log.Debug("Create called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("create"),
			errs.WithObject("role"),
		)
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err := transaction.Start(ctx, r.DB)
		if err != nil {
			return err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	q, a, _ := r.gen.
		Insert(r.tableName).
		Columns(utils.GetColumns(user)...).
		Values(utils.GetArguments(user)...).
		ToSql()

	if _, err = tx.Exec(q, a...); err != nil {
		log.Errorw("failed to create user", "err", err)
		return err
	}
	return nil
}
