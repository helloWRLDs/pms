package grpchandler

import (
	"context"

	"pms.pkg/errs"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) CreateDocumentTemplate(ctx context.Context,
	req *pb.CreateDocumentTemplateRequest) (res *pb.CreateDocumentTemplateResponse, err error) {
	defer func() {
		err = errs.WrapGRPC(err)
	}()

	res = new(pb.CreateDocumentTemplateResponse)
	res.Success = false

	docID, err := s.logic.CreateReportTemplate(ctx, req.Creation)
	if err != nil {
		return res, err
	}
	res.Success = true
	res.DocId = docID

	return res, nil
}
