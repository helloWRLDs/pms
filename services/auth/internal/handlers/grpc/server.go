package grpc

import (
	"context"

	pb "pms.pkg/protobuf/services"
)

type ServerGRPC struct {
	pb.UnimplementedAuthServiceServer
}

func (s *ServerGRPC) CreateCompany(context.Context, *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	return nil, nil
}

func (s *ServerGRPC) DeleteCompany(context.Context, *pb.DeleteCompanyRequest) (*pb.DeleteCompanyResponse, error) {
	return nil, nil
}

func (s *ServerGRPC) DeleteParticipant(context.Context, *pb.DeleteParticipantRequest) (*pb.DeleteParticipantResponse, error) {
	return nil, nil
}

func (s *ServerGRPC) DeleteUser(context.Context, *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}

func (s *ServerGRPC) GetCompany(context.Context, *pb.GetCompanyRequest) (*pb.GetCompanyResponse, error) {
	return nil, nil
}

func (s *ServerGRPC) GetParticipant(context.Context, *pb.GetParticipantRequest) (*pb.GetParticipantResponse, error) {
	return nil, nil
}

func (s *ServerGRPC) GetUser(context.Context, *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, nil
}

func (s *ServerGRPC) LoginUser(context.Context, *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	return nil, nil
}

func (s *ServerGRPC) RegisterParticipant(context.Context, *pb.RegisterParticipantRequest) (*pb.RegisterParticipantResponse, error) {
	return nil, nil
}

func (s *ServerGRPC) RegisterUser(context.Context, *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	return nil, nil
}

func (s *ServerGRPC) UpdateCompany(context.Context, *pb.UpdateCompanyRequest) (*pb.UpdateCompanyResponse, error) {
	return nil, nil
}

func (s *ServerGRPC) UpdateParticipant(context.Context, *pb.UpdateParticipantRequest) (*pb.UpdateParticipantResponse, error) {
	return nil, nil
}

func (s *ServerGRPC) UpdateUser(context.Context, *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return nil, nil
}

func (s *ServerGRPC) ValidateToken(context.Context, *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	return nil, nil
}
