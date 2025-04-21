package projectdata

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
	"pms.project/internal/entity"
)

func (r *Repository) Create(ctx context.Context, project entity.Project) (err error) {
	log := r.log.With(
		zap.String("func", "Create"),
		zap.Any("new_project", project),
	)
	log.Debug("Create called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("create"),
			errs.WithObject("project"),
		)
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err := transaction.Start(ctx, r.DB)
		if err != nil {
			return err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	q, a, _ := r.gen.
		Insert(r.tableName).
		Columns(utils.GetColumns(project)...).
		Values(utils.GetArguments(project)...).ToSql()

	_, err = tx.ExecContext(ctx, q, a...)
	if err != nil {
		log.Errorw("failed to insert project", "err", err)
		return err
	}

	return nil
}
