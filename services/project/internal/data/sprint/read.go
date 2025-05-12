package sprintdata

import (
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"

	"pms.pkg/errs"
	"pms.pkg/type/list"
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
		Where(sq.Eq{"id": id}).ToSql()

	if err = r.DB.QueryRowx(q, a...).StructScan(&sprint); err != nil {
		log.Errorw("failed to fetch sprint", "err", err)
		return models.Sprint{}, err
	}

	return sprint, err
}
func (r *Repository) Exists(ctx context.Context, field string, value interface{}) (exists bool) {
	log := r.log.With(
		zap.String("func", "Exists"),
		zap.Any("condition", fmt.Sprintf("%s: %v", field, value)),
	)
	log.Debug("Exists called")

	q := `SELECT EXISTS(SELECT id FROM Sprint WHERE %s = ?)`

	if err := r.DB.QueryRowx(fmt.Sprintf(q, field), value).Scan(&exists); err != nil {
		return false
	}
	return exists
}
func (r *Repository) Count(ctx context.Context, filter list.Filters) (count int64) {
	builder := r.gen.
		Select("COUNT(*)").
		From("Sprint s").
		LeftJoin("Project p on p.id = s.project_id")
	if filter.Date.From != "" {
		builder = builder.Where(sq.GtOrEq{"s.created_at": filter.Date.From})
	}
	if filter.Date.To != "" {
		builder = builder.Where(sq.LtOrEq{"s.created_at": filter.Date.To})
	}

	if filter.Order.By != "" {
		builder = builder.OrderBy(filter.Order.String())
	} else {
		builder = builder.OrderBy("s.created_at DESC")
	}
	for k, v := range filter.Fields {
		builder = builder.Where(sq.Eq{k: v})
	}
	for k, v := range filter.InFields {
		builder = builder.Where(fmt.Sprintf("%s in (%v)", k, strings.Join(v, ",")))
	}
	query, args, _ := builder.ToSql()
	if err := r.DB.QueryRowx(query, args...).Scan(&count); err != nil {
		return 0
	}
	return count
}
func (r *Repository) List(ctx context.Context, filter list.Filters) (res list.List[models.Sprint], err error) {
	log := r.log.With(
		zap.String("funct", "List"),
		zap.String("filter", filter.String()),
	)
	log.Debug("List called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("list"),
			errs.WithObject("sprints"))
	}()

	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.PerPage <= 0 {
		filter.PerPage = 10
	}

	builder := r.gen.
		Select("s.*").
		From("Sprint s").
		LeftJoin("Project p ON p.id = s.project_id")
	if filter.Date.From != "" {
		builder = builder.Where(sq.GtOrEq{"s.created_at": filter.Date.From})
	}
	if filter.Date.To != "" {
		builder = builder.Where(sq.LtOrEq{"s.created_at": filter.Date.To})
	}
	if filter.Order.By != "" {
		builder = builder.OrderBy(filter.Order.String())
	} else {
		builder = builder.OrderBy("s.created_at DESC")
	}
	for k, v := range filter.Fields {
		builder = builder.Where(sq.Eq{k: v})
	}
	for k, v := range filter.InFields {
		builder = builder.Where(fmt.Sprintf("s.%s IN (%v)", k, v))

	}
	{
		countQuery, countArgs, _ := builder.ToSql()
		if err := r.DB.QueryRowx(strings.ReplaceAll(countQuery, "SELECT s.*", "SELECT COUNT(*)"), countArgs...).Scan(&res.TotalItems); err != nil {
			log.Errorw("failed to count sprints", "err", err)
			return list.List[models.Sprint]{}, err
		}
		res.Page = filter.Page
		res.PerPage = filter.PerPage
		res.TotalPages = (res.TotalItems + filter.PerPage - 1) / filter.PerPage
	}
	builder = builder.Limit(uint64(filter.PerPage)).Offset(uint64((filter.Page - 1) * filter.PerPage))
	query, args, _ := builder.ToSql()
	log.Debugw("query built", "query", query, "args", args)
	rows, err := r.DB.Queryx(query, args...)
	if err != nil {
		log.Errorw("failed to fetch sprints", "err", err)
		return list.List[models.Sprint]{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var sprint models.Sprint
		if err := rows.StructScan(&sprint); err != nil {
			log.Errorw("failed to scan Sprint", "err", err)
			return list.List[models.Sprint]{}, err
		}
		res.Items = append(res.Items, sprint)
	}
	return res, nil
}
