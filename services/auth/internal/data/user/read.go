package userdata

import (
	"context"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

func (r *Repository) Count(ctx context.Context, filter *dto.UserFilter) (count int64) {
	log := r.log.With(
		zap.String("func", "Count"),
		zap.Any("filters", filter),
	)
	log.Debug("Count called")

	builder := r.gen.
		Select("u.*").
		From("\"User\" u")

	if filter.CompanyId != "" || filter.Role != "" || filter.CompanyCodename != "" || filter.CompanyName != "" {
		builder = builder.
			LeftJoin("\"Participant\" p ON u.id = p.user_id").
			LeftJoin("\"Company\" c ON p.company_id = c.id")

		if filter.CompanyId != "" {
			builder = builder.Where(sq.Eq{"p.company_id": filter.CompanyId})
		}
		if filter.Role != "" {
			builder = builder.Where(sq.Eq{"p.role": filter.Role})
		}
		if filter.CompanyCodename != "" {
			builder = builder.Where(sq.Eq{"c.codename": filter.CompanyCodename})
		}
		if filter.CompanyName != "" {
			builder = builder.Where(sq.Eq{"c.name": filter.CompanyName})
		}
	}
	if filter.UserId != "" {
		builder = builder.Where(sq.Eq{"u.id": filter.UserId})
	}
	if filter.UserEmail != "" {
		builder = builder.Where(sq.Eq{"u.email": filter.UserEmail})
	}
	if filter.UserName != "" {
		builder = builder.Where(sq.Eq{"u.name": filter.UserName})
	}
	if filter.UserPhone != "" {
		builder = builder.Where(sq.Eq{"u.phone": filter.UserPhone})
	}

	q, a, _ := builder.ToSql()
	log.Info("query ", q)

	r.DB.QueryRowx(q, a...).Scan(&count)
	return count
}

func (r *Repository) Exists(ctx context.Context, field string, value interface{}) (exists bool) {
	log := r.log.With(
		zap.String("func", "Exists"),
		zap.Any("condition", fmt.Sprintf("%s: %v", field, value)),
	)
	log.Debug("Exists called")

	q := `SELECT EXISTS(SELECT id FROM "User" WHERE %s = $1)`

	if err := r.DB.QueryRowx(fmt.Sprintf(q, field), value).Scan(&exists); err != nil {
		return false
	}
	return exists
}

// p - Particpant fields
func (r *Repository) List(ctx context.Context, filter *dto.UserFilter) (res list.List[User], err error) {
	log := r.log.With(
		zap.String("func", "List"),
		zap.Any("filters", filter),
	)
	log.Debug("List called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("list"),
			errs.WithObject("users"),
		)
	}()

	builder := r.gen.
		Select("u.*").
		From("\"User\" u")

	if filter.CompanyId != "" || filter.Role != "" || filter.CompanyCodename != "" || filter.CompanyName != "" {
		builder = builder.
			LeftJoin("\"Participant\" p ON u.id = p.user_id").
			LeftJoin("\"Company\" c ON p.company_id = c.id")

		if filter.CompanyId != "" {
			builder = builder.Where(sq.Eq{"p.company_id": filter.CompanyId})
		}
		if filter.Role != "" {
			builder = builder.Where(sq.Eq{"p.role": filter.Role})
		}
		if filter.CompanyCodename != "" {
			builder = builder.Where(sq.Eq{"c.codename": filter.CompanyCodename})
		}
		if filter.CompanyName != "" {
			builder = builder.Where(sq.Eq{"c.name": filter.CompanyName})
		}
	}
	if filter.UserId != "" {
		builder = builder.Where(sq.Eq{"u.id": filter.UserId})
	}
	if filter.UserEmail != "" {
		builder = builder.Where(sq.Eq{"u.email": filter.UserEmail})
	}
	if filter.UserName != "" {
		builder = builder.Where(sq.Eq{"u.name": filter.UserName})
	}
	if filter.UserPhone != "" {
		builder = builder.Where(sq.Eq{"u.phone": filter.UserPhone})
	}

	if filter.DateFrom != "" {
		builder = builder.Where(sq.GtOrEq{"u.created_at": filter.DateFrom})
	}
	if filter.DateTo != "" {
		builder = builder.Where(sq.LtOrEq{"u.created_at": filter.DateTo})
	}

	{ // build pagination info
		filter.Page = utils.If(filter.Page <= 0, 1, filter.Page)
		filter.PerPage = utils.If(filter.PerPage <= 0, 10, filter.PerPage)

		countQuery, countArgs, _ := builder.ToSql()
		if err := r.DB.QueryRowx(strings.ReplaceAll(countQuery, "SELECT u.*", "SELECT COUNT(*)"), countArgs...).Scan(&res.TotalItems); err != nil {
			log.Errorw("failed to count users", "err", err)
			return list.List[User]{}, err
		}
		res.Page = int(filter.Page)
		res.PerPage = int(filter.PerPage)
		res.TotalPages = int((int32(res.TotalItems) + filter.PerPage - 1) / filter.PerPage)
	}

	if filter.OrderBy != "" {
		builder = builder.OrderBy(filter.OrderBy + " " + filter.OrderDirection)
	} else {
		builder = builder.OrderBy("u.created_at DESC")
	}

	builder = builder.Limit(uint64(filter.PerPage)).Offset(uint64((filter.Page - 1) * filter.PerPage))

	query, args, _ := builder.ToSql()
	log.Debugw("query built", "query", query, "args", args)

	rows, err := r.DB.QueryxContext(ctx, query, args...)
	if err != nil {
		log.Errorw("failed to fetch user", "err", err)
		return list.List[User]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var usr User
		if err := rows.StructScan(&usr); err != nil {
			log.Errorw("failed to scan user", "err", err)
			return list.List[User]{}, err
		}
		res.Items = append(res.Items, usr)
	}

	return res, nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (user User, err error) {
	log := r.log.With(
		zap.String("func", "GetByEmail"),
		zap.String("email", email),
	)
	log.Debug("GetByEmail called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("get"),
			errs.WithObject("user"),
			errs.WithField("email", email),
		)
	}()

	q, a, _ := r.gen.
		Select("*").
		From(r.tableName).
		Where(sq.Eq{"email": email}).
		ToSql()

	if err = r.DB.QueryRowx(q, a...).StructScan(&user); err != nil {
		log.Error("failed to fetch user by email", "err", err)
		return User{}, err
	}
	return user, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (user User, err error) {
	log := r.log.With(
		zap.String("func", "GetByID"),
		zap.String("id", id),
	)
	log.Debug("GetByID called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("get"),
			errs.WithObject("users"),
			errs.WithField("id", id),
		)
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err := transaction.Start(ctx, r.DB)
		if err != nil {
			return user, err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	q, a, _ := r.gen.
		Select("*").
		From(r.tableName).
		Where(sq.Eq{"id": id}).
		ToSql()

	if err = tx.QueryRowx(q, a...).StructScan(&user); err != nil {
		log.Errorw("failed to fetch user by id", "err", err)
		return User{}, err
	}
	return user, nil
}

// func (r *Repository) Exists(ctx context.Context, )
