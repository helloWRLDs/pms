package userdata

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"pms.auth/internal/entity"
	"pms.pkg/errs"
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
		From("User u").
		LeftJoin("Participant p ON u.id = p.user_id")

	if filter.Date.From != "" {
		builder = builder.Where(sq.GtOrEq{"u.created_at": filter.Date.From})
	}
	if filter.Date.To != "" {
		builder = builder.Where(sq.LtOrEq{"u.created_at": filter.Date.To})
	}

	if filter.Order.By != "" {
		builder = builder.OrderBy(filter.Order.String())
	} else {
		builder = builder.OrderBy("u.created_at DESC")
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
		builder = builder.Where(fmt.Sprintf("u.%s IN (%v)", k, v))
	}
	q, a, _ := builder.ToSql()
	log.Info("query ", q)

	r.DB.QueryRowx(q, a...).Scan(&count)
	return count
}

// p - Particpant fields
func (r *Repository) List(ctx context.Context, filter list.Filters) (res list.List[entity.User], err error) {
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
		From("User u").
		LeftJoin("Participant p ON u.id = p.user_id")

	if filter.Date.From != "" {
		builder = builder.Where(sq.GtOrEq{"u.created_at": filter.Date.From})
	}
	if filter.Date.To != "" {
		builder = builder.Where(sq.LtOrEq{"u.created_at": filter.Date.To})
	}

	if filter.Order.By != "" {
		builder = builder.OrderBy(filter.Order.String())
	} else {
		builder = builder.OrderBy("u.created_at DESC")
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
		builder = builder.Where(fmt.Sprintf("u.%s IN (%v)", k, v))
	}
	res.TotalItems = int(r.Count(ctx, filter))
	res.Page = filter.Page
	res.PerPage = filter.PerPage
	res.TotalPages = (res.TotalItems + filter.PerPage - 1) / filter.PerPage

	builder = builder.Limit(uint64(filter.PerPage)).Offset(uint64((filter.Page - 1) * filter.PerPage))

	query, args, _ := builder.ToSql()
	log.Debugw("query built", "query", query, "args", args)

	rows, err := r.DB.QueryxContext(ctx, query, args...)
	if err != nil {
		log.Errorw("failed to fetch user", "err", err)
		return list.List[entity.User]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var company entity.User
		if err := rows.StructScan(&company); err != nil {
			log.Errorw("failed to scan user", "err", err)
			return list.List[entity.User]{}, err
		}
		res.Items = append(res.Items, company)
	}

	return res, nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (user entity.User, err error) {
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
		return entity.User{}, err
	}
	return user, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (user entity.User, err error) {
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

	q, a, _ := r.gen.
		Select("*").
		From(r.tableName).
		Where(sq.Eq{"id": id}).
		ToSql()

	if err = r.DB.QueryRowx(q, a...).StructScan(&user); err != nil {
		log.Errorw("failed to fetch user by id", "err", err)
		return entity.User{}, err
	}
	return user, nil
}

// func (r *Repository) Exists(ctx context.Context, )
