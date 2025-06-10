package assignmentdata

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
)

func (r *Repository) Delete(ctx context.Context, assignment AssignmentData) (err error) {
	log := r.log.Named("Delete").With(
		zap.String("task_id", assignment.TaskID),
		zap.String("user_id", assignment.UserID),
	)
	log.Debug("Delete called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("delete"),
			errs.WithField("task_id", assignment.TaskID),
			errs.WithObject("task_assignment"),
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
		Where("task_id = $1 AND user_id = $2", assignment.TaskID, assignment.UserID).
		ToSql()

	if _, err = tx.Exec(q, a...); err != nil {
		log.Errorw("failed to delete assignment", "err", err)
		return err
	}
	return nil
}

func (r *Repository) Create(ctx context.Context, assignment AssignmentData) (err error) {
	log := r.log.Named("Create").With(
		zap.String("task_id", assignment.TaskID),
		zap.String("user_id", assignment.UserID),
	)
	log.Debug("Create called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("create"),
			errs.WithField("task_id", assignment.TaskID),
			errs.WithObject("task_assignment"),
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
		Columns(utils.GetColumns(assignment)...).
		Values(utils.GetArguments(assignment)...).
		ToSql()

	if _, err = tx.Exec(q, a...); err != nil {
		log.Errorw("failed to create assignment", "err", err)
		return err
	}
	return nil
}
