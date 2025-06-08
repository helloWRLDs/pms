package roledata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
)

func (r *Repository) Create(ctx context.Context, newRole Role) (err error) {
	log := r.log.With(
		zap.String("func", "Create"),
		zap.Any("role", newRole),
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

func (r *Repository) Update(ctx context.Context, name string, updated Role) (err error) {
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

func (r *Repository) Delete(ctx context.Context, name string) (err error) {
	log := r.log.With(
		zap.String("func", "DeleteRole"),
		zap.String("name", name),
	)
	log.Debug("DeleteRole called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("delete"),
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

	q, a, _ := r.gen.Delete(r.tableName).Where(sq.Eq{"name": name}).ToSql()
	_, err = tx.Exec(q, a...)
	if err != nil {
		log.Errorw("failed to delete role", "err", err)
		return err
	}
	return nil
}
