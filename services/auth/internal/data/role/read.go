package roledata

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

func (r *Repository) Count(ctx context.Context, filter list.Filters) (count int64) {
	log := r.log.With(
		zap.String("func", "Count"),
		zap.Any("filters", filter),
	)
	log.Debug("Count called")

	builder := r.gen.
		Select("COUNT(*)").
		From("\"Role\" r").
		LeftJoin("\"Company\" c ON c.id = r.company_id")

	if filter.Date.From != "" {
		builder = builder.Where(sq.GtOrEq{"r.created_at": filter.Date.From})
	}
	if filter.Date.To != "" {
		builder = builder.Where(sq.LtOrEq{"r.created_at": filter.Date.To})
	}

	if filter.Order.By != "" {
		builder = builder.OrderBy(filter.Order.String())
	} else {
		builder = builder.OrderBy("r.created_at DESC")
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
		builder = builder.Where(fmt.Sprintf("r.%s IN (%v)", k, v))
	}
	q, a, _ := builder.ToSql()
	log.Info("query ", q)

	r.DB.QueryRowx(q, a...).Scan(&count)
	return count
}

func (r *Repository) List(ctx context.Context, filter *dto.RoleFilter) (res list.List[Role], err error) {
	log := r.log.With(
		zap.String("func", "List"),
		zap.Any("filters", filter),
	)
	log.Debug("List called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("list"),
			errs.WithObject("role"),
		)
	}()

	builder := r.gen.
		Select("r.*").
		From("\"Role\" r")

	if filter.CompanyName != "" {
		builder = builder.LeftJoin("\"Company\" c ON c.id = r.company_id")

		builder = builder.Where(sq.Eq{"c.name": filter.CompanyName})
	}

	if filter.UserId != "" {
		builder = builder.LeftJoin("\"Participant\" p ON p.role = r.name")
		builder = builder.Where(sq.Eq{"p.user_id": filter.UserId})
	}

	if filter.DateFrom != "" {
		builder = builder.Where(sq.GtOrEq{"r.created_at": filter.DateFrom})
	}
	if filter.DateTo != "" {
		builder = builder.Where(sq.LtOrEq{"r.created_at": filter.DateTo})
	}

	if filter.WithDefault && filter.CompanyId != "" {
		builder = builder.Where(sq.Or{
			sq.Eq{"r.company_id": filter.CompanyId},
			sq.Eq{"r.company_id": nil},
		})
	} else if !filter.WithDefault && filter.CompanyId != "" {
		builder = builder.Where(sq.Eq{"r.company_id": filter.CompanyId})
	} else {
		builder = builder.Where(sq.Eq{"r.company_id": nil})
	}

	if filter.Name != "" {
		builder = builder.Where(sq.Eq{"r.name": filter.Name})
	}

	{ // build pagination info
		filter.Page = utils.If(filter.Page <= 0, 1, filter.Page)
		filter.PerPage = utils.If(filter.PerPage <= 0, 10, filter.PerPage)

		countQuery, countArgs, _ := builder.ToSql()
		if err := r.DB.QueryRowx(strings.ReplaceAll(countQuery, "SELECT r.*", "SELECT COUNT(*)"), countArgs...).Scan(&res.TotalItems); err != nil {
			log.Errorw("failed to count roles", "err", err)
			return list.List[Role]{}, err
		}
		res.Page = int(filter.Page)
		res.PerPage = int(filter.PerPage)
		res.TotalPages = (res.TotalItems + int(filter.PerPage) - 1) / int(filter.PerPage)
	}

	if filter.OrderBy != "" {
		builder = builder.OrderBy(filter.OrderBy)
	} else {
		builder = builder.OrderBy("r.created_at DESC")
	}
	builder = builder.Limit(uint64(filter.PerPage)).Offset(uint64((filter.Page - 1) * filter.PerPage))

	query, args, _ := builder.ToSql()
	log.Debugw("query built", "query", query, "args", args)

	rows, err := r.DB.Queryx(query, args...)
	if err != nil {
		log.Errorw("failed to fetch roles", "err", err)
		return list.List[Role]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var company Role
		if err := rows.StructScan(&company); err != nil {
			log.Errorw("failed to scan role", "err", err)
			return list.List[Role]{}, err
		}
		res.Items = append(res.Items, company)
	}

	return res, nil
}

func (r *Repository) GetByName(ctx context.Context, name string) (role Role, err error) {
	log := r.log.With(
		zap.String("func", "GetByName"),
		zap.String("name", name),
	)
	log.Debug("GetByName called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("get"),
			errs.WithObject("role"),
			errs.WithField("name", name),
		)
	}()

	q, a, _ := r.gen.Select("*").From("\"Role\"").Where(sq.Eq{"name": name}).ToSql()
	if err = r.DB.QueryRowx(q, a...).StructScan(&role); err != nil {
		log.Errorw("failed to get role", "err", err)
		return Role{}, err
	}
	return role, nil
}
