package logic

import (
	"context"
	"time"

	"go.uber.org/zap"
	"pms.api-gateway/internal/models"
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
	log := l.log.With(
		zap.String("func", "logic.LoginUser"),
		zap.String("email", creds.Email),
	)
	log.Debug("LoginUser called")

	req := pb.LoginUserRequest{
		Credentials: creds,
	}
	loginRes, err := l.authClient.LoginUser(ctx, &req)
	if err != nil {
		log.Errorw("failed to login user", "err", err)
		return nil, err
	}
	compRes, err := l.authClient.ListCompanies(ctx, &pb.ListCompaniesRequest{
		Filter: &dto.CompanyFilter{
			Page:    1,
			PerPage: 1000,
			UserId:  loginRes.Payload.User.Id,
		},
	})
	if err != nil {
		log.Errorw("failed to list companies", "err", err)
	}
	session := models.Session{
		ID:           loginRes.Payload.SessionId,
		UserID:       loginRes.Payload.User.Id,
		AccessToken:  loginRes.Payload.AccessToken,
		RefreshToken: loginRes.Payload.RefreshToken,
		Expires:      time.Unix(loginRes.Payload.Exp, 0),
		Companies: func() (compList []string) {
			compList = make([]string, 0)
			if compRes == nil {
				return
			}
			for _, comp := range compRes.Companies.Items {
				compList = append(compList, comp.Id)
			}
			return
		}(),
	}
	if err := l.Sessions.Set(ctx, session.ID, session, 24); err != nil {
		log.Errorw("failed to setup session", "err", err)
		return nil, err
	}
	return loginRes.Payload, nil
}
