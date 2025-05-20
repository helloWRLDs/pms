package companydata

import (
	"context"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

// p - Participant
func (r *Repository) List(ctx context.Context, filter *dto.CompanyFilter) (res list.List[Company], err error) {
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
		From("\"Company\" c")

	if filter.UserId != "" || filter.Role != "" {
		builder = builder.LeftJoin("\"Participant\" p ON c.id = p.company_id")
		if filter.UserId != "" {
			builder = builder.Where(sq.Eq{"p.user_id": filter.UserId})
		}
		if filter.Role != "" {
			builder = builder.Where(sq.Eq{"p.role": filter.Role})
		}
	}

	if filter.DateFrom != "" {
		builder = builder.Where(sq.GtOrEq{"c.created_at": filter.DateFrom})
	}
	if filter.DateTo != "" {
		builder = builder.Where(sq.LtOrEq{"c.created_at": filter.DateTo})
	}

	if filter.CodeName != "" {
		builder = builder.Where(sq.Eq{"c.codename": filter.CodeName})
	}
	if filter.CompanyName != "" {
		builder = builder.Where(sq.Eq{"c.name": filter.CompanyName})
	}
	if filter.CompanyId != "" {
		builder = builder.Where(sq.Eq{"c.id": filter.CompanyId})
	}

	{
		filter.Page = utils.If(filter.Page <= 0, 1, filter.Page)
		filter.PerPage = utils.If(filter.PerPage <= 0, 10, filter.PerPage)

		countQuery, countArgs, _ := builder.ToSql()
		if err := r.DB.QueryRowx(strings.ReplaceAll(countQuery, "SELECT c.*", "SELECT COUNT(*)"), countArgs...).Scan(&res.TotalItems); err != nil {
			log.Errorw("failed to count companies", "err", err)
			return list.List[Company]{}, err
		}
		res.Page = int(filter.Page)
		res.PerPage = int(filter.PerPage)
		res.TotalPages = int((int32(res.TotalItems) + filter.PerPage - 1) / filter.PerPage)
	}

	if filter.OrderBy != "" {
		builder = builder.OrderBy(filter.OrderBy + " " + filter.OrderDirection)
	} else {
		builder = builder.OrderBy("c.created_at DESC")
	}

	builder = builder.Limit(uint64(filter.PerPage)).Offset(uint64((filter.Page - 1) * filter.PerPage))

	query, args, _ := builder.ToSql()
	log.Debugw("query built", "query", query, "args", args)

	rows, err := r.DB.Queryx(query, args...)
	if err != nil {
		log.Errorw("failed to fetch companies", "err", err)
		return list.List[Company]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var company Company
		if err := rows.StructScan(&company); err != nil {
			log.Errorw("failed to scan company", "err", err)
			return list.List[Company]{}, err
		}
		res.Items = append(res.Items, company)
	}

	return res, nil
}

func (r *Repository) Count(ctx context.Context, filter *dto.CompanyFilter) (count int64) {
	log := r.log.With(zap.String("func", "Count"))
	log.Debug("Count called")

	builder := r.gen.
		Select("COUNT(*)").
		From("\"Company\" c")

	if filter.UserId != "" || filter.Role != "" {
		builder = builder.LeftJoin("\"Participant\" p ON c.id = p.company_id")
		if filter.UserId != "" {
			builder = builder.Where(sq.Eq{"p.user_id": filter.UserId})
		}
		if filter.Role != "" {
			builder = builder.Where(sq.Eq{"p.role": filter.Role})
		}
	}

	if filter.DateFrom != "" {
		builder = builder.Where(sq.GtOrEq{"c.created_at": filter.DateFrom})
	}
	if filter.DateTo != "" {
		builder = builder.Where(sq.LtOrEq{"c.created_at": filter.DateTo})
	}

	if filter.CodeName != "" {
		builder = builder.Where(sq.Eq{"c.codename": filter.CodeName})
	}
	if filter.CompanyName != "" {
		builder = builder.Where(sq.Eq{"c.name": filter.CompanyName})
	}
	if filter.CompanyId != "" {
		builder = builder.Where(sq.Eq{"c.id": filter.CompanyId})
	}

	q, a, _ := builder.ToSql()

	r.DB.QueryRow(q, a...).Scan(&count)
	return count
}

func (r *Repository) Exists(ctx context.Context, field, value string) (exists bool) {
	query, args, _ := r.gen.
		Select("COUNT(*) > 0").
		From(r.tableName).
		Where(sq.Eq{field: value}).
		ToSql()

	err := r.DB.QueryRow(query, args...).Scan(&exists)
	return err == nil && exists
}

func (r *Repository) GetByID(ctx context.Context, companyID string) (company Company, err error) {
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
		From(r.tableName).
		Where(sq.Eq{"id": companyID}).
		ToSql()

	if err := tx.QueryRowx(query, args...).StructScan(&company); err != nil {
		log.Warnw("failed to fetch company by ID", "err", err)
		return company, err
	}
	return company, nil
}
