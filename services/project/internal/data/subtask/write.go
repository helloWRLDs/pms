package subtaskdata

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
)

func (r *Repository) Create(ctx context.Context, task SubTask) (err error) {
	log := r.log.With(
		zap.String("func", "Create"),
		zap.Any("new_sub_task", task),
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
		Columns(utils.GetColumns(task)...).
		Values(utils.GetArguments(task)...).ToSql()

	log.Debugw("query built", "query", q, "args", a)

	if _, err = tx.ExecContext(ctx, q, a...); err != nil {
		log.Errorw("failed to insert task", "err", err)
		return err
	}
	return nil
}
