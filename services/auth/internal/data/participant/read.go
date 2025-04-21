package participantdata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.auth/internal/entity"
	"pms.pkg/errs"
)

// func (r *Repository) GetParticipantByUserID(ctx context.Context, userID string) (participant *entity.Participant, err error) {
// 	log := r.log.With(
// 		zap.String("func", "GetParticipantByUserID"),
// 		zap.String("user_id", userID),
// 	)
// 	log.Debug("GetParticipantByUserID called")

// 	defer func() {
// 		err = r.errctx.MapSQL(err, errs.WithOperation("get"), errs.WithField("id", userID))
// 	}()

// 	tx := transaction.Retrieve(ctx)
// 	if tx == nil {
// 		log.Debug("tx not found")
// 		ctx, err = transaction.Start(ctx, r.DB)
// 		if err != nil {
// 			return nil, err
// 		}
// 		tx = transaction.Retrieve(ctx)
// 		defer func() {
// 			transaction.End(ctx, err)
// 		}()
// 	}

// 	q, args, _ := r.gen.Select("*").From("Participant").Where("user_id = ?", userID).ToSql()

// }

func (r *Repository) GetByUserID(ctx context.Context, userID string) (participations []entity.Participant, err error) {
	log := r.log.With(
		zap.String("func", "GetByUserID"),
		zap.String("user_id", userID),
	)
	log.Debug("GetByUserID called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("get"),
			errs.WithObject("participant"),
			errs.WithField("user_id", userID),
		)
	}()

	q, a, _ := r.gen.Select("*").From(r.tableName).Where(sq.Eq{"user_id": userID}).ToSql()
	rows, err := r.DB.Queryx(q, a...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p entity.Participant
		if err = rows.StructScan(&p); err != nil {
			return nil, err
		}
		participations = append(participations, p)
	}
	return participations, nil
}
