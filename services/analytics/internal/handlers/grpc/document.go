package grpchandler

import (
	"context"

	"pms.pkg/errs"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) GetDocument(ctx context.Context, req *pb.GetDocumentRequest) (res *pb.GetDocumentResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.GetDocumentResponse)
	res.Success = false

	doc, err := s.logic.GetDocument(ctx, req.Id)
	if err != nil {
		return res, err
	}
	res.Success = true
	res.Doc = doc

	return res, nil
}

func (s *ServerGRPC) ListDocuments(ctx context.Context, req *pb.ListDocumentsRequest) (res *pb.ListDocumentsResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.ListDocumentsResponse)
	res.Success = false

	list, err := s.logic.ListDocuments(ctx, req.Filter)
	if err != nil {
		return res, err
	}
	res.Success = true
	res.Docs = list

	return res, nil
}

func (s *ServerGRPC) UpdateDocument(ctx context.Context, req *pb.UpdateDocumentRequest) (res *pb.UpdateDocumentResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.UpdateDocumentResponse)
	res.Success = false

	if err := s.logic.UpdateDocument(ctx, req.DocId, req.UpdatedDoc); err != nil {
		return res, err
	}
	res.Success = true

	return res, nil
}
