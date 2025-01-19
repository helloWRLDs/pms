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

	// err = fs.WalkDir(migrations, "docs", func(path string, d fs.DirEntry, err error) error {
	// 	log := log.WithField("path", path)
	// 	if strings.HasSuffix(path, "sql") {
	// 		q, err := migrations.ReadFile(path)
	// 		if err != nil {
	// 			log.WithError(err).Error("failed read script file")
	// 			return err
	// 		}
	// 		_, err = tx.Exec(string(q))
	// 		if err != nil {
	// 			log.WithError(err).Error("failed to apply migration")
	// 			return err
	// 		}
	// 		log.Debug("applied migration")
	// 		return nil
	// 	}
	// 	return nil
	// })
	// if err != nil {
	// 	return err
	// }
	log.Debug("migrations applied successfuly")
	return nil
}
