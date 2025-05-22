package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (l *Logic) CreateReportTemplate(ctx context.Context, docCreation *dto.DocumentCreation) (string, error) {
	log := l.log.Named("CreateReportTemplate").With(
		zap.Any("creation", docCreation),
	)
	log.Info("CreateReportTemplate called")

	reportRes, err := l.analyticsClient.CreateDocumentTemplate(ctx, &pb.CreateDocumentTemplateRequest{
		Creation: docCreation,
	})
	if err != nil {
		log.Errorw("failed to create template", "err", err)
		return "", err
	}
	return reportRes.DocId, nil
}

func (l *Logic) GetDocument(ctx context.Context, docID string) (*dto.Document, error) {
	log := l.log.Named("GetDocument").With(
		zap.String("doc_id", docID),
	)
	log.Info("GetDocument called")

	docRes, err := l.analyticsClient.GetDocument(ctx, &pb.GetDocumentRequest{
		Id: docID,
	})
	if err != nil {
		log.Errorw("failed to get document", "err", err)
		return nil, err
	}
	doc := docRes.Doc

	return doc, nil
}

func (l *Logic) ListDocuments(ctx context.Context, filter *dto.DocumentFilter) (*dto.DocumentList, error) {
	log := l.log.Named("ListDocuments").With(
		zap.Any("filter", filter),
	)
	log.Info("ListDocuments called")

	docRes, err := l.analyticsClient.ListDocuments(ctx, &pb.ListDocumentsRequest{
		Filter: filter,
	})
	if err != nil {
		log.Errorw("failed to list documents", "err", err)
		return nil, err
	}
	docList := docRes.Docs
	return docList, nil
}

func (l *Logic) UpdateDocument(ctx context.Context, id string, doc *dto.Document) error {
	log := l.log.Named("UpdateDocument").With(
		zap.Any("updated", doc),
		zap.String("id", id),
	)
	log.Info("UpdateDocument called")

	docRes, err := l.analyticsClient.UpdateDocument(ctx, &pb.UpdateDocumentRequest{
		DocId:      id,
		UpdatedDoc: doc,
	})
	if err != nil {
		log.Errorw("failed to update document", "err", err)
		return err
	}
	log.Infow("result", "res", docRes)

	return nil
}
