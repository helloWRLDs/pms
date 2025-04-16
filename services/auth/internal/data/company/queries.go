package companydata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"go.uber.org/zap"
	comp "pms.auth/internal/entity/company"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/type/list"
)

func (r *Repository) Create(ctx context.Context, newCompany comp.Company) (err error) {
	log := r.log.With(
		zap.String("func", "CreateCompany"),
		zap.String("name", newCompany.Name),
		zap.String("codename", newCompany.Codename),
	)
	log.Debug("CreateCompany called")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithOperation("create"), errs.WithField("codename", newCompany.Codename))
	}()
	tx := transaction.Retrieve(ctx)
	if tx == nil {
		log.Debug("tx not found")
		ctx, err = transaction.Start(ctx, r.DB)
		if err != nil {
			return err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	query, args, _ := r.gen.
		Insert("Company").
		Columns("name", "codename").
		Values(newCompany.Name, newCompany.Codename).
		ToSql()
	if _, err := tx.Exec(query, args...); err != nil {
		log.Errorw("failed to create company", "err", err)
		return err
	}
	log.Debug("Company created")
	return nil
}
func (r *Repository) GetByID(ctx context.Context, id string) (company comp.Company, err error) {
	log := r.log.With(
		zap.String("func", "GetByID"),
		zap.String("id", id),
	)
	log.Debug("GetCompanyByID called")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithField("id", id), errs.WithOperation("retrieve"))
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
		Where(sq.Eq{"id": id}).
		ToSql()

	if err := tx.QueryRowx(query, args...).StructScan(&company); err != nil {
		log.Warnw("failed to fetch company by ID", "err", err)
		return company, err
	}
	return company, nil
}
func (r *Repository) ExistsCompany(ctx context.Context, id string) (exists bool) {
	query, args, _ := r.gen.
		Select("COUNT(*) > 0").
		From("Company").
		Where(sq.Eq{"id": id}).
		ToSql()

	err := r.DB.QueryRow(query, args...).Scan(&exists)
	return err == nil && exists
}
func (r *Repository) Count(ctx context.Context, filter list.Filters) (int, error) {
	log := r.log.With(zap.String("func", "CountCompanies"))
	log.Debug("Count companies called")

	var count int

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err := transaction.Start(ctx, r.DB)
		if err != nil {
			return 0, err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	queryBuilder := r.gen.Select("COUNT(*)").From("Company")

	if filter.Created.CreatedFrom != "" {
		queryBuilder = queryBuilder.Where(sq.GtOrEq{"created_at": filter.Created.CreatedFrom})
	}
	if filter.Created.CreatedTo != "" {
		queryBuilder = queryBuilder.Where(sq.LtOrEq{"created_at": filter.Created.CreatedTo})
	}

	query, args, _ := queryBuilder.ToSql()

	err := tx.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *Repository) ListCompanies(ctx context.Context, filter list.Filters) (list.List[comp.Company], error) {
	log := r.log.With(zap.String("func", "ListCompanies"))
	log.Debug("ListCompanies called")

	var result list.List[comp.Company]

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err := transaction.Start(ctx, r.DB)
		if err != nil {
			return result, err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	queryBuilder := r.gen.Select("id", "name", "codename", "created_at", "updated_at").
		From("Company")

	if filter.Created.CreatedFrom != "" {
		queryBuilder = queryBuilder.Where(sq.GtOrEq{"created_at": filter.Created.CreatedFrom})
	}
	if filter.Created.CreatedTo != "" {
		queryBuilder = queryBuilder.Where(sq.LtOrEq{"created_at": filter.Created.CreatedTo})
	}

	if filter.OrderBy != "" {
		queryBuilder = queryBuilder.OrderBy(filter.OrderBy)
	} else {
		queryBuilder = queryBuilder.OrderBy("created_at DESC")
	}

	if filter.Page > 0 && filter.PerPage > 0 {
		offset := (filter.Page - 1) * filter.PerPage
		queryBuilder = queryBuilder.Limit(uint64(filter.PerPage)).Offset(uint64(offset))
	}

	query, args, _ := queryBuilder.ToSql()

	rows, err := tx.Queryx(query, args...)
	if err != nil {
		log.Errorw("failed to fetch companies", "err", err)
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var company comp.Company
		if err := rows.StructScan(&company); err != nil {
			log.Errorw("failed to scan company", "err", err)
			return result, err
		}
		result.Items = append(result.Items, company)
	}

	totalItems, err := r.Count(ctx, filter)
	if err != nil {
		return result, err
	}

	totalPages := (totalItems + filter.PerPage - 1) / filter.PerPage

	result.Pagination = list.Pagination{
		Page:       filter.Page,
		PerPage:    filter.PerPage,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}

	return result, nil
}
func (r *Repository) UpdateCompany(ctx context.Context, c comp.Company) (err error) {
	log := r.log.With(
		zap.String("func", "Update сompany"),
		zap.String("company_id", c.ID.String()),
	)
	log.Debug("called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("update"),
			errs.WithField("company_id", c.ID.String()),
		)
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err = transaction.Start(ctx, r.DB)
		if err != nil {
			return err
		}
		tx = transaction.Retrieve(ctx)
		defer transaction.End(ctx, err)
	}

	query, args, _ := r.gen.
		Update("Сompany").
		Set("name", c.Name).
		Set("codename", c.Codename).
		Set("updated_at", c.UpdatedAt.Time).
		Where(sq.Eq{"id": c.ID}).
		ToSql()

	_, err = tx.Exec(query, args...)
	if err != nil {
		log.Error("failed to update company", zap.Error(err))
	}
	return err
}
func (r *Repository) DeleteCompany(ctx context.Context, companyID uuid.UUID) (err error) {
	log := r.log.With(
		zap.String("func", "DeleteCompany"),
		zap.String("company_id", companyID.String()),
	)
	log.Debug("called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("delete"),
			errs.WithField("company_id", companyID.String()),
		)
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err = transaction.Start(ctx, r.DB)
		if err != nil {
			return err
		}
		tx = transaction.Retrieve(ctx)
		defer transaction.End(ctx, err)
	}

	query, args, _ := r.gen.
		Delete("Company").
		Where(sq.Eq{"id": companyID}).
		ToSql()

	_, err = tx.Exec(query, args...)
	if err != nil {
		log.Error("failed to delete company", zap.Error(err))
	}
	return err
}
