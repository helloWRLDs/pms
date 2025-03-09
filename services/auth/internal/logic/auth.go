package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"pms.auth/internal/entity"
	"pms.pkg/errs"
	"pms.pkg/tools/jwtoken"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/claims"
	"pms.pkg/utils/validators"
)

func (l *Logic) LoginUser(ctx context.Context, creds *dto.UserCredentials) (payload *dto.AuthPayload, err error) {
	log := l.log.With(
		zap.String("func", "LoginUser"),
		zap.String("email", creds.Email),
		zap.String("password", creds.Password),
	)
	log.Debug("LoginUser called")
	payload = new(dto.AuthPayload)
	if creds.Email == "" {
		return nil, errs.ErrInvalidInput{
			Object: "email",
			Reason: "cannot be empty",
		}
	}
	if creds.Password == "" {
		return nil, errs.ErrInvalidInput{
			Object: "password",
			Reason: "cannot be empty",
		}
	}
	existingUser, err := l.repo.User.GetByEmail(ctx, creds.GetEmail())
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(creds.Password)); err != nil {
		return nil, err
	}

	claims := claims.AccessTokenClaims{
		Email:     existingUser.Email,
		UserID:    existingUser.ID.String(),
		SessionID: uuid.NewString(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	sessionID := uuid.NewString()
	accessToken, err := jwtoken.GenerateAccessToken(claims, &l.conf.JWT)
	if err != nil {
		return nil, errs.ErrInternal{
			Reason: "failed to generate jwt token",
		}
	}

	payload.SessionId = sessionID
	payload.User = existingUser.DTO()
	payload.AccessToken = accessToken
	payload.Exp = claims.ExpiresAt.Time.Unix()

	return payload, nil
}

func (l *Logic) RegisterUser(ctx context.Context, newUser *dto.NewUser) (createdUser *dto.User, err error) {
	log := l.log.With(
		zap.String("func", "RegsterUser"),
		zap.String("email", newUser.Email),
		zap.String("name", newUser.Name),
	)
	log.Debug("RegsterUser called")

	tx, err := l.repo.StartTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		log.Infow("err check", "isNil", err == nil)
		l.repo.EndTx(tx, err)
	}()

	if exists := l.repo.User.Exists(tx, newUser.Email); exists {
		log.Errorf("user with email = %s already exists", newUser.Email)
		return nil, errs.ErrConflict{
			Reason: fmt.Sprintf("user with email = %s already exists", newUser.Email),
		}
	}

	if err := validators.ValidateEmail(newUser.Email); err != nil {
		log.Warnw("invalid email", "err", err)
		return nil, err
	}
	if err := validators.ValidatePassword(newUser.Password); err != nil {
		log.Warnw("invalid password", "err", err)
		return nil, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.ErrInvalidInput{
			Object: "password",
			Reason: err.Error(),
		}
	}
	user := entity.User{
		Name:      newUser.GetName(),
		Email:     newUser.GetEmail(),
		Password:  string(hashedPassword),
		AvatarIMG: newUser.GetAvatarImg(),
	}

	if err := l.repo.User.Create(tx, user); err != nil {
		log.Errorw("failed to create user", "err", err)
		return nil, err
	}
	created, err := l.repo.User.GetByEmail(tx, user.Email)
	if err != nil {
		return nil, err
	}
	createdUser = created.DTO()
	return createdUser, nil
}
