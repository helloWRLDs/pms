package companydata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.auth/internal/entity"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/utils"
)

func (r *Repository) Create(ctx context.Context, newCompany entity.Company) (err error) {
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
		Columns(utils.GetColumns(newCompany)...).
		Values(utils.GetArguments(newCompany)...).
		ToSql()
	if _, err := tx.Exec(query, args...); err != nil {
		log.Errorw("failed to create company", "err", err)
		return err
	}
	log.Debug("Company created")
	return nil
}

func (r *Repository) DeleteCompany(ctx context.Context, companyID string) (err error) {
	log := r.log.With(
		zap.String("func", "DeleteCompany"),
		zap.String("company_id", companyID),
	)
	log.Debug("DeleteCompany called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("delete"),
			errs.WithField("company_id", companyID),
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
		Delete(r.tableName).
		Where(sq.Eq{"id": companyID}).
		ToSql()

	_, err = tx.Exec(query, args...)
	if err != nil {
		log.Error("failed to delete company", zap.Error(err))
	}
	return err
}

func (r *Repository) UpdateCompany(ctx context.Context, companyID string, c entity.Company) (err error) {
	log := r.log.With(
		zap.String("func", "UpdateCompany"),
		zap.String("company_id", c.ID.String()),
	)
	log.Debug("UpdateCompany called")

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

	cols := utils.GetColumns(c)
	args := utils.GetArguments(c)

	builder := r.gen.
		Update(r.tableName)

	for i, col := range cols {
		builder = builder.Set(col, args[i])
	}
	q, a, _ := builder.Where(sq.Eq{"id": companyID}).ToSql()

	if _, err = tx.Exec(q, a...); err != nil {
		log.Error("failed to update company", "err", err)
	}
	return nil
}

func (r *Repository) AddParticipant(ctx context.Context, userID, comapanyID string) (err error) {
	log := r.log.With(
		zap.String("func", "ListCompanies"),
		zap.String("user_id", userID),
		zap.String("company_id", comapanyID),
	)
	log.Debug("ListCompanies called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("list"),
			errs.WithObject("comapanies"),
		)
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err := transaction.Start(ctx, r.DB)
		if err != nil {
			return err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	q, a, _ := r.gen.
		Insert("Participant").
		Columns("user_id, company_id, role").
		Values(userID, comapanyID, "admin").ToSql()

	_, err = tx.Exec(q, a...)
	if err != nil {
		log.Errorw("failed to add participant", "err", err)
		return err
	}
	return nil
}
