package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.pkg/errs"
	assignmentdata "pms.project/internal/data/assignment"
)

func (l *Logic) AssignTask(ctx context.Context, userID, taskID string) (err error) {
	log := l.log.Named("AssignTask").With(
		zap.String("user_id", userID),
		zap.String("task_id", taskID),
	)
	log.Debug("AssignTask called")

	task, err := l.Repo.Task.GetByID(ctx, taskID)
	if err != nil {
		return err
	}
	if task.ID == "" {
		return errs.ErrNotFound{
			Object: "task",
			Field:  "id",
			Value:  taskID,
		}
	}

	tx, err := l.Repo.StartTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		l.Repo.EndTx(tx, err)
	}()

	existing, _ := l.Repo.TaskAssignment.GetByTask(tx, taskID)

	if existing != nil {
		if err := l.Repo.TaskAssignment.Delete(tx, assignmentdata.AssignmentData{
			TaskID: taskID,
			UserID: existing.UserID,
		}); err != nil {
			return err
		}
	}

	err = l.Repo.TaskAssignment.Create(tx, assignmentdata.AssignmentData{
		TaskID: taskID,
		UserID: userID,
	})
	if err != nil {
		return err
	}

	return nil

}

func (l *Logic) UnassignTask(ctx context.Context, userID, taskID string) (err error) {
	log := l.log.Named("UnassignTask").With(
		zap.String("user_id", userID),
		zap.String("task_id", taskID),
	)
	log.Debug("UnassignTask called")

	existing, _ := l.Repo.TaskAssignment.Get(ctx, taskID, userID)
	if existing == nil {
		return nil
	}

	if err := l.Repo.TaskAssignment.Delete(ctx, assignmentdata.AssignmentData{
		UserID: userID,
		TaskID: taskID,
	}); err != nil {
		return err
	}
	return nil
}
