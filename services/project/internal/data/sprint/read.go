package sprintdata

import (
	"context"

	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.project/internal/data/models"
)

func (r *Repository) GetByID(ctx context.Context, id string) (sprint models.Sprint, err error) {
	log := r.log.With(
		zap.String("func", "GetByID"),
		zap.Any("id", id),
	)
	log.Debug("GetByID called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("get"),
			errs.WithField("id", id),
		)
	}()

	q, a, _ := r.gen.
		Select("*").
		From(r.tableName).
		Where(squirrel.Eq{"id": id}).ToSql()

	if err = r.DB.QueryRowx(q, a...).StructScan(&sprint); err != nil {
		log.Errorw("failed to fetch sprint", "err", err)
		return models.Sprint{}, err
	}

	return sprint, err
}
