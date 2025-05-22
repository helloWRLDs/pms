package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (l *Logic) GetUserProfile(ctx context.Context, userID string) (*dto.User, error) {
	log := l.log.With(
		zap.String("func", "ListCompanies"),
	)
	log.Debug("ListCompanies called")

	// session, err := l.GetSessionInfo(ctx)
	// if err != nil {
	// 	log.Errorw("failed to get session", "err", err)
	// 	return nil, err
	// }

	res, err := l.authClient.GetUser(ctx, &pb.GetUserRequest{UserID: userID})
	if err != nil {
		log.Errorw("failed to get user", "err", err)
		return nil, err
	}

	return res.User, nil
}

func (l *Logic) ListUsers(ctx context.Context, filter *dto.UserFilter) (*dto.UserList, error) {
	log := l.log.Named("ListUsers").With()
	log.Debug("ListUsers called")

	userRes, err := l.authClient.ListUsers(ctx, &pb.ListUsersRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}
	userList := userRes.UserList
	return userList, nil
}
