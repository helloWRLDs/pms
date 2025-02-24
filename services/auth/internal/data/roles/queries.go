package roledata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
	roledomain "pms.auth/internal/domain/role"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/type/nullable"
)

func (r *Repository) Create(ctx context.Context, newRole roledomain.Role) (err error) {
	log := logrus.WithFields(logrus.Fields{
		"func":        "roledata.Create",
		"name":        newRole.Name,
		"permissions": newRole.Permissions.StringArray(),
	})
	log.Debug("Create called")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithOperation("create"), errs.WithField("name", newRole.Name))
	}()

	tx, err := transaction.RetrieveCTX(ctx)
	if err != nil {
		tx, err = transaction.Start(r.DB)
		if err != nil {
			return err
		}
		defer func() {
			transaction.End(tx, err)
		}()
	}

	query, args, _ := r.gen.
		Insert("role").
		Columns("name", "permissions", "org_id").
		Values(newRole.Name, newRole.Permissions, nullable.String(newRole.OrganizationID)).
		ToSql()

	if err := tx.QueryRowx(query, args...).Err(); err != nil {
		log.WithError(err).Error("failed to create role")
	}

	return nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (role roledomain.Role, err error) {
	log := logrus.WithFields(logrus.Fields{
		"func":    "roledata.GetByID",
		"role_id": id,
	})
	log.Debug("GetByID called")

	query, args, _ := r.gen.
		Select("*").
		From("role").
		Where(sq.Eq{"id": id}).
		ToSql()

	err = r.DB.QueryRowx(query, args...).StructScan(&role)
	if err != nil {
		log.WithError(err).Error("failed to retrieve role from db")
		return
	}
	return role, nil
}

func (r *Repository) Get(ctx context.Context, orgID *string) (roles []roledomain.Role, err error) {
	log := logrus.WithFields(logrus.Fields{
		"func": "roledata.Get",
	})
	log.Debug("Get called")

	builder := r.gen.
		Select("*").
		From("role")

	if orgID != nil {
		builder = builder.Where(sq.Or{sq.Eq{"org_id": "NULL"}, sq.Eq{"org_id": orgID}})
	} else {
		builder = builder.Where(sq.Eq{"org_id": "NULL"})
	}
	query, args, _ := builder.ToSql()
	err = r.DB.QueryRowx(query, args...).StructScan(&roles)
	if err != nil {
		log.WithError(err).Error("failed to retrieve role from db")
		return
	}
	return roles, nil
}
