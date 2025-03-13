package logic

import (
	"context"
	"time"

	"go.uber.org/zap"
	"pms.api-gateway/internal/models"
	"pms.pkg/logger"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
	"pms.pkg/utils"
)

func (l *Logic) RegisterUser(ctx context.Context, newUser *dto.NewUser) (*dto.User, error) {
	req := pb.RegisterUserRequest{
		NewUser: newUser,
	}
	registered, err := l.authClient.RegisterUser(ctx, &req)
	if err != nil {
		print(utils.JSON(err))
		return nil, err
	}
	return registered.User, nil
}

func (l *Logic) LoginUser(ctx context.Context, creds *dto.UserCredentials) (*dto.AuthPayload, error) {
	log := logger.Enabled(l.log, true).With(
		zap.String("func", "logic.LoginUser"),
		zap.String("email", creds.Email),
	)
	log.Debug("LoginUser called")

	req := pb.LoginUserRequest{
		Credentials: creds,
	}
	res, err := l.authClient.LoginUser(ctx, &req)
	if err != nil {
		log.Errorw("failed to login user", "err", err)
		return nil, err
	}
	session := models.Session{
		ID:           res.Payload.SessionId,
		UserID:       res.Payload.User.Id,
		AccessToken:  res.Payload.AccessToken,
		RefreshToken: res.Payload.RefreshToken,
		Expires:      time.Unix(res.Payload.Exp, 0),
	}
	if err := l.Sessions.Set(ctx, session.ID, session, 24); err != nil {
		log.Errorw("failed to setup session", "err", err)
		return nil, err
	}
	return res.Payload, nil
}
