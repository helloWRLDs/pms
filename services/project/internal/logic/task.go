package logic

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	"pms.pkg/type/list"
	"pms.pkg/utils"
	taskdata "pms.project/internal/data/task"
)

func (l *Logic) UpdateTask(ctx context.Context, id string, task *dto.Task) (err error) {
	log := l.log.Named("UpdateTask").With(
		zap.String("id", id),
		zap.Any("task", task),
	)
	log.Info("UpdateTask called")

	tx, err := l.Repo.StartTx(ctx)
	if err != nil {
		log.Errorw("failed to start tx", "err", err)
		return err
	}
	defer func() {
		l.Repo.EndTx(tx, err)
	}()

	if err := l.Repo.Task.Update(tx, id, *taskdata.Entity(task)); err != nil {
		log.Errorw("failed to update task", "err", err)
		return err
	}
	return nil
}

func (l *Logic) CreateTask(ctx context.Context, creation *dto.TaskCreation) (id string, err error) {
	log := l.log.With(
		zap.String("func", "CreateTask"),
		zap.Any("new_task", utils.JSON(creation)),
	)
	log.Debug("CreateTask called")

	tx, err := l.Repo.StartTx(ctx)
	if err != nil {
		log.Errorw("failed to start tx", "err", err)
		return "", err
	}
	defer func() {
		log.Debugw("err check", "isNil", err == nil, "err", err)
		l.Repo.EndTx(tx, err)
	}()

	newTask := taskdata.Task{
		ID:        uuid.NewString(),
		Title:     creation.GetTitle(),
		Body:      creation.GetBody(),
		Status:    consts.TaskStatus(creation.GetStatus()),
		Priority:  utils.Ptr(int(creation.GetPriority())),
		ProjectID: creation.GetProjectId(),
		SprintID:  utils.Ptr(creation.GetSprintId()),
		DueDate:   utils.Ptr(creation.GetDueDate().AsTime()),
		Created:   time.Now(),
		Code:      l.Repo.Project.GetCode(ctx, creation.ProjectId),
		Type:      consts.TaskType(creation.GetType()),
	}

	if err = l.Repo.Task.Create(tx, newTask); err != nil {
		log.Errorw("failed to create new task", "err", err)
		return "", err
	}

	return newTask.ID, nil
}

func (l *Logic) GetTask(ctx context.Context, id string) (task *dto.Task, err error) {
	log := l.log.With(
		zap.String("func", "GetTask"),
		zap.String("id", id),
	)
	log.Debug("GetTask called")

	tx, err := l.Repo.StartTx(ctx)
	if err != nil {
		log.Errorw("failed to start tx", "err", err)
		return nil, err
	}
	defer func() {
		log.Debugw("err check", "isNil", err == nil, "err", err)
		l.Repo.EndTx(tx, err)
	}()

	task = new(dto.Task)
	if t, err := l.Repo.Task.GetByID(tx, id); err != nil {
		return nil, err
	} else {
		task = t.DTO()
	}

	// TODO: Get AssigneeID, etc..

	return
}

func (l *Logic) ListTasks(ctx context.Context, filter *dto.TaskFilter) (result list.List[*dto.Task], err error) {
	log := l.log.With(
		zap.String("func", "ListTasks"),
		zap.Any("filter", filter.String()),
	)
	log.Debug("ListTasks called")

	tasks, err := l.Repo.Task.List(ctx, filter)
	if err != nil {
		log.Errorw("failed to list tasks", "err", err)
		return list.List[*dto.Task]{}, err
	}

	result = list.List[*dto.Task]{
		Pagination: list.Pagination{
			Page:       tasks.Page,
			PerPage:    tasks.PerPage,
			TotalPages: tasks.TotalPages,
			TotalItems: tasks.TotalItems,
		},
	}
	result.Items = make([]*dto.Task, len(tasks.Items))
	for i, t := range tasks.Items {
		result.Items[i] = t.DTO()
		if assignment, err := l.Repo.TaskAssignment.GetByTask(ctx, t.ID); err == nil {
			result.Items[i].AssigneeId = assignment.UserID
		}
	}

	return result, nil
}
