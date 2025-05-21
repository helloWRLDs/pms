package projectdata

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

func (r *Repository) GetCode(ctx context.Context, projectID string) (code string) {
	log := r.log.Named("GetCode").With(
		zap.String("id", projectID),
	)
	log.Debug("GetCode called")

	nestedQ, _, _ := r.gen.
		Select("CAST(split_part(t.code, '-', 2) AS INT)").
		From("\"Task\" t").
		Where(sq.Eq{"t.project_id": "p.id"}).
		OrderBy("t.created_at DESC").
		Limit(1).
		ToSql()

	q, a, _ := r.gen.
		Select(fmt.Sprintf("CONCAT(COALESCE(p.code_prefix, p.codename), '-', COALESCE((%s), 0) + 1)", nestedQ)).
		From("\"Project\" p").
		Where(sq.Eq{"p.id": projectID}).
		ToSql()

	if err := r.DB.QueryRowx(q, a...).Scan(&code); err != nil {
		log.Errorw("failed to setup code")
		return ""
	}
	return code
}

func (r *Repository) GetCodeName(ctx context.Context, projectID string) (code string, err error) {
	log := r.log.With(
		zap.String("func", "GetCodeName"),
		zap.String("id", projectID),
	)
	log.Debug("GetCodeName called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("get code name"),
			errs.WithObject("project"),
			errs.WithField("id", projectID),
		)
	}()

	q, a, _ := r.gen.Select("COALESCE(code_prefix, codename)").From(r.tableName).ToSql()
	if err = r.DB.QueryRow(q, a...).Scan(&code); err != nil {
		log.Errorw("failed to scan project", "err", err)
		return "", err
	}
	return code, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (project Project, err error) {
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
		return Project{}, err
	}
	return project, nil
}

func (r *Repository) Count(ctx context.Context, filter *dto.ProjectFilter) (count int64) {
	log := r.log.With(zap.String("func", "Count"))
	log.Debug("Count called")

	builder := r.gen.
		Select("p.*").
		From("\"Project\" p")

	if filter.CompanyId != "" {
		builder = builder.Where(sq.Eq{"p.company_id": filter.CompanyId})
	}
	if filter.Title != "" {
		builder = builder.Where(sq.Eq{"p.title": filter.Title})
	}
	if filter.Description != "" {
		builder = builder.Where(sq.Eq{"p.description": filter.Description})
	}
	if filter.Status != "" {
		builder = builder.Where(sq.Eq{"p.status": filter.Status})
	}

	if filter.DateFrom != "" {
		builder = builder.Where(sq.GtOrEq{"p.created_at": filter.DateFrom})
	}
	if filter.DateTo != "" {
		builder = builder.Where(sq.LtOrEq{"p.created_at": filter.DateTo})
	}

	q, a, _ := builder.ToSql()

	r.DB.QueryRow(q, a...).Scan(&count)
	return count
}

func (r *Repository) List(ctx context.Context, filter *dto.ProjectFilter) (res list.List[Project], err error) {
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
		From("\"Project\" p")

	if filter.CompanyId != "" {
		builder = builder.Where(sq.Eq{"p.company_id": filter.CompanyId})
	}
	if filter.Title != "" {
		builder = builder.Where(sq.Eq{"p.title": filter.Title})
	}
	if filter.Description != "" {
		builder = builder.Where(sq.Eq{"p.description": filter.Description})
	}
	if filter.Status != "" {
		builder = builder.Where(sq.Eq{"p.status": filter.Status})
	}

	if filter.DateFrom != "" {
		builder = builder.Where(sq.GtOrEq{"p.created_at": filter.DateFrom})
	}
	if filter.DateTo != "" {
		builder = builder.Where(sq.LtOrEq{"p.created_at": filter.DateTo})
	}

	{
		filter.Page = utils.If(filter.Page <= 0, 1, filter.Page)
		filter.PerPage = utils.If(filter.PerPage <= 0, 10, filter.PerPage)

		countQuery, countArgs, _ := builder.ToSql()
		if err := r.DB.QueryRowx(strings.ReplaceAll(countQuery, "SELECT p.*", "SELECT COUNT(*)"), countArgs...).Scan(&res.TotalItems); err != nil {
			log.Errorw("failed to count projects", "err", err)
			return list.List[Project]{}, err
		}
		// res.TotalItems = int(totalItems)
		res.Page = int(filter.Page)
		res.PerPage = int(filter.PerPage)
		res.TotalPages = (res.TotalItems + int(filter.Page) - 1) / int(filter.PerPage)
	}

	if filter.OrderBy != "" {
		builder = builder.OrderBy(filter.OrderBy + " " + filter.OrderDirection)
	} else {
		builder = builder.OrderBy("p.created_at DESC")
	}

	builder = builder.Limit(uint64(filter.PerPage)).Offset(uint64((filter.Page - 1) * filter.PerPage))

	query, args, _ := builder.ToSql()
	log.Debugw("query built", "query", query, "args", args)

	rows, err := r.DB.Queryx(query, args...)
	if err != nil {
		log.Errorw("failed to fetch companies", "err", err)
		return list.List[Project]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var project Project
		if err := rows.StructScan(&project); err != nil {
			log.Errorw("failed to scan project", "err", err)
			return list.List[Project]{}, err
		}
		res.Items = append(res.Items, project)
	}

	return res, nil
}
