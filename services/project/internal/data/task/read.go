package taskdata

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

func (r *Repository) Exists(ctx context.Context, field string, value interface{}) (exists bool) {
	log := r.log.With(
		zap.String("func", "Exists"),
		zap.Any("condition", fmt.Sprintf("%s: %v", field, value)),
	)
	log.Debug("Exists called")

	q := `SELECT EXISTS(SELECT id FROM Task WHERE %s = ?)`

	if err := r.DB.QueryRowx(fmt.Sprintf(q, field), value).Scan(&exists); err != nil {
		return false
	}
	return exists
}

func (r *Repository) GetByID(ctx context.Context, id string) (task models.Task, err error) {
	log := r.log.With(
		zap.String("func", "GetByID"),
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
		Where(sq.Eq{"id": id}).ToSql()

	if err = r.DB.QueryRowx(q, a...).StructScan(&task); err != nil {
		log.Errorw("failed to scan task", "err", err)
		return models.Task{}, err
	}
	return task, nil
}

// p - project, ta - taskAssignment
func (r *Repository) Count(ctx context.Context, filter list.Filters) (count int64) {
	builder := r.gen.
		Select("COUNT(*)").
		From("Task t").
		LeftJoin("TaskAssignment ta ON ta.task_id = t.id").
		LeftJoin("Project p ON p.id = t.project_id")

	if filter.Date.From != "" {
		builder = builder.Where(sq.GtOrEq{"t.created_at": filter.Date.From})
	}
	if filter.Date.To != "" {
		builder = builder.Where(sq.LtOrEq{"t.created_at": filter.Date.To})
	}

	if filter.Order.By != "" {
		builder = builder.OrderBy(filter.Order.String())
	} else {
		builder = builder.OrderBy("t.created_at DESC")
	}
	for k, v := range filter.Fields {
		builder = builder.Where(sq.Eq{k: v})
	}
	for k, v := range filter.InFields {
		builder = builder.Where(fmt.Sprintf("%s IN (%v)", k, strings.Join(v, ",")))
	}
	query, args, _ := builder.ToSql()
	if err := r.DB.QueryRowx(query, args...).Scan(&count); err != nil {
		return 0
	}
	return count
}

func (r *Repository) List(ctx context.Context, filter list.Filters) (res list.List[models.Task], err error) {
	log := r.log.With(
		zap.String("func", "List"),
		zap.String("filter", filter.String()),
	)
	log.Debug("List called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("list"),
			errs.WithObject("tasks"),
		)
	}()

	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.PerPage <= 0 {
		filter.PerPage = 10
	}

	builder := r.gen.
		Select("t.*").
		From("Task t").
		LeftJoin("TaskAssignment a ON a.task_id = t.id")

	if filter.Date.From != "" {
		builder = builder.Where(sq.GtOrEq{"t.created_at": filter.Date.From})
	}
	if filter.Date.To != "" {
		builder = builder.Where(sq.LtOrEq{"t.created_at": filter.Date.To})
	}

	if filter.Order.By != "" {
		builder = builder.OrderBy(filter.Order.String())
	} else {
		builder = builder.OrderBy("t.created_at DESC")
	}
	for k, v := range filter.Fields {
		builder = builder.Where(sq.Eq{k: v})
	}
	for k, v := range filter.InFields {
		builder = builder.Where(fmt.Sprintf("t.%s IN (%v)", k, v))
	}

	{
		countQuery, countArgs, _ := builder.ToSql()
		if err := r.DB.QueryRowx(strings.ReplaceAll(countQuery, "SELECT t.*", "SELECT COUNT(*)"), countArgs...).Scan(&res.TotalItems); err != nil {
			log.Errorw("failed to count tasks", "err", err)
			return list.List[models.Task]{}, err
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
		log.Errorw("failed to fetch companies", "err", err)
		return list.List[models.Task]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		if err := rows.StructScan(&task); err != nil {
			log.Errorw("failed to scan Task", "err", err)
			return list.List[models.Task]{}, err
		}
		res.Items = append(res.Items, task)
	}

	return res, nil
}
