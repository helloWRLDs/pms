package transactions

import (
	"context"

	"github.com/jmoiron/sqlx"
)

const txKey = "tx"

func Start(ctx context.Context, db *sqlx.DB) context.Context {
	tx, _ := db.BeginTxx(ctx, nil)
	ctx = context.WithValue(ctx, txKey, tx)
	return ctx
}

func End(ctx context.Context, err error) {
	tx, ok := ctx.Value(txKey).(*sqlx.Tx)
	if !ok {
		return
	}
	if err == nil {
		if commitErr := tx.Commit(); commitErr != nil {
			println("failed to commit tx")
		}
		println("tx commited")
	}
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		println("tx rollback")
	}
	println("failed to rollback tx")
}
