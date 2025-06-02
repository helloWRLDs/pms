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
		l.Repo.EndTx(ctx, err)
	}()

	existing, _ := l.Repo.TaskAssignment.GetByTask(tx, taskID)

	if existing != nil {
		l.Repo.TaskAssignment.Delete(tx, assignmentdata.AssignmentData{
			TaskID: taskID,
			UserID: existing.UserID,
		})
	}
	newAssignment := assignmentdata.AssignmentData{
		TaskID: taskID,
		UserID: userID,
	}
	err = l.Repo.TaskAssignment.Create(tx, newAssignment)
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

	tx, err := l.Repo.StartTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		l.Repo.EndTx(ctx, err)
	}()

	existing, _ := l.Repo.TaskAssignment.Get(tx, userID, taskID)
	if existing == nil {
		return nil
	}

	if err := l.Repo.TaskAssignment.Delete(tx, assignmentdata.AssignmentData{
		UserID: userID,
		TaskID: taskID,
	}); err != nil {
		return err
	}
	return nil
}
