package logic

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"pms.pkg/transport/grpc/dto"
	taskcommentdata "pms.project/internal/data/taskcomment"
)

func (l *Logic) CreateTaskComment(ctx context.Context, creation *dto.TaskCommentCreation) (comment *dto.TaskComment, err error) {
	log := l.log.Named("CreateTaskComment").With(
		zap.Any("creation", creation),
	)
	log.Debug("CreateTaskComment called")

	tx, err := l.Repo.StartTx(ctx)
	if err != nil {
		log.Errorw("failed to start tx", "err", err)
		return nil, err
	}
	defer func() {
		log.Debugw("err check", "isNil", err == nil, "err", err)
		l.Repo.EndTx(tx, err)
	}()

	newComment := taskcommentdata.TaskComment{
		ID:        uuid.NewString(),
		Body:      creation.Body,
		TaskID:    creation.TaskId,
		UserID:    creation.UserId,
		CreatedAt: time.Now(),
	}

	if err = l.Repo.TaskComment.Create(tx, newComment); err != nil {
		log.Errorw("failed to create comment", "err", err)
		return nil, err
	}

	return newComment.DTO(), nil
}

func (l *Logic) ListTaskComments(ctx context.Context, filter *dto.TaskCommentFilter) (comments *dto.TaskCommentList, err error) {
	log := l.log.Named("ListTaskComments").With(
		zap.Any("filter", filter),
	)
	log.Debug("ListTaskComments called")

	commentEnts, err := l.Repo.TaskComment.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	comments = &dto.TaskCommentList{
		Page:       int32(commentEnts.Page),
		PerPage:    int32(commentEnts.PerPage),
		TotalPages: int32(commentEnts.TotalPages),
		TotalItems: int32(commentEnts.TotalItems),
	}

	if len(commentEnts.Items) > 0 {
		comments.Items = make([]*dto.TaskComment, len(commentEnts.Items))
		for i, c := range commentEnts.Items {
			comments.Items[i] = c.DTO()
		}
	}

	return comments, nil
}

func (l *Logic) GetTaskComment(ctx context.Context, id string) (*dto.TaskComment, error) {
	log := l.log.Named("GetTaskComment").With(
		zap.String("id", id),
	)
	log.Debug("GetTaskComment called")

	comment, err := l.getTaskComment(ctx, id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (l *Logic) getTaskComment(ctx context.Context, id string) (*dto.TaskComment, error) {
	log := l.log.Named("getTaskComment").With(
		zap.String("id", id),
	)
	log.Debug("getTaskComment called")

	comment, err := l.Repo.TaskComment.GetByID(ctx, id)
	if err != nil {
		log.Errorw("failed to get comment", "err", err)
		return nil, err
	}
	return comment.DTO(), nil
}
