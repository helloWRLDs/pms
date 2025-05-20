package taskcommentdata

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

func (r *Repository) List(ctx context.Context, filter *dto.TaskCommentFilter) (res list.List[TaskComment], err error) {
	log := r.log.Named("List").With(
		zap.Any("filter", filter),
	)
	log.Debug("List called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("list"),
		)
	}()

	builder := r.gen.
		Select("ta.*").
		From("\"TaskComment\" ta")

	if filter.TaskId != "" {
		builder = builder.Where(sq.Eq{"ta.task_id": filter.TaskId})
	}
	if filter.UserId != "" {
		builder = builder.Where(sq.Eq{"ta.user_id": filter.UserId})
	}

	if filter.DateFrom != "" {
		builder = builder.Where(sq.GtOrEq{"ta.created_at": filter.DateFrom})
	}
	if filter.DateTo != "" {
		builder = builder.Where(sq.LtOrEq{"ta.created_at": filter.DateTo})
	}

	{
		filter.Page = utils.If(filter.Page <= 0, 1, filter.Page)
		filter.PerPage = utils.If(filter.PerPage <= 0, 10, filter.PerPage)
		var totalItems int64
		countQuery, countArgs, _ := builder.ToSql()
		log.Debugw("count query built", "q", countQuery, "args", countArgs)
		if err := r.DB.QueryRowx(strings.ReplaceAll(countQuery, "SELECT ta.*", "SELECT COUNT(*)"), countArgs...).Scan(&totalItems); err != nil {
			log.Errorw("failed to count task comments", "err", err)
			return list.List[TaskComment]{}, err
		}
		res.TotalItems = int(totalItems)
		res.Page = int(filter.Page)
		res.PerPage = int(filter.PerPage)
		res.TotalPages = (res.TotalItems + int(filter.PerPage) - 1) / int(filter.PerPage)
	}

	if filter.OrderBy != "" {
		builder = builder.OrderBy("ta.", filter.OrderBy+" "+filter.OrderDirection)
	} else {
		builder = builder.OrderBy("ta.created_at DESC")
	}

	builder = builder.Limit(uint64(filter.PerPage)).Offset(uint64((filter.Page - 1) * filter.PerPage))
	query, args, _ := builder.ToSql()
	log.Debugw("query built", "query", query, "args", args)

	rows, err := r.DB.Queryx(query, args...)
	if err != nil {
		log.Errorw("failed to fetch task comments", "err", err)
		return list.List[TaskComment]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment TaskComment
		if err := rows.StructScan(&comment); err != nil {
			log.Errorw("failed to scan task comment", "err", err)
			return list.List[TaskComment]{}, err
		}
		res.Items = append(res.Items, comment)
	}

	return res, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (comment TaskComment, err error) {
	log := r.log.Named("GetByID").With(
		zap.String("id", id),
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
		Where(sq.Eq{"id": id}).
		ToSql()

	if err = r.DB.QueryRowx(q, a...).StructScan(&comment); err != nil {
		log.Errorw("failed to fetch task comment", "err", err)
		return comment, err
	}
	return comment, nil
}
