package sqlite

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/cznic/sqlite"
)

func Open(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Close(db *sqlx.DB) error {
	return db.Close()
}
