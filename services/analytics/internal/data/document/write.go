package documentdata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
)

func (r *Repository) Update(ctx context.Context, id string, updatedDoc Document) (err error) {
	log := r.log.Named("Update").With(
		zap.Any("updated", updatedDoc),
		zap.String("id", id),
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
	cols := utils.GetColumns(updatedDoc)
	args := utils.GetArguments(updatedDoc)

	builder := r.gen.Update(r.tableName)

	for i, col := range cols {
		builder = builder.Set(col, args[i])
	}
	q, a, _ := builder.Where(sq.Eq{"id": id}).ToSql()

	if _, err = tx.Exec(q, a...); err != nil {
		log.Errorw("failed to update doc", "err", err)
		return err
	}
	return nil
}

func (r *Repository) Create(ctx context.Context, document Document) (err error) {
	log := r.log.Named("Create").With(
		zap.Any("name", document),
	)
	log.Debug("Create called")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithOperation("create"))
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
	q, a, _ := r.gen.
		Insert(r.tableName).
		Columns(utils.GetColumns(document)...).
		Values(utils.GetArguments(document)...).
		ToSql()

	if res, err := tx.Exec(q, a...); err != nil {
		log.Errorw("failed to create document", "res", res, "err", err)
		return err
	} else {
		log.Infof("result: %#v", res)
	}
	return nil
}
