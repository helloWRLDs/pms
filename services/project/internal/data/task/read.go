package taskdata

import (
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

func (r *Repository) Exists(ctx context.Context, field string, value interface{}) (exists bool) {
	log := r.log.With(
		zap.String("func", "Exists"),
		zap.Any("condition", fmt.Sprintf("%s: %v", field, value)),
	)
	log.Debug("Exists called")

	q := `SELECT EXISTS(SELECT id FROM "Task" WHERE %s = ?)`

	if err := r.DB.QueryRowx(fmt.Sprintf(q, field), value).Scan(&exists); err != nil {
		return false
	}
	return exists
}

func (r *Repository) GetByID(ctx context.Context, id string) (task Task, err error) {
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
		return Task{}, err
	}
	return task, nil
}

func (r *Repository) Count(ctx context.Context, filter *dto.TaskFilter) (count int64) {
	builder := r.gen.
		Select("COUNT(*)").
		From("\"Task\" t")

	if filter.AssigneeId != "" {
		builder = builder.
			LeftJoin("\"TaskAssignment\" a ON a.task_id = t.id").
			Where(sq.Eq{"a.user_id": filter.AssigneeId})
	}
	if filter.ProjectName != "" {
		builder = builder.
			LeftJoin("\"Project\" p ON p.id = t.project_id").
			Where(sq.Eq{"p.name": filter.ProjectName})
	}
	if filter.SprintName != "" {
		builder = builder.
			LeftJoin("\"Sprint\" s ON s.id = t.sprint_id").
			Where(sq.Eq{"s.title": filter.SprintName})
	}

	if filter.Code != "" {
		builder = builder.Where(sq.Eq{"t.code": filter.Code})
	}
	if filter.DateFrom != "" {
		builder = builder.Where(sq.GtOrEq{"t.created_at": filter.DateFrom})
	}
	if filter.DateTo != "" {
		builder = builder.Where(sq.LtOrEq{"t.created_at": filter.DateTo})
	}

	if filter.Title != "" {
		builder = builder.Where(sq.Eq{"t.title": filter.Title})
	}
	if filter.Priority != 0 {
		builder = builder.Where(sq.Eq{"t.priority": filter.Priority})
	}
	if filter.Status != "" {
		builder = builder.Where(sq.Eq{"t.status": filter.Status})
	}
	if filter.ProjectId != "" {
		builder = builder.Where(sq.Eq{"t.project_id": filter.ProjectId})
	}
	if filter.SprintId != "" {
		builder = builder.Where(sq.Eq{"t.sprint_id": filter.SprintId})
	}

	query, args, _ := builder.ToSql()
	if err := r.DB.QueryRowx(query, args...).Scan(&count); err != nil {
		return 0
	}
	return count
}

func (r *Repository) List(ctx context.Context, filter *dto.TaskFilter) (res list.List[Task], err error) {
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

	builder := r.gen.
		Select("t.*").
		From("\"Task\" t")

	if filter.AssigneeId != "" {
		builder = builder.
			LeftJoin("\"TaskAssignment\" a ON a.task_id = t.id").
			Where(sq.Eq{"a.user_id": filter.AssigneeId})
	}
	if filter.ProjectName != "" {
		builder = builder.
			LeftJoin("\"Project\" p ON p.id = t.project_id").
			Where(sq.ILike{"p.name": "%" + filter.ProjectName + "%"})
	}
	if filter.SprintName != "" {
		builder = builder.
			LeftJoin("\"Sprint\" s ON s.id = t.sprint_id").
			Where(sq.ILike{"s.title": "%" + filter.SprintName + "%"})
	}

	if filter.DateFrom != "" {
		builder = builder.Where(sq.GtOrEq{"t.created_at": filter.DateFrom})
	}
	if filter.DateTo != "" {
		builder = builder.Where(sq.LtOrEq{"t.created_at": filter.DateTo})
	}

	if filter.Title != "" {
		builder = builder.Where(sq.ILike{"t.title": "%" + filter.Title + "%"})
	}
	if filter.Priority != 0 {
		builder = builder.Where(sq.Eq{"t.priority": filter.Priority})
	}
	if filter.Status != "" {
		builder = builder.Where(sq.Eq{"t.status": filter.Status})
	}
	if filter.ProjectId != "" {
		builder = builder.Where(sq.Eq{"t.project_id": filter.ProjectId})
	}
	if filter.SprintId != "" {
		builder = builder.Where(sq.Eq{"t.sprint_id": filter.SprintId})
	}
	if filter.Type != "" {
		builder = builder.Where(sq.Eq{"t.type": filter.Type})
	}

	{
		filter.Page = utils.If(filter.Page <= 0, 1, filter.Page)
		filter.PerPage = utils.If(filter.PerPage <= 0, 10, filter.PerPage)
		var totalItems int64
		countQuery, countArgs, _ := builder.ToSql()
		log.Debugw("count query built", "q", countQuery, "args", countArgs)
		if err := r.DB.QueryRowx(strings.ReplaceAll(countQuery, "SELECT t.*", "SELECT COUNT(*)"), countArgs...).Scan(&totalItems); err != nil {
			log.Errorw("failed to count tasks", "err", err)
			return list.List[Task]{}, err
		}
		res.TotalItems = int(totalItems)
		res.Page = int(filter.Page)
		res.PerPage = int(filter.PerPage)
		res.TotalPages = (res.TotalItems + int(filter.PerPage) - 1) / int(filter.PerPage)
	}

	if filter.OrderBy != "" {
		builder = builder.OrderBy(filter.OrderBy + " " + filter.OrderDirection)
	} else {
		builder = builder.OrderBy("t.created_at DESC")
	}

	builder = builder.Limit(uint64(filter.PerPage)).Offset(uint64((filter.Page - 1) * filter.PerPage))

	query, args, _ := builder.ToSql()
	log.Debugw("query built", "query", query, "args", args)

	rows, err := r.DB.Queryx(query, args...)
	if err != nil {
		log.Errorw("failed to fetch companies", "err", err)
		return list.List[Task]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		if err := rows.StructScan(&task); err != nil {
			log.Errorw("failed to scan Task", "err", err)
			return list.List[Task]{}, err
		}
		res.Items = append(res.Items, task)
	}

	return res, nil
}
