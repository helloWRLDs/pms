package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (l *Logic) ListTaskComments(ctx context.Context, filter *dto.TaskCommentFilter) (*dto.TaskCommentList, error) {
	log := l.log.Named("ListTaskComments").With(
		zap.Any("filter", filter),
	)
	log.Info("ListTaskComments called")

	listRes, err := l.projectClient.ListTaskComments(ctx, &pb.ListTaskCommentsRequest{
		Filter: filter,
	})
	if err != nil {
		log.Errorw("failed to list task comments", "err", err)
		return nil, err
	}

	for i, c := range listRes.List.Items {
		userRes, err := l.authClient.GetUser(ctx, &pb.GetUserRequest{
			UserID: c.User.Id,
		})
		if err == nil {
			listRes.List.Items[i].User = userRes.User
		}
	}
	return listRes.List, nil
}

func (l *Logic) CreateTaskComment(ctx context.Context, creation *dto.TaskCommentCreation) (*dto.TaskComment, error) {
	log := l.log.Named("CreateTaskComment").With(
		zap.Any("comment_creation", creation),
	)
	log.Info("CreateTaskComment called")

	creationRes, err := l.projectClient.CreateTaskComment(ctx, &pb.CreateTaskCommentRequest{
		Creation: creation,
	})
	if err != nil {
		log.Errorw("failed to create task comment", "err", err)
		return nil, err
	}
	return creationRes.CreatedComment, nil
}

func (l *Logic) GetTaskComment(ctx context.Context, commentiD string) (*dto.TaskComment, error) {
	log := l.log.Named("GetTaskComment").With(
		zap.String("comment_id", commentiD),
	)
	log.Info("GetTaskComment called")

	commentRes, err := l.projectClient.GetTaskComment(ctx, &pb.GetTaskCommentRequest{
		Id: commentiD,
	})
	if err != nil {
		return nil, err
	}
	return commentRes.Comment, nil
}
