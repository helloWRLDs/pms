package projectdata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
)

func (r *Repository) Update(ctx context.Context, id string, updated Project) (err error) {
	log := r.log.With(
		zap.String("func", "Update"),
		zap.String("id", id),
		zap.Any("updated_project", updated),
	)
	log.Debug("Update called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("update"),
			errs.WithObject("project"),
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
		Update(r.tableName).
		Where(sq.Eq{"id": id})

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

func (r *Repository) Create(ctx context.Context, project Project) (err error) {
	log := r.log.With(
		zap.String("func", "Create"),
		zap.Any("new_project", project),
	)
	log.Debug("Create called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("create"),
			errs.WithObject("project"),
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
		Columns(utils.GetColumns(project)...).
		Values(utils.GetArguments(project)...).ToSql()

	_, err = tx.ExecContext(ctx, q, a...)
	if err != nil {
		log.Errorw("failed to insert project", "err", err)
		return err
	}

	return nil
}
