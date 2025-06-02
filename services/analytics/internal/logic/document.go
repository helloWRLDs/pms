package logic

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	documentdata "pms.analytics/internal/data/document"
	htmlmodule "pms.analytics/internal/modules/htmlgen"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
)

func (l *Logic) DownloadDocument(ctx context.Context, docID string) (pdf *dto.DocumentPDF, err error) {
	log := l.log.Named("DownloadDocumment").With(
		zap.String("id", docID),
	)
	log.Debug("DownloadDocument called")

	doc, err := l.getDocument(ctx, docID)
	if err != nil {
		log.Errorw("failed to get document from db", "err", err)
		return nil, err
	}

	pdfData, err := htmlmodule.PDF(doc.Body)
	if err != nil {
		log.Errorw("failed to transform html to pdf", "err", err)
		return nil, err
	}
	pdf = new(dto.DocumentPDF)
	pdf.Body = pdfData
	pdf.Title = doc.Title
	pdf.DocId = doc.Id

	return pdf, nil
}

func (l *Logic) getDocument(ctx context.Context, id string) (*dto.Document, error) {
	log := l.log.Named("getDocument").With(
		zap.String("id", id),
	)
	log.Debug("getDocument called")

	if _, err := uuid.Parse(id); err != nil {
		return nil, errs.ErrBadGateway{
			Object: "doc_id",
		}
	}

	doc, err := l.Repo.Document.GetByID(ctx, id)
	if err != nil {
		log.Errorw("failed to get doc", "err", err)
		return nil, err
	}
	return doc.DTO(), nil
}

func (l *Logic) GetDocument(ctx context.Context, id string) (*dto.Document, error) {
	log := l.log.Named("GetDocument").With(
		zap.String("id", id),
	)
	log.Debug("GetDocument called")

	doc, err := l.getDocument(ctx, id)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (l *Logic) UpdateDocument(ctx context.Context, id string, updatedDoc *dto.Document) error {
	log := l.log.Named("UpdateDocument").With(
		zap.String("id", id),
		zap.Any("updatedDoc", updatedDoc),
	)
	log.Debug("UpdateDocument called")

	if err := l.Repo.Document.Update(ctx, id, *documentdata.Entity(updatedDoc)); err != nil {
		log.Errorw("failed to update doc", "err", err)
		return err
	}

	log.Info("doc updated")
	return nil
}

func (l *Logic) ListDocuments(ctx context.Context, filter *dto.DocumentFilter) (*dto.DocumentList, error) {
	log := l.log.Named("ListDocuments").With(
		zap.Any("filter", filter),
	)
	log.Debug("ListDocuments called")

	docList, err := l.Repo.Document.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	result := dto.DocumentList{
		Page:       int32(docList.Page),
		PerPage:    int32(docList.PerPage),
		TotalPages: int32(docList.TotalPages),
		TotalItems: int32(docList.TotalItems),
		Items:      make([]*dto.Document, len(docList.Items)),
	}
	for i, item := range docList.Items {
		result.Items[i] = item.DTO()
	}

	return &result, nil
}
