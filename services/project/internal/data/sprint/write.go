package sprintdata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
	"pms.project/internal/data/models"
)

func (r *Repository) Create(ctx context.Context, new models.Sprint) (err error) {
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
