package projectdata

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/type/list"
	"pms.project/internal/entity"
)

func (r *Repository) GetByID(ctx context.Context, id string) (project entity.Project, err error) {
	log := r.log.With(
		zap.String("func", "GetByID"),
		zap.String("id", id),
	)
	log.Debug("GetByID called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("get"),
			errs.WithObject("project"),
			errs.WithField("id", id),
		)
	}()

	q, a, _ := r.gen.Select("*").From(r.tableName).Where(sq.Eq{"id": id}).ToSql()
	if err = r.DB.QueryRowx(q, a...).StructScan(&project); err != nil {
		log.Errorw("failed to scan project", "err", err)
		return entity.Project{}, err
	}
	return project, nil
}

func (r *Repository) Count(ctx context.Context, filter list.Filters) (count int64) {
	log := r.log.With(zap.String("func", "Count"))
	log.Debug("Count called")

	builder := r.gen.
		Select("COUNT(*)").
		From("Project p")
		// LeftJoin("Participant p ON c.id = p.company_id")

	if filter.Date.From != "" {
		builder = builder.Where(sq.GtOrEq{"p.created_at": filter.Date.From})
	}
	if filter.Date.To != "" {
		builder = builder.Where(sq.LtOrEq{"p.created_at": filter.Date.To})
	}

	if filter.Order.By != "" {
		builder = builder.OrderBy(filter.Order.String())
	} else {
		builder = builder.OrderBy("p.created_at DESC")
	}
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.PerPage <= 0 {
		filter.PerPage = 10
	}
	for k, v := range filter.Fields {
		builder = builder.Where(sq.Eq{k: v})
	}
	for k, v := range filter.InFields {
		builder = builder.Where(fmt.Sprintf("p.%s IN (%v)", k, v))
	}

	q, a, _ := builder.ToSql()

	r.DB.QueryRow(q, a...).Scan(&count)
	return count
}

func (r *Repository) List(ctx context.Context, filter list.Filters) (res list.List[entity.Project], err error) {
	log := r.log.With(
		zap.String("func", "List"),
		zap.String("filter", filter.String()),
	)
	log.Debug("List called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("list"),
			errs.WithObject("projects"),
		)
	}()

	builder := r.gen.
		Select("p.*").
		From("Project p")
		// LeftJoin("Participant p ON c.id = p.company_id")

	if filter.Date.From != "" {
		builder = builder.Where(sq.GtOrEq{"p.created_at": filter.Date.From})
	}
	if filter.Date.To != "" {
		builder = builder.Where(sq.LtOrEq{"p.created_at": filter.Date.To})
	}

	if filter.Order.By != "" {
		builder = builder.OrderBy(filter.Order.String())
	} else {
		builder = builder.OrderBy("p.created_at DESC")
	}
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.PerPage <= 0 {
		filter.PerPage = 10
	}
	for k, v := range filter.Fields {
		builder = builder.Where(sq.Eq{k: v})
	}
	for k, v := range filter.InFields {
		builder = builder.Where(fmt.Sprintf("p.%s IN (%v)", k, v))
	}

	res.TotalItems = int(r.Count(ctx, filter))
	res.Page = filter.Page
	res.PerPage = filter.PerPage
	res.TotalPages = (res.TotalItems + filter.PerPage - 1) / filter.PerPage

	builder = builder.Limit(uint64(filter.PerPage)).Offset(uint64((filter.Page - 1) * filter.PerPage))

	query, args, _ := builder.ToSql()
	log.Debugw("query built", "query", query, "args", args)

	rows, err := r.DB.Queryx(query, args...)
	if err != nil {
		log.Errorw("failed to fetch companies", "err", err)
		return list.List[entity.Project]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var company entity.Project
		if err := rows.StructScan(&company); err != nil {
			log.Errorw("failed to scan company", "err", err)
			return list.List[entity.Project]{}, err
		}
		res.Items = append(res.Items, company)
	}

	return res, nil
}
