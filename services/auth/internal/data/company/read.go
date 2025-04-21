package companydata

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

// p - Participant
func (r *Repository) List(ctx context.Context, filter list.Filters) (res list.List[entity.Company], err error) {
	log := r.log.With(
		zap.String("func", "ListCompanies"),
		zap.Any("filter", filter),
	)
	log.Debug("ListCompanies called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("list"),
			errs.WithObject("comapanies"),
		)
	}()

	builder := r.gen.
		Select("c.*").
		From("Company c").
		LeftJoin("Participant p ON c.id = p.company_id")

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

	rows, err := r.DB.Queryx(query, args...)
	if err != nil {
		log.Errorw("failed to fetch companies", "err", err)
		return list.List[entity.Company]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var company entity.Company
		if err := rows.StructScan(&company); err != nil {
			log.Errorw("failed to scan company", "err", err)
			return list.List[entity.Company]{}, err
		}
		res.Items = append(res.Items, company)
	}

	return res, nil
}

func (r *Repository) Count(ctx context.Context, filter list.Filters) (count int64) {
	log := r.log.With(zap.String("func", "Count"))
	log.Debug("Count called")

	builder := r.gen.
		Select("COUNT(*)").
		From("Company c").
		LeftJoin("Participant p ON c.id = p.company_id")

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

	q, a, _ := builder.ToSql()

	r.DB.QueryRow(q, a...).Scan(&count)
	return count
}

func (r *Repository) Exists(ctx context.Context, id string) (exists bool) {
	query, args, _ := r.gen.
		Select("COUNT(*) > 0").
		From("Company").
		Where(sq.Eq{"id": id}).
		ToSql()

	err := r.DB.QueryRow(query, args...).Scan(&exists)
	return err == nil && exists
}

func (r *Repository) GetByID(ctx context.Context, companyID string) (company entity.Company, err error) {
	log := r.log.With(
		zap.String("func", "GetByID"),
		zap.String("id", companyID),
	)
	log.Debug("GetCompanyByID called")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithField("id", companyID), errs.WithOperation("get"))
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err = transaction.Start(ctx, r.DB)
		if err != nil {
			return
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	query, args, _ := r.gen.
		Select("*").
		From("Company").
		Where(sq.Eq{"id": companyID}).
		ToSql()

	if err := tx.QueryRowx(query, args...).StructScan(&company); err != nil {
		log.Warnw("failed to fetch company by ID", "err", err)
		return company, err
	}
	return company, nil
}
