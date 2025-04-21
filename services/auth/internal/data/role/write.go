package roledata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.auth/internal/entity"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
)

func (r *Repository) Create(ctx context.Context, newRole entity.Role) (err error) {
	log := r.log.With(
		zap.String("func", "ListCompanies"),
		zap.Any("user_id", newRole),
	)
	log.Debug("ListCompanies called")

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
		Insert("Role").
		Columns(utils.GetColumns(newRole)...).
		Values(utils.GetArguments(newRole)...).
		ToSql()

	_, err = tx.Exec(q, a...)
	if err != nil {
		log.Errorw("failed to add participant", "err", err)
		return err
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, name string, updated entity.Role) (err error) {
	log := r.log.With(
		zap.String("func", "UpdateRole"),
		zap.String("name", name),
		zap.Any("updated_role", updated),
	)
	log.Debug("UpdateRole called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("update"),
			errs.WithObject("role"),
			errs.WithField("name", name),
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

	builder := r.gen.Update("Role")

	cols := utils.GetColumns(updated)
	args := utils.GetArguments(updated)
	for i, col := range cols {
		builder = builder.Set(col, args[i])
	}
	builder = builder.Where(sq.Eq{"name": name})
	q, a, _ := builder.ToSql()

	_, err = tx.Exec(q, a...)
	if err != nil {
		log.Errorw("failed to update role", "err", err)
		return err
	}
	return nil
}
