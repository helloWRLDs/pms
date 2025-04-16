package userdata

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"pms.auth/internal/entity"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/type/list"
)

func (r *Repository) Create(ctx context.Context, newUser entity.User) (err error) {
	log := r.log.With(
		zap.String("func", "CreateUser"),
		zap.String("email", newUser.Email),
		zap.String("name", newUser.Name),
	)
	log.Debug("CreateUser called")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithOperation("create"), errs.WithField("email", newUser.Email))
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
		Insert("User").
		Columns("name", "email", "password").
		Values(newUser.Name, newUser.Email, newUser.Password).
		ToSql()

	if _, err := tx.Exec(query, args...); err != nil {
		log.Errorw("failed to create user", "err", err)
		return err
	}
	log.Debug("user created")
	return nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (user entity.User, err error) {
	log := r.log.With(
		zap.String("func", "GetByEmail"),
		zap.String("email", email),
	)
	log.Debug("CreateUser called")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithField("email", email), errs.WithOperation("retrieve"))
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
		From("User").
		Where(sq.Eq{"email": email}).
		ToSql()

	if err := tx.QueryRowx(query, args...).StructScan(&user); err != nil {
		log.Warnw("failed to fetch user by email", "err", err)
		return user, err
	}
	return user, nil
}

func (r *Repository) Exists(ctx context.Context, email string) bool {
	var exists bool
	var err error
	log := r.log.With(
		zap.String("func", "Exists"),
		zap.Bool("exists", exists),
	)
	// defer func() {
	// 	err := r.errctx.MapSQL(err, errs.WithOperation("Exists"))
	// }()
	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err = transaction.Start(ctx, r.DB)
		if err != nil {
			return false
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}
	query, args, _ := r.gen.
		Select("*").
		From("user").
		Where(sq.Eq{"email": email}).
		ToSql()
	if err := r.DB.QueryRowx(query, args...).Scan(&exists); err != nil {
		log.Info("User does not exist")
		return false
	}
	return exists
}

func (r *Repository) GetByID(ctx context.Context, id string) (user entity.User, err error) {
	log := logrus.
		WithField("id", id).
		WithField("func", "Get")

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
		From("user").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err = tx.QueryRowx(query, args...).StructScan(&user); err != nil {
		log.WithError(err).Warn("failed to fetch user by id")
		return user, err
	}
	return user, nil
}

func (r *Repository) ListUsers(ctx context.Context, filter list.Filters) (list.List[entity.User], error) {

	log := r.log.With(
		zap.String("func", "ListUsers"),
		zap.String("filters", "filters"),
	)
	log.Debug("ListUsers called")

	var result list.List[entity.User]

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		var err error
		ctx, err := transaction.Start(ctx, r.DB)
		if err != nil {
			return result, err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}
	queryBuilder := r.gen.Select("id", "name", "email", "password", "avatar_img", "created_at", "updated_at").
		From("User")
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
		log.Errorw("failed to fetch users", "err", err)
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err := rows.StructScan(&user); err != nil {
			log.Errorw("failed to scan user", "err", err)
			return result, err
		}
		result.Items = append(result.Items, user)
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
func (r *Repository) Count(ctx context.Context, filter list.Filters) (int, error) {
	var count int
	var err error
	log := r.log.With(zap.String("func", "Count"))
	log.Debug("Count called", zap.Any("filters", filter))

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithOperation("count"))
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err = transaction.Start(ctx, r.DB)
		if err != nil {
			return -1, err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	queryBuilder := r.gen.Select("COUNT(*)").From("User")

	if filter.Created.CreatedFrom != "" {
		queryBuilder = queryBuilder.Where(sq.GtOrEq{"created_at": filter.Created.CreatedFrom})
	}
	if filter.Created.CreatedTo != "" {
		queryBuilder = queryBuilder.Where(sq.LtOrEq{"created_at": filter.Created.CreatedTo})
	}

	query, args, _ := queryBuilder.ToSql()

	if err := tx.QueryRow(query, args...).Scan(&count); err != nil {
		log.Errorw("failed to count users", "err", err)
		return 0, err
	}

	return count, nil
}

func (r *Repository) GetAll(ctx context.Context) (users []entity.User, err error) {
	log := r.log.With(zap.String("func", "GetAll"))
	log.Debug("GetAll called")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithOperation("retrieve all"))
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
		From("User").
		ToSql()

	rows, err := tx.Queryx(query, args...)
	if err != nil {
		log.Errorw("failed to fetch all users", "err", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err := rows.StructScan(&user); err != nil {
			log.Errorw("failed to scan user", "err", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
func (r *Repository) Update(ctx context.Context, u entity.User) (err error) {
	log := r.log.With(
		zap.String("func", "UpdateUser"),
		zap.String("user_id", u.ID.String()),
	)
	log.Debug("Update user called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("update"),
			errs.WithField("user_id", u.ID.String()),
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
		Update("User").
		Set("name", u.Name).
		Set("email", u.Email).
		Set("password", u.Password).
		Set("avatar_img", u.AvatarIMG).
		Set("updated_at", u.UpdatedAt.Time).
		Where(sq.Eq{"id": u.ID}).
		ToSql()

	_, err = tx.Exec(query, args...)
	if err != nil {
		log.Error("failed to update user", zap.Error(err))
	}
	return err
}
func (r *Repository) DeleteUser(ctx context.Context, userID uuid.UUID) (err error) {
	log := r.log.With(
		zap.String("func", "Delete User"),
		zap.String("user_id", userID.String()),
	)
	log.Debug("called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("delete"),
			errs.WithField("user_id", userID.String()),
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
		Delete("User").
		Where(sq.Eq{"id": userID}).
		ToSql()

	_, err = tx.Exec(query, args...)
	if err != nil {
		log.Error("failed to delete user", zap.Error(err))
	}
	return err
}
