package auth

import (
	"context"

	"pms.pkg/protobuf/dto"
	pb "pms.pkg/protobuf/services"
)

type AuthService struct {
	pb.AuthClient
}

func (a *AuthService) LoginUser(ctx context.Context, email, password string) error {
	creds := dto.UserCreds{Email: email, Password: password}
	a.AuthClient.LoginUser(ctx, &pb.LoginUserRequest{
		Credentials: &creds,
	})

	return nil
}
