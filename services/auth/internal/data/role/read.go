package roledata

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.auth/internal/entity"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/type/list"
)

func (r *Repository) Count(ctx context.Context, filter list.Filters) (count int64) {
	log := r.log.With(
		zap.String("func", "Count"),
		zap.Any("filters", filter),
	)
	log.Debug("Count called")

	builder := r.gen.
		Select("COUNT(*)").
		From("Role r").
		LeftJoin("Company c ON c.id = r.company_id")

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

func (r *Repository) List(ctx context.Context, filter list.Filters) (res list.List[entity.Role], err error) {
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

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err := transaction.Start(ctx, r.DB)
		if err != nil {
			return list.List[entity.Role]{}, err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	builder := r.gen.
		Select("r.*").
		From("Role r").
		LeftJoin("Company c ON c.id = r.company_id")

	if filter.Date.From != "" {
		builder = builder.Where(sq.GtOrEq{"c.created_at": filter.Date.From})
	}
	if filter.Date.To != "" {
		builder = builder.Where(sq.LtOrEq{"c.created_at": filter.Date.To})
	}

	if filter.Order.By != "" {
		builder = builder.OrderBy(filter.Order.String())
	} else {
		builder = builder.OrderBy("c.created_at DESC")
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
		builder = builder.Where(fmt.Sprintf("c.%s IN (%v)", k, v))
	}
	res.TotalItems = int(r.Count(ctx, filter))
	res.Page = filter.Page
	res.PerPage = filter.PerPage
	res.TotalPages = (res.TotalItems + filter.PerPage - 1) / filter.PerPage

	builder = builder.Limit(uint64(filter.PerPage)).Offset(uint64((filter.Page - 1) * filter.PerPage))

	query, args, _ := builder.ToSql()
	log.Debugw("query built", "query", query, "args", args)

	rows, err := tx.Queryx(query, args...)
	if err != nil {
		log.Errorw("failed to fetch roles", "err", err)
		return list.List[entity.Role]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var company entity.Role
		if err := rows.StructScan(&company); err != nil {
			log.Errorw("failed to scan role", "err", err)
			return list.List[entity.Role]{}, err
		}
		res.Items = append(res.Items, company)
	}

	return res, nil
}

func (r *Repository) GetByName(ctx context.Context, name string) (role entity.Role, err error) {
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

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err := transaction.Start(ctx, r.DB)
		if err != nil {
			return entity.Role{}, err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	q, a, _ := r.gen.Select("*").From("Role").Where(sq.Eq{"name": name}).ToSql()
	if err = tx.QueryRowx(q, a...).StructScan(&role); err != nil {
		log.Errorw("failed to get role", "err", err)
		return entity.Role{}, err
	}
	return role, nil
}
