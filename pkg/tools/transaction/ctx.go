package transaction

import (
	"context"

	"github.com/jmoiron/sqlx"
	"pms.pkg/errs"
	"pms.pkg/logger"
)

type ContextKey string

const TX_KEY ContextKey = "TX"

func Start(ctx context.Context, db *sqlx.DB) (context.Context, error) {
	log := logger.Log.Named("tx.Start")

	if tx := Retrieve(ctx); tx != nil {
		log.Debug("tx already exists")
		return ctx, nil
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Errorw("failed to start tx", "err", err)
		return ctx, errs.ErrInternal{
			Reason: "failed to start tx",
		}
	}
	log.Debug("tx started")
	ctx = context.WithValue(ctx, TX_KEY, tx)
	return ctx, nil
}

func End(ctx context.Context, err error) {
	log := logger.Log.Named("tx.Retrieve")

	tx := Retrieve(ctx)
	if tx == nil {
		return
	}

	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Errorw("failed to rollback tx", "err", err)
		} else {
			log.Debug("tx rolled back")
		}
	} else {
		if commitErr := tx.Commit(); commitErr != nil {
			log.Errorw("failed to commit tx", "err", err)
		} else {
			log.Debug("tx committed")
		}
	}
}

func Retrieve(ctx context.Context) *sqlx.Tx {
	log := logger.Log.Named("tx.Retrieve")

	tx, ok := ctx.Value(TX_KEY).(*sqlx.Tx)
	if !ok {
		log.Debug("tx not found")
		return nil
	}
	log.Debug("tx found")
	return tx
}

func Commit(ctx context.Context) error {
	log := logger.Log.Named("tx.Commit")

	tx := Retrieve(ctx)
	if tx == nil {
		log.Debug("tx not found")
		return nil
	}

	if err := tx.Commit(); err != nil {
		log.Errorw("failed to commit tx manually")
		return err
	}
	log.Debug("tx commited manually")
	return nil
}

func Rollback(ctx context.Context) error {
	log := logger.Log.Named("tx.Rollback")

	tx := Retrieve(ctx)
	if tx == nil {
		return nil
	}
	if err := tx.Rollback(); err != nil {
		log.Errorw("failed to rollback tx manually")
		return err
	}
	log.Debug("tx rollbacked manually")
	return nil
}
