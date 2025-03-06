package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.api-gateway/internal/models"
	"pms.pkg/protobuf/dto"
	pb "pms.pkg/protobuf/services"
	"pms.pkg/utils"
)

func (l *Logic) RegisterUser(ctx context.Context, newUser *dto.NewUser) (*dto.User, error) {
	req := pb.RegisterUserRequest{
		NewUser: newUser,
	}
	registered, err := l.AuthClient.RegisterUser(ctx, &req)
	if err != nil {
		print(utils.JSON(err))
		return nil, err
	}
	return registered.User, nil
}

func (l *Logic) LoginUser(ctx context.Context, creds *dto.UserCredentials) (*dto.AuthPayload, error) {
	log := l.log.With(
		zap.String("func", "logic.LoginUser"),
		zap.String("email", creds.Email),
	)
	log.Debug("LoginUser called")

	req := pb.LoginUserRequest{
		Credentials: creds,
	}
	payload, err := l.AuthClient.LoginUser(ctx, &req)
	if err != nil {
		log.Errorw("failed to login user", "err", err)
		return nil, err
	}
	session := models.Session{
		ID:           payload.Payload.SessionId,
		UserID:       payload.Payload.User.Id,
		AccessToken:  payload.Payload.AccessToken,
		RefreshToken: payload.Payload.RefreshToken,
	}
	if err := l.Sessions.Set(ctx, payload.Payload.SessionId, session, 24); err != nil {
		log.Errorw("failed to setup session", "err", err)
		return nil, err
	}
	return payload.Payload, nil
}
