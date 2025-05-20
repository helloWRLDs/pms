package sprintdata

import (
	"context"
	"strings"

	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

func (r *Repository) GetByID(ctx context.Context, id string) (sprint Sprint, err error) {
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
		return Sprint{}, err
	}

	return sprint, err
}

func (r *Repository) List(ctx context.Context, filter *dto.SprintFilter) (res list.List[Sprint], err error) {
	log := r.log.With(
		zap.String("func", "List"),
		zap.Any("filter", filter),
	)
	log.Debug("List called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("list"),
			errs.WithObject("sprints"),
		)
	}()

	builder := r.gen.
		Select("s.*").
		From("\"Sprint\" s")

	if filter.ProjectName != "" {
		builder = builder.LeftJoin("\"Project\" p ON p.id = s.project_id").Where(squirrel.Eq{"p.name": filter.ProjectName})
	}
	if filter.Title != "" {
		builder = builder.Where(squirrel.Eq{"s.title": filter.Title})
	}
	if filter.Description != "" {
		builder = builder.Where(squirrel.Eq{"s.description": filter.Description})
	}
	if filter.ProjectId != "" {
		builder = builder.Where(squirrel.Eq{"s.project_id": filter.ProjectId})
	}

	{
		filter.Page = utils.If(filter.Page <= 0, 1, filter.Page)
		filter.PerPage = utils.If(filter.PerPage <= 0, 10, filter.PerPage)
		var totalItems int64
		countQuery, countArgs, _ := builder.ToSql()
		if err := r.DB.QueryRowx(strings.ReplaceAll(countQuery, "SELECT s.*", "SELECT COUNT(*)"), countArgs...).Scan(&totalItems); err != nil {
			log.Errorw("failed to count sprints", "err", err)
			return list.List[Sprint]{}, err
		}
		res.TotalItems = int(totalItems)
		res.Page = int(filter.Page)
		res.PerPage = int(filter.PerPage)
		res.TotalPages = (res.TotalItems + int(filter.PerPage) - 1) / int(filter.PerPage)
	}

	if filter.DateFrom != "" {
		builder = builder.Where(squirrel.GtOrEq{"s.created_at": filter.DateFrom})
	}
	if filter.DateTo != "" {
		builder = builder.Where(squirrel.LtOrEq{"s.created_at": filter.DateTo})
	}
	if filter.OrderBy != "" {
		builder = builder.OrderBy(filter.OrderBy + " " + filter.OrderDirection)
	} else {
		builder = builder.OrderBy("s.created_at DESC")
	}

	builder = builder.Limit(uint64(filter.PerPage)).Offset(uint64((filter.Page - 1) * filter.PerPage))

	query, args, _ := builder.ToSql()
	log.Debugw("query built", "query", query, "args", args)

	rows, err := r.DB.Queryx(query, args...)
	if err != nil {
		log.Errorw("failed to fetch sprints", "err", err)
		return list.List[Sprint]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var sprint Sprint
		if err := rows.StructScan(&sprint); err != nil {
			log.Errorw("failed to scan sprint", "err", err)
			return list.List[Sprint]{}, err
		}
		res.Items = append(res.Items, sprint)
	}

	return res, nil
}
