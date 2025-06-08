package logic

import (
	"context"
	"time"

	"go.uber.org/zap"
	"pms.api-gateway/internal/models"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (l *Logic) CompleteOAuth2(ctx context.Context, provider string, code string) (*dto.User, *dto.AuthPayload, error) {
	log := l.log.Named("CompleteOAuth2").With(
		zap.String("provider", provider),
		zap.String("code", code),
	)
	log.Debug("CompleteOAuth2 called")

	req := pb.CompleteOAuth2Request{
		Provider: provider,
		Code:     code,
	}
	oauthRes, err := l.authClient.CompleteOAuth2(ctx, &req)
	if err != nil {
		log.Errorw("failed to complete oauth2", "err", err)
		return nil, nil, err
	}
	payload := oauthRes.Payload

	compRes, err := l.authClient.ListCompanies(ctx, &pb.ListCompaniesRequest{
		Filter: &dto.CompanyFilter{
			Page:    1,
			PerPage: 1000,
			UserId:  payload.User.Id,
		},
	})
	if err != nil {
		log.Errorw("failed to list companies", "err", err)
	}
	session := models.Session{
		ID:           payload.SessionId,
		UserID:       payload.User.Id,
		AccessToken:  payload.AccessToken,
		RefreshToken: payload.RefreshToken,
		Permissions: func() (perm map[string][]consts.Permission) {
			perm = make(map[string][]consts.Permission)
			for companyID, permissions := range payload.User.Permissions {
				perm[companyID] = make([]consts.Permission, len(permissions.Values))
				for i, value := range permissions.Values {
					perm[companyID][i] = consts.Permission(value)
				}
			}
			return
		}(),
		Expires: time.Unix(payload.Exp, 0),
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
		return nil, nil, err
	}
	return oauthRes.User, oauthRes.Payload, nil
}

func (l *Logic) InitiateOAuth2(ctx context.Context, provider string) (string, error) {
	log := l.log.Named("InitiateOAuth2").With(
		zap.String("provider", provider),
	)
	log.Debug("InitiateOAuth2 called")

	req := pb.InitiateOAuth2Request{
		Provider: provider,
	}
	oauthRes, err := l.authClient.InitiateOAuth2(ctx, &req)
	if err != nil {
		log.Errorw("failed to initiate oauth2", "err", err)
		return "", err
	}

	return oauthRes.AuthUrl, nil
}

func (l *Logic) RegisterUser(ctx context.Context, newUser *dto.NewUser) (*dto.User, error) {
	log := l.log.Named("RegisterUser").With(
		zap.Any("new_user", newUser),
	)
	log.Debug("RegisterUser called")

	req := pb.RegisterUserRequest{
		NewUser: newUser,
	}
	registered, err := l.authClient.RegisterUser(ctx, &req)
	if err != nil {
		log.Errorw("failed to register user", "err", err)
		return nil, err
	}
	return registered.User, nil
}

func (l *Logic) LoginUser(ctx context.Context, creds *dto.UserCredentials) (*dto.AuthPayload, error) {
	log := l.log.Named("LoginUser").With(
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
		Permissions: func() (perm map[string][]consts.Permission) {
			perm = make(map[string][]consts.Permission)
			for companyID, permissions := range loginRes.Payload.User.Permissions {
				perm[companyID] = make([]consts.Permission, len(permissions.Values))
				for i, value := range permissions.Values {
					perm[companyID][i] = consts.Permission(value)
				}
			}
			return
		}(),
		Expires: time.Unix(loginRes.Payload.Exp, 0),
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

func (l *Logic) GetUserRole(ctx context.Context, userID string, companyID string) (*dto.Role, error) {
	log := l.log.Named("GetUserRole").With(
		zap.String("user_id", userID),
		zap.String("company_id", companyID),
	)
	log.Debug("GetUserRole called")

	req := pb.GetUserRoleRequest{
		UserId:    userID,
		CompanyId: companyID,
	}
	role, err := l.authClient.GetUserRole(ctx, &req)
	if err != nil {
		log.Errorw("failed to get user role", "err", err)
		return nil, err
	}
	return role.Role, nil
}
