package data

import (
	"context"

	"pms.pkg/tools/transaction"
)

func (r *Repository) StartTx(ctx context.Context) (context.Context, error) {
	ctx, err := transaction.Start(ctx, r.db)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func (r *Repository) EndTx(ctx context.Context, err error) {
	transaction.End(ctx, err)
}
