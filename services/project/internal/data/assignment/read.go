package assignmentdata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
)

func (r *Repository) GetByTask(ctx context.Context, taskID string) (assignment *AssignmentData, err error) {
	log := r.log.Named("GetByTask").With(
		zap.String("task_id", taskID),
	)
	log.Debug("GetByTask called")

	assignment = new(AssignmentData)

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("get"),
			errs.WithField("task_id", taskID),
			errs.WithObject("task_assignment"),
		)
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err := transaction.Start(ctx, r.DB)
		if err != nil {
			return nil, err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	q, a, _ := r.gen.
		Select("*").
		From(r.tableName).
		Where(sq.Eq{"task_id": taskID}).
		Limit(1).ToSql()

	if err = tx.QueryRowx(q, a...).StructScan(assignment); err != nil {
		return nil, err
	}
	return assignment, nil
}

func (r *Repository) Get(ctx context.Context, taskID, userID string) (assignment *AssignmentData, err error) {
	log := r.log.Named("Get").With(
		zap.String("task_id", taskID),
		zap.String("user_id", userID),
	)
	log.Debug("Get called")

	assignment = new(AssignmentData)

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("get"),
			errs.WithField("task_id", taskID),
			errs.WithObject("task_assignment"),
		)
	}()

	q, a, _ := r.gen.
		Select("*").
		From(r.tableName).
		Where(sq.Eq{"task_id": taskID}).
		Where(sq.Eq{"user_id": userID}).
		ToSql()

	if err = r.DB.QueryRowx(q, a...).StructScan(assignment); err != nil {
		log.Errorw("failed to select assignment", "err", err)
		return nil, err
	}

	return assignment, nil
}
