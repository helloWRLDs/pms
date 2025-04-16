package companydata

import (
	"context"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"go.uber.org/zap"
	part "pms.auth/internal/entity/participant"
	"pms.pkg/errs"
	"pms.pkg/tools/transaction"
	"pms.pkg/type/list"
	"pms.pkg/type/timestamp"
)

func (r *Repository) GetParticipantById(ctx context.Context, id int32) (participant part.Participant, err error) {
	log := r.log.With(
		zap.Int32("id", id),
		zap.String("func", "Get"),
	)
	defer func() {
		err = r.errctx.MapSQL(err, errs.WithField("id", strconv.Itoa(int(id))), errs.WithOperation("retrieve"))
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
		From("participant").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err = tx.QueryRowx(query, args...).StructScan(&participant); err != nil {
		log.Errorw("Failed to retrieve participant by id", err)
		return participant, err
	}
	return participant, nil
}
func (r *Repository) GetParticipantByCompanyId(ctx context.Context, cId string) (participant part.Participant, err error) {
	log := r.log.With(
		zap.String("company id", cId),
		zap.String("func", "Get"),
	)
	defer func() {
		err = r.errctx.MapSQL(err, errs.WithField("company id", cId), errs.WithOperation("retrieve"))
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
		From("participant").
		Where(sq.Eq{"company_id": cId}).
		ToSql()
	if err = tx.QueryRowx(query, args...).StructScan(&participant); err != nil {
		log.Errorw("Failed to retrieve participant by id", err)
		return participant, err
	}
	return participant, nil
}
func (r *Repository) CountParticipants(ctx context.Context, filters list.Filters) (list.List[part.Participant], error) {
	log := r.log.With(
		zap.String("func", "Count Participants"))
	log.Debug("Count participants called")
	var err error
	var result list.List[part.Participant]
	var count int
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
	qb := r.gen.Select("COUNT(*)").From("Participant")

	if filters.Created.CreatedFrom != "" {
		qb = qb.Where(sq.GtOrEq{"created_at": filters.Created.CreatedFrom})
	}
	if filters.Created.CreatedTo != "" {
		qb = qb.Where(sq.LtOrEq{"created_at": filters.Created.CreatedTo})
	}
	query, args, _ := qb.ToSql()
	err = tx.QueryRow(query, args...).Scan(&count)
	if err != nil {
		log.Error("failed to count participants", zap.Error(err))
		return result, err
	}
	perPage := filters.PerPage
	if perPage == 0 {
		perPage = 10
	}
	page := filters.Page
	if page == 0 {
		page = 1
	}
	totalPages := (count + perPage - 1) / perPage
	result.Pagination = list.Pagination{
		Page:       page,
		PerPage:    perPage,
		TotalItems: count,
		TotalPages: totalPages,
	}
	result.Items = nil // since count only
	log.Debug("pagination ended", zap.Any("pagination", result.Pagination))
	return result, err
}
func (r *Repository) ListParticipants(ctx context.Context, filters list.Filters) (list.List[part.Participant], error) {
	log := r.log.With(zap.String("func", "ListParticipants"))
	log.Debug("List participants called")

	var result list.List[part.Participant]
	var err error
	defer func() {
		err = r.errctx.MapSQL(err, errs.WithOperation("list"))
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err = transaction.Start(ctx, r.DB)
		if err != nil {
			return result, err
		}
	}
	tx = transaction.Retrieve(ctx)
	defer func() {
		transaction.End(ctx, err)
	}()

	page := filters.Page
	if page <= 0 {
		page = 1
	}

	perPage := filters.PerPage
	if perPage <= 0 {
		perPage = 10
	}
	offset := (page - 1) * perPage
	qb := r.gen.
		Select("*").
		From("Participant").
		Limit(uint64(perPage)).
		Offset(uint64(offset))

	if filters.Created.CreatedFrom != "" {
		qb = qb.Where(sq.GtOrEq{"created_at": filters.Created.CreatedFrom})
	}
	if filters.Created.CreatedTo != "" {
		qb.Where(sq.LtOrEq{"created_at": filters.Created.CreatedTo})
	}
	if filters.OrderBy != "" {
		qb = qb.OrderBy(filters.OrderBy)

	} else {
		qb = qb.OrderBy("created_at DESC")
	}
	query, args, _ := qb.ToSql()
	rows, err := tx.Queryx(query, args...)
	if err != nil {
		log.Error("failed to list participants", zap.Error(err))
		return result, err
	}
	defer rows.Close()

	var participants []part.Participant
	for rows.Next() {
		var p part.Participant
		if err := rows.StructScan(&p); err != nil {
			log.Error("failed to scan participant", zap.Error(err))
			return result, err
		}
		participants = append(participants, p)
	}
	countResults, err := r.CountParticipants(ctx, filters)
	if err != nil {
		return result, err
	}
	result = list.List[part.Participant]{
		Pagination: countResults.Pagination,
		Items:      participants,
	}
	result.Pagination.Page = page
	log.Debug("participants listed", zap.Int("count", len(participants)))
	return result, nil
}
func (r *Repository) ListParticipantsInfo(ctx context.Context, userID uuid.UUID) (participants []part.ParticipantInfo, err error) {
	log := r.log.With(zap.String("Func", "List ParticipantInfo"), zap.String("user_id", userID.String()))
	log.Debug("List partInfo called")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithOperation("retrieve"), errs.WithField("user_id", userID.String()))

	}()
	tx := transaction.Retrieve(ctx)
	if tx == nil {
		ctx, err = transaction.Start(ctx, r.DB)
		if err != nil {
			return nil, err
		}
		tx = transaction.Retrieve(ctx)
		defer transaction.End(ctx, err)
	}
	query := `
		SELECT
			p.company_id,
			c.name AS company_name,
			r.ID AS role_id,
			r.name AS role_name,
			r.permissions, 
			r.created_at AS role_created_at, 
			r.updated_at AS role_updated_at,
			p.created_at AS joined_at
		FROM participant p
		JOIN company c ON c.id = p.company_id
		JOIN role r ON r.id = p.role_id
		WHERE p.user_id = $1
	`
	rows, err := tx.QueryxContext(ctx, query, userID)
	if err != nil {
		log.Error("failed to list participant info", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pI part.ParticipantInfo
		var roleID uuid.UUID
		var createdAt, updatedAt, joinedAt timestamp.Timestamp
		var permissions []string

		err := rows.Scan(
			&pI.CompanyID,
			&pI.CompanyName,
			&roleID,
			&pI.Role.Name,
			pq.Array(&permissions),
			&createdAt,
			&updatedAt,
			&joinedAt,
		)
		if err != nil {
			log.Error("failed to scan row", zap.Error(err))
			return nil, err
		}
		pI.Role.ID = roleID
		pI.Role.Permissions = permissions
		pI.Role.CreatedAt = createdAt
		pI.Role.UpdatedAt = updatedAt
		pI.JoinedAt = joinedAt

		participants = append(participants, pI)
	}
	return participants, nil
}
func (r *Repository) UpdateParticipant(ctx context.Context, p part.Participant) (err error) {
	log := r.log.With(
		zap.String("func", "Update Participant"),
		zap.String("user_id", p.UserId),
		zap.String("company_id", p.CompanyId),
	)
	log.Debug("Update participant called")
	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("update"),
			errs.WithField("user_id", p.UserId),
			errs.WithField("company_id", p.CompanyId))

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
		Update("Participant").
		Set("role_id", p.RoleId).
		Set("updated_at", p.UpdatedAt.Time).
		Where(sq.Eq{"user_id": p.UserId, "company_id": p.CompanyId}).
		ToSql()
	_, err = tx.Exec(query, args...)
	if err != nil {
		log.Error("failed to update participant", zap.Error(err))
		return err
	}

	log.Debug("participant updated")
	return nil
}
func (r *Repository) DeleteParticipant(ctx context.Context, userID, companyID string) (err error) {
	log := r.log.With(
		zap.String("func", "Delete Participant"),
		zap.String("user_id", userID),
		zap.String("company_id", companyID),
	)
	log.Debug("Delete participant called")
	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("update"),
			errs.WithField("user_id", userID),
			errs.WithField("company_id", companyID))

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
		Update("Participant").
		Where(sq.Eq{"user_id": userID, "company_id": companyID}).
		ToSql()
	_, err = tx.Exec(query, args...)
	if err != nil {
		log.Error("failed to delete participant", zap.Error(err))
		return err
	}

	log.Debug("participant deleted")
	return nil
}
