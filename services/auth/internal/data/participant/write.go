package participantdata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
)

func (r *Repository) Create(ctx context.Context, participant Participant) (err error) {
	log := r.log.With(
		zap.String("func", "Create"),
		zap.String("user_id", participant.UserID),
		zap.String("company_id", participant.CompanyID),
	)
	log.Debug("Create called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("create"),
			errs.WithObject("participant"),
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
		Columns(utils.GetColumns(participant)...).
		Values(utils.GetArguments(participant)...).
		ToSql()

	log.Info("Query: ", q)

	if _, err = tx.Exec(q, a...); err != nil {
		log.Errorw("failed to create participant", "err", err)
		return err
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, participantID string, updated Participant) (err error) {
	log := r.log.With(
		zap.String("func", "Update"),
		zap.String("user_id", participantID),
	)
	log.Debug("Update called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("update"),
			errs.WithObject("participant"),
			errs.WithField("id", participantID),
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

	builder := r.gen.
		Update(r.tableName)

	cols := utils.GetColumns(updated)
	args := utils.GetArguments(updated)

	for i, col := range cols {
		builder = builder.Set(col, args[i])
	}
	builder = builder.Where(sq.Eq{"id": participantID})
	q, a, _ := builder.ToSql()
	if _, err = tx.Exec(q, a...); err != nil {
		log.Errorw("failed to update participant", "err", err)
		return err
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, participantID string) (err error) {
	log := r.log.With(
		zap.String("func", "Delete"),
		zap.String("user_id", participantID),
	)
	log.Debug("Delete called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("delete"),
			errs.WithObject("participant"),
			errs.WithField("id", participantID),
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

	q, a, _ := r.gen.Delete(r.tableName).Where(sq.Eq{"id": participantID}).ToSql()
	if _, err = tx.Exec(q, a...); err != nil {
		log.Errorw("failed to delete participant", "err", err)
		return err
	}
	return nil
}
