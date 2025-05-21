package grpchandler

import (
	"context"

	"pms.pkg/errs"
	pb "pms.pkg/transport/grpc/services"
)

func (s *ServerGRPC) AddParticipant(ctx context.Context, req *pb.AddParticipantRequest) (res *pb.AddParticipantResponse, err error) {
	log := s.log.With("func", "RemoveParticipant", "pkg", "grpchandler")
	log.Debug("RemoveParticipant called")

	res = new(pb.AddParticipantResponse)
	res.Success = false

	defer func() {
		err = errs.WrapGRPC(err)
	}()

	participant, err := s.logic.AddParticipant(ctx, req.UserId, req.CompanyId)
	if err != nil {
		log.Errorw("failed to remove participant", "err", err)
		return nil, err
	}
	res.Success = true
	res.Participant = participant
	return res, nil
}

func (s *ServerGRPC) RemoveParticipant(ctx context.Context, req *pb.RemoveParticipantRequest) (res *pb.RemoveParticipantResponse, err error) {
	log := s.log.With("func", "RemoveParticipant", "pkg", "grpchandler")
	log.Debug("RemoveParticipant called")

	res = new(pb.RemoveParticipantResponse)
	res.Success = false

	defer func() {
		err = errs.WrapGRPC(err)
	}()

	if err := s.logic.DeleteParticipant(ctx, req.UserId, req.CompanyId); err != nil {
		return nil, err
	}
	res.Success = true
	return
}
