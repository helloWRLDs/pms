package userdata

import (
	"embed"
	"io/fs"
	"strings"

	"github.com/sirupsen/logrus"
	"pms.pkg/tools/transaction"
)

//go:embed docs/*.sql
var migrations embed.FS

func (r *Repository) ApplyMigrations() error {
	log := logrus.WithFields(logrus.Fields{
		"func": "ApplyMigrations",
	})

	tx, err := transaction.Start(r.DB)
	if err != nil {
		return err
	}
	defer func() {
		transaction.End(tx, err)
	}()

	err = fs.WalkDir(migrations, "docs", func(path string, d fs.DirEntry, err error) error {
		log := log.WithField("path", path)
		if strings.HasSuffix(path, "sql") {
			q, err := migrations.ReadFile(path)
			if err != nil {
				log.WithError(err).Error("failed read script file")
				return err
			}
			_, err = tx.Exec(string(q))
			if err != nil {
				log.WithError(err).Error("failed to apply migration")
				return err
			}
			log.Debug("applied migration")
			return nil
		}
		return nil
	})
	if err != nil {
		return err
	}
	log.Debug("migrations applied successfuly")
	return nil
}
