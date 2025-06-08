package logic

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	userdata "pms.auth/internal/data/user"
	"pms.pkg/consts"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

func (l *Logic) InitiateOAuth2(provider string) (authURL string, err error) {
	log := l.log.Named("InitiateOAuth2")
	log.Debug("InitiateOAuth2 called")

	switch consts.Provider(provider) {
	case consts.ProviderGoogle:
		authURL = l.googleClient.AuthURL(provider)
	// case consts.ProviderGitHub:
	// 	authURL = l.githubClient.AuthURL(provider)
	default:
		return "", errs.ErrNotFound{
			Object: "provider",
			Field:  "name",
			Value:  provider,
		}
	}

	return authURL, nil
}

func (l *Logic) CompleteOAuth2(ctx context.Context, provider string, code string) (*dto.User, *dto.AuthPayload, error) {
	log := l.log.Named("CompleteOAuth2 function").With(
		zap.String("provider", provider),
		zap.String("code", code),
	)
	log.Debug("CompleteOAuth2 called")

	switch consts.Provider(provider) {
	case consts.ProviderGoogle:

		if err := l.googleClient.SetToken(code); err != nil {
			log.Errorw("failed to set token", "err", err)
			return nil, nil, err
		}

		googleUser, err := l.googleClient.GetUserData()
		if err != nil {
			log.Errorw("failed to get user data", "err", err)
			return nil, nil, err
		}

		if exists := l.Repo.User.Exists(ctx, "email", googleUser.Email); !exists {
			user := userdata.User{
				ID:        uuid.NewString(),
				FirstName: googleUser.GivenName,
				LastName:  utils.Ptr(googleUser.FamilyName),
				Email:     googleUser.Email,
				AvatarURL: utils.Ptr(googleUser.Picture),
			}
			log.Debug("user created", zap.Any("user", user))
			if err := l.Repo.User.Create(ctx, user); err != nil {
				log.Errorw("failed to create user", "err", err)
				return nil, nil, err
			}
		}

		user, err := l.Repo.User.GetByEmail(ctx, googleUser.Email)
		if err != nil {
			log.Errorw("failed to get user by email", "err", err)
			return nil, nil, err
		}

		payload, err := l.LoginUser(ctx, &provider, &dto.UserCredentials{
			Email:    googleUser.Email,
			Password: "",
		})
		if err != nil {
			log.Errorw("failed to login user", "err", err)
			return nil, nil, err
		}
		participations, err := l.Repo.Participant.GetByUserID(ctx, payload.User.Id)
		if err == nil {
			payload.User.Permissions = make(map[string]*dto.StringArray)
			for _, participation := range participations {
				log.Infow("participation check", "company_id", participation.CompanyID, "role", participation.Role)
				role, err := l.Repo.Role.GetByName(ctx, participation.Role)
				if err != nil {
					log.Errorw("failed to get role", "err", err)
					continue
				}
				payload.User.Permissions[participation.CompanyID] = &dto.StringArray{
					Values: role.Permissions,
				}
			}
		} else {
			log.Errorw("failed to get participations", "err", err)
		}
		log.Infow("user permissions", "permissions", payload.User.Permissions)

		return user.DTO(), payload, nil
	}

	return nil, nil, errs.ErrInternal{
		Reason: "unresolved oauth2 provider",
	}
}
