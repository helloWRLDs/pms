package logic

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/utils"
)

func (l *Logic) GetProfile(ctx context.Context, userID string) (profile *dto.User, err error) {
	log := l.log.With(
		zap.String("func", "GetProfile"),
		zap.String("user_id", userID),
	)
	log.Debug("GetProfile called")

	user, err := l.Repo.User.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	profile = new(dto.User)

	profile.Id = user.ID.String()
	profile.Name = user.Name
	profile.Email = user.Email
	profile.AvatarImg = user.AvatarIMG
	profile.Phone = utils.Value(user.Phone)
	profile.Bio = utils.Value(user.Bio)
	profile.CreatedAt = timestamppb.New(user.CreatedAt.Time)
	profile.UpdatedAt = timestamppb.New(user.UpdatedAt.Time)

	return profile, nil
}
