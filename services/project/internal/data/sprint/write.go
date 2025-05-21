package sprintdata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
)

func (r *Repository) Update(ctx context.Context, id string, updated Sprint) (err error) {
	log := r.log.Named("Update").With(
		zap.Any("updated", updated),
	)
	log.Debug("Update called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("update"),
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

	cols := utils.GetColumns(updated)
	args := utils.GetArguments(updated)

	builder := r.gen.Update(r.tableName)
	for i, col := range cols {
		builder = builder.Set(col, args[i])
	}
	builder = builder.Where(sq.Eq{"id": id})

	q, a, _ := builder.ToSql()

	if _, err = tx.ExecContext(ctx, q, a...); err != nil {
		log.Errorw("failed to update sprint", "err", err)
		return err
	}
	return nil
}

func (r *Repository) Create(ctx context.Context, new Sprint) (err error) {
	log := r.log.With(
		zap.String("func", "Create"),
		zap.Any("new_sprint", new),
	)
	log.Debug("Create called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("create"),
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
		Columns(utils.GetColumns(new)...).
		Values(utils.GetArguments(new)...).ToSql()

	if _, err = tx.ExecContext(ctx, q, a...); err != nil {
		log.Errorw("failed to create sprint", "err", err)
		return err
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) (err error) {
	log := r.log.With(
		zap.String("func", "Delete"),
		zap.String("id", id),
	)
	log.Debug("Delete called")
	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("delete"),
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
		Delete(r.tableName).
		Where(sq.Eq{"id": id}).ToSql()

	if _, err = tx.ExecContext(ctx, q, a...); err != nil {
		log.Errorw("failed to delete sprint", "err", err)
		return err
	}
	return nil
}
func (r *Repository) Update(ctx context.Context, id string, updated models.Sprint) (err error) {
	log := r.log.With(
		zap.String("func", "Update"),
		zap.String("id", id),
		zap.Any("updated_sprint", updated),
	)
	log.Debug("Update called")
	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("update"),
			errs.WithObject("sprint"),
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

	cols := utils.GetColumns(updated)
	args := utils.GetArguments(updated)

	builder := r.gen.
		Update(r.tableName)
	for i, col := range cols {
		builder = builder.Set(col, args[i])
	}
	q, a, _ := builder.ToSql()

	if _, err = tx.ExecContext(ctx, q, a...); err != nil {
		log.Errorw("failed to update task", "err", err)
		return err
	}
	return nil
}
