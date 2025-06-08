package taskcommentdata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
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

func (r *Repository) Delete(ctx context.Context, commentID string) (err error) {
	log := r.log.Named("Delete").With(
		zap.String("comment_id", commentID),
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

	q, a, _ := r.gen.Delete(r.tableName).Where(sq.Eq{"id": commentID}).ToSql()

	if res, err := tx.Exec(q, a...); err != nil {
		log.Errorw("failed to delete comment", "err", err)
		return err
	} else {
		log.Infow("deleted comment", "result", res)
	}
	return nil
}
