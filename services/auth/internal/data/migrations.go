package data

import (
	"embed"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

//go:embed migrations/*.sql
var migrations embed.FS

func (r *Repository) ApplyMigrations(db *sqlx.DB) error {
	log := logrus.WithFields(logrus.Fields{
		"func": "ApplyMigrations",
	})

	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.WithError(err).Error("failed to set dialect - sqlite3")
		return err
	}

	if err := goose.Up(db.DB, "migrations"); err != nil {
		log.WithError(err).Error("failed to apply migrations")
		return err
	}

	log.Debug("migrations applied successfuly")
	return nil
}

func (r *Repository) RevertMigrations(db *sqlx.DB) error {
	log := logrus.WithFields(logrus.Fields{
		"func": "RevertMigrations",
	})

	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.WithError(err).Error("failed to set dialect - sqlite3")
		return err
	}

	if err := goose.Down(db.DB, "migrations"); err != nil {
		log.WithError(err).Error("failed to revert migrations")
		return err
	}
	log.Debug("failed to revert migrations")
	return nil
}
