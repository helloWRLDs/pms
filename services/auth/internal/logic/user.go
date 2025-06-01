package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
	"pms.pkg/utils"
)

func (l *Logic) ListUsers(ctx context.Context, filter *dto.UserFilter) (result list.List[*dto.User], err error) {
	log := l.log.With(
		zap.String("func", "ListUsers"),
		zap.Any("filter", filter),
	)
	log.Debug("ListUsers called")

	entities, err := l.Repo.User.List(ctx, filter)
	if err != nil {
		log.Errorw("failed to list users", "err", err)
		return list.List[*dto.User]{}, err
	}
	result.Page = entities.Page
	result.PerPage = entities.PerPage
	result.TotalPages = entities.TotalPages
	result.TotalItems = entities.TotalItems

	for _, usr := range entities.Items {
		result.Items = append(result.Items, func() (u *dto.User) {
			u = usr.DTO()
			participantList, err := l.Repo.Participant.List(ctx, &dto.ParticipantFilter{
				Page:    1,
				PerPage: 1000,
				UserId:  usr.ID,
			})
			if err == nil {
				for _, p := range participantList.Items {
					u.Participants = append(u.Participants, p.DTO())
				}
			}
			return u
		}())
	}

	return result, nil
}

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

	profile = user.DTO()

	return profile, nil
}

func (l *Logic) UpdateUser(ctx context.Context, id string, user *dto.User) (updated *dto.User, err error) {
	log := l.log.With(
		zap.String("func", "UpdateUser"),
		zap.String("user_id", id),
	)
	log.Debug("UpdateUser called")

	existing, err := l.Repo.User.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	isChanged := false
	{
		if existing.FirstName != user.FirstName {
			existing.FirstName = user.FirstName
			isChanged = true
		}
		if existing.LastName != utils.Ptr(user.LastName) {
			existing.LastName = utils.Ptr(user.LastName)
			isChanged = true
		}
		if existing.Email != user.Email {
			existing.Email = user.Email
			isChanged = true
		}
		if utils.Value(existing.Phone) != user.Phone {
			existing.Phone = &user.Phone
			isChanged = true
		}
		if utils.Value(existing.Bio) != user.Bio {
			existing.Bio = &user.Bio
			isChanged = true
		}
		if utils.Value(existing.AvatarURL) != user.AvatarUrl {
			existing.AvatarURL = utils.Ptr(user.AvatarUrl)
			isChanged = true
		}
	}

	if isChanged {
		if err := l.Repo.User.Update(ctx, id, existing); err != nil {
			return nil, err
		}
	}

	return l.GetProfile(ctx, id)
}
