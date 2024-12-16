package ctx

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func (c *Context) StartTx(db *sqlx.DB) (*sqlx.Tx, error) {
	log := logrus.WithField("func", "StartTx")

	if c.tx != nil {
		log.Debug("tx already started")
		return c.tx, nil
	}

	tx, err := db.BeginTxx(c, nil)
	if err != nil {
		log.WithError(err).Error("failed to start tx")
		return nil, err
	}

	log.Debug("tx started")
	c.tx = tx
	return c.tx, nil
}

func (c *Context) EndTx(err error) {
	log := logrus.WithField("func", "EndTx")

	if c.tx == nil {
		log.Debug("tx already ended")
		return
	}

	if err == nil {
		c.Commit()
	} else {
		c.Rollback()
	}

	log.Debug("tx ended")
	c.tx = nil
}

func (c *Context) Commit() {
	log := logrus.WithField("func", "Commit")
	if c.tx == nil {
		log.Debug("tx already ended")
		return
	}
	defer func() {
		c.tx = nil
	}()

	if err := c.tx.Commit(); err != nil {
		log.WithError(err).Error("failed to commit tx")
		return
	}
	log.Debug("commit successful")
}

func (c *Context) Rollback() {
	log := logrus.WithField("func", "Rollback")
	if c.tx == nil {
		return
	}
	defer func() {
		c.tx = nil
	}()
	if err := c.tx.Rollback(); err != nil {
		log.WithError(err).Error("failed to rollback tx")
		return
	}
	log.Debug("rollback successful")
}
