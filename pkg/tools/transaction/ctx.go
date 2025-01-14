package transaction

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"pms.pkg/errs"
)

type ContextKey string

const TX_KEY ContextKey = "TX"

func StartCTX(ctx context.Context, db *sqlx.DB) (context.Context, error) {
	log := logrus.WithFields(logrus.Fields{
		"func": "transaction.Start",
	})

	_, err := RetrieveCTX(ctx)
	if err == nil {
		log.Debug("tx already started")
		return ctx, nil
	}

	tx, err := db.Beginx()
	if err != nil {
		log.WithError(err).Error("failed to start tx")
		return ctx, errs.ErrInternal{
			Reason: "failed to start tx",
		}
	}
	log.Debug("tx started")
	ctx = context.WithValue(ctx, TX_KEY, tx)
	return ctx, nil
}

func RetrieveCTX(ctx context.Context) (*sqlx.Tx, error) {
	log := logrus.WithFields(logrus.Fields{
		"func": "transaction.Retrieve",
	})

	tx, ok := ctx.Value(TX_KEY).(*sqlx.Tx)
	if !ok {
		log.Debug("tx not found")
		return nil, errs.ErrNotFound{
			Object: "tx",
		}
	}
	return tx, nil
}

func EndCTX(ctx context.Context, err error) {
	log := logrus.WithFields(logrus.Fields{
		"func":   "transaction.End",
		"hasErr": err != nil,
	})

	tx, retieveErr := RetrieveCTX(ctx)
	if retieveErr != nil {
		log.WithError(err).Error("failed to retrieve tx")
		return
	}
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.WithError(rollbackErr).Error("failed to rollback tx")
		} else {
			log.Debug("tx rollback successful")
		}
	} else {
		if commitErr := tx.Commit(); commitErr != nil {
			log.WithError(commitErr).Error("failed to commit tx")
		} else {
			log.Debug("tx commit successful")
		}
	}
}
