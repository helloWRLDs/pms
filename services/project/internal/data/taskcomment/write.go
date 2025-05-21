package taskcommentdata

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
)

func (r *Repository) Create(ctx context.Context, newComment TaskComment) (err error) {
	log := r.log.Named("Create").With(
		zap.Any("new_comment", newComment),
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
		Columns(utils.GetColumns(newComment)...).
		Values(utils.GetArguments(newComment)...).
		ToSql()

	if res, err := tx.Exec(q, a...); err != nil {
		log.Errorw("failed to create comment", "err", err)
		return err
	} else {
		log.Infow("created comment", "result", res)
	}
	return nil
}
