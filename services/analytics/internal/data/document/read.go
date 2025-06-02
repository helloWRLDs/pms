package documentdata

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

func (r *Repository) GetByID(ctx context.Context, docID string) (doc Document, err error) {
	log := r.log.Named("GetByID").With(
		zap.Any("id", docID),
	)
	log.Debug("GetByID called")

	defer func() {
		err = r.errctx.MapSQL(err,
			errs.WithOperation("get"),
			errs.WithField("id", docID),
		)
	}()

	tx := transaction.Retrieve(ctx)
	if tx == nil {
		log.Debug("tx not found")
		ctx, err = transaction.Start(ctx, r.DB)
		if err != nil {
			return Document{}, err
		}
		tx = transaction.Retrieve(ctx)
		defer func() {
			transaction.End(ctx, err)
		}()
	}

	q, a, _ := r.gen.
		Select("*").
		From(r.tableName).
		Where(sq.Eq{"id": docID}).
		ToSql()

	if err = tx.QueryRowx(q, a...).StructScan(&doc); err != nil {
		return Document{}, err
	}
	return doc, nil
}

func (r *Repository) List(ctx context.Context, filter *dto.DocumentFilter) (res list.List[Document], err error) {
	log := r.log.Named("List").With(
		zap.Any("filter", filter),
	)
	log.Debug("List called")

	defer func() {
		err = r.errctx.MapSQL(err, errs.WithOperation("list"))
	}()

	builder := r.gen.
		Select("d.*").
		From("\"Document\" d")

	if filter.ProjectId != "" {
		builder = builder.Where(sq.Eq{"d.project_id": filter.ProjectId})
	}
	if filter.Title != "" {
		builder = builder.Where(sq.Eq{"d.title": filter.Title})
	}
	{
		filter.Page = utils.If(filter.Page <= 0, 1, filter.Page)
		filter.PerPage = utils.If(filter.PerPage <= 0, 10, filter.PerPage)

		countQuery, countArgs, _ := builder.ToSql()
		if err := r.DB.QueryRowx(strings.ReplaceAll(countQuery, "SELECT d.*", "SELECT COUNT(*)"), countArgs...).Scan(&res.TotalItems); err != nil {
			log.Errorw("failed to count users", "err", err)
			return list.List[Document]{}, err
		}
		res.Page = int(filter.Page)
		res.PerPage = int(filter.PerPage)
		res.TotalPages = int((int32(res.TotalItems) + filter.PerPage - 1) / filter.PerPage)
	}

	if filter.OrderBy != "" {
		builder = builder.OrderBy(filter.OrderBy + " " + filter.OrderDirection)
	} else {
		builder = builder.OrderBy("d.created_at DESC")
	}
	builder = builder.Limit(uint64(filter.PerPage)).Offset(uint64((filter.Page - 1) * filter.PerPage))

	query, args, _ := builder.ToSql()
	log.Debugw("query built", "query", query, "args", args)

	rows, err := r.DB.QueryxContext(ctx, query, args...)
	if err != nil {
		log.Errorw("failed to fetch doc", "err", err)
		return list.List[Document]{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var doc Document
		if err := rows.StructScan(&doc); err != nil {
			log.Errorw("failed to scan user", "err", err)
			return list.List[Document]{}, err
		}
		res.Items = append(res.Items, doc)
	}

	return res, nil
}
