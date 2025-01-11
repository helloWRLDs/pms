package transaction

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func Start(db *sqlx.DB) (*sqlx.Tx, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func End(tx *sqlx.Tx, err error) {
	log := logrus.WithFields(logrus.Fields{
		"func": "transaction.Handle",
	})
	log.Debug("func called")

	if tx == nil {
		log.Debug("tx is nil, skipping rollback/commit")
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
