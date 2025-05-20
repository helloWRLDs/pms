package participantdata

import (
	"context"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

func (r *Repository) List(ctx context.Context, filter *dto.ParticipantFilter) (res list.List[Participant], err error) {
	log := r.log.With(
		zap.String("func", "List"),
		zap.Any("filter", filter),
	)
	log.Debug("List called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("list"),
			errs.WithObject("participants"),
		)
	}()

	builder := r.gen.
		Select("p.*").
		From("\"Participant\" p")

	if filter.UserId != "" {
		builder = builder.Where(sq.Eq{"p.user_id": filter.UserId})
	}
	if filter.Role != "" {
		builder = builder.Where(sq.Eq{"p.role": filter.Role})
	}
	if filter.CompanyId != "" {
		builder = builder.Where(sq.Eq{"p.company_id": filter.CompanyId})
	}
	if filter.DateFrom != "" {
		builder = builder.Where(sq.GtOrEq{"p.created_at": filter.DateFrom})
	}
	if filter.DateTo != "" {
		builder = builder.Where(sq.LtOrEq{"p.created_at": filter.DateTo})
	}

	{ // build pagination info
		filter.Page = utils.If(filter.Page <= 0, 1, filter.Page)
		filter.PerPage = utils.If(filter.PerPage <= 0, 10, filter.PerPage)

		countQuery, countArgs, _ := builder.ToSql()
		if err := r.DB.QueryRowx(strings.ReplaceAll(countQuery, "SELECT p.*", "SELECT COUNT(*)"), countArgs...).Scan(&res.TotalItems); err != nil {
			log.Errorw("failed to count users", "err", err)
			return list.List[Participant]{}, err
		}
		res.Page = int(filter.Page)
		res.PerPage = int(filter.PerPage)
		res.TotalPages = int((int32(res.TotalItems) + filter.PerPage - 1) / filter.PerPage)
	}
	if filter.OrderBy != "" {
		builder = builder.OrderBy(filter.OrderBy + " " + filter.OrderDirection)
	} else {
		builder = builder.OrderBy("p.created_at DESC")
	}

	builder = builder.Limit(uint64(filter.PerPage)).Offset(uint64((filter.Page - 1) * filter.PerPage))

	query, args, _ := builder.ToSql()
	log.Debugw("query built", "query", query, "args", args)

	rows, err := r.DB.QueryxContext(ctx, query, args...)
	if err != nil {
		log.Errorw("failed to fetch participant", "err", err)
		return list.List[Participant]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var participant Participant
		if err := rows.StructScan(&participant); err != nil {
			log.Errorw("failed to scan participant", "err", err)
			return list.List[Participant]{}, err
		}
		res.Items = append(res.Items, participant)
	}

	return res, nil

}

func (r *Repository) Exists(ctx context.Context, userID, companyID string) (exists bool) {
	query, args, _ := r.gen.
		Select("COUNT(*) > 0").
		From(r.tableName).
		Where(sq.Eq{"company_id": companyID}).
		Where(sq.Eq{"user_id": userID}).
		ToSql()

	err := r.DB.QueryRow(query, args...).Scan(&exists)
	return err == nil && exists
}

func (r *Repository) Get(ctx context.Context, userID, companyID string) (participant Participant, err error) {
	log := r.log.With(
		zap.String("func", "Get"),
		zap.String("user_id", userID),
		zap.String("company_id", companyID),
	)
	log.Debug("Get called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("get"),
			errs.WithObject("participant"),
			errs.WithField("user_id", userID),
		)
	}()

	q, a, _ := r.gen.
		Select("*").
		From(r.tableName).
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"company_id": companyID}).
		ToSql()

	if err := r.DB.QueryRowx(q, a...).StructScan(&participant); err != nil {
		log.Errorw("failed to fetch participant", "err", err)
		return participant, err
	}
	return participant, nil
}

func (r *Repository) GetByUserID(ctx context.Context, userID string) (participations []Participant, err error) {
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
		var p Participant
		if err = rows.StructScan(&p); err != nil {
			return nil, err
		}
		participations = append(participations, p)
	}
	return participations, nil
}
