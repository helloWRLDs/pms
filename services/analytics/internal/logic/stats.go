package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.analytics/internal/utils"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (l *Logic) GetUserTaskStats(ctx context.Context, companyID string) (result []*dto.UserTaskStats, err error) {
	log := l.log.Named("GetUserTaskStats").With(
		zap.String("company_id", companyID),
	)
	log.Debug("GetUserTaskStats called")

	var (
		users    []*dto.User    = make([]*dto.User, 0)
		projects []*dto.Project = make([]*dto.Project, 0)
		sprints  []*dto.Sprint  = make([]*dto.Sprint, 0)
	)

	usersResp, err := l.authClient.ListUsers(ctx, &pb.ListUsersRequest{
		Filter: &dto.UserFilter{
			CompanyId: companyID,
			Page:      1,
			PerPage:   1000,
		},
	})
	if err != nil {
		log.Error("failed to list users", zap.Error(err))
		return nil, err
	}
	users = usersResp.UserList.Items
	projectsResp, err := l.projectClient.ListProjects(ctx, &pb.ListProjectsRequest{
		Filter: &dto.ProjectFilter{
			CompanyId: companyID,
			Page:      1,
			PerPage:   1000,
		},
	})
	if err != nil {
		log.Error("failed to list projects", zap.Error(err))
		return nil, err
	}
	projects = projectsResp.Projects.Items
	for _, project := range projects {
		sprintsResp, err := l.projectClient.ListSprints(ctx, &pb.ListSprintsRequest{
			Filter: &dto.SprintFilter{
				ProjectId: project.Id,
				Page:      1,
				PerPage:   1000,
			},
		})
		if err != nil {
			log.Error("failed to list sprints", zap.Error(err))
			return nil, err
		}
		sprints = append(sprints, sprintsResp.Sprints.Items...)
	}
	log.Info("total sprints", zap.Int("count", len(sprints)))

	for _, user := range users {
		userStats := &dto.UserTaskStats{
			UserId:    user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Stats:     make(map[string]*dto.TaskStats),
		}

		for _, sprint := range sprints {
			tasksResp, err := l.projectClient.ListTasks(ctx, &pb.ListTasksRequest{
				Filter: &dto.TaskFilter{
					AssigneeId: user.Id,
					SprintId:   sprint.Id,
					Page:       1,
					PerPage:    10000,
				},
			})
			if err != nil {
				log.Error("failed to list tasks for user",
					zap.String("user_id", user.Id),
					zap.String("sprint_id", sprint.Id),
					zap.Error(err))
				continue
			}
			userStats.Stats[sprint.Id] = &dto.TaskStats{
				TotalTasks:      tasksResp.Tasks.TotalItems,
				DoneTasks:       0,
				InProgressTasks: 0,
				ToDoTasks:       0,
				TotalPoints:     0,
			}

			var totalPoints int32 = 0

			for _, task := range tasksResp.Tasks.Items {
				if task.Status == string(consts.TASK_STATUS_DONE) {
					userStats.Stats[sprint.Id].DoneTasks++
				} else if task.Status == string(consts.TASK_STATUS_IN_PROGRESS) {
					userStats.Stats[sprint.Id].InProgressTasks++
				} else if task.Status == string(consts.TASK_STATUS_CREATED) {
					userStats.Stats[sprint.Id].ToDoTasks++
				}

				taskPoints := utils.CalculateTaskPoints(task)
				totalPoints += taskPoints

				log.Debug("calculated points for task",
					zap.String("task_id", task.Id),
					zap.String("task_type", task.Type),
					zap.Int32("priority", task.Priority),
					zap.String("status", task.Status),
					zap.Int32("points", taskPoints))
			}

			userStats.Stats[sprint.Id].TotalPoints = totalPoints
		}

		func() {
			defer func() {
				result = append(result, userStats)
			}()
		}()

		tasksResp, err := l.projectClient.ListTasks(ctx, &pb.ListTasksRequest{
			Filter: &dto.TaskFilter{
				AssigneeId: user.Id,
				// CompanyId:  companyID,
				Page:    1,
				PerPage: 10000,
			},
		})
		if err != nil {
			log.Error("failed to list tasks for user",
				zap.String("user_id", user.Id),
				zap.Error(err))
			continue
		}
		userStats.Stats["overall"] = &dto.TaskStats{
			TotalTasks:      tasksResp.Tasks.TotalItems,
			DoneTasks:       0,
			InProgressTasks: 0,
			ToDoTasks:       0,
			TotalPoints:     0,
		}
		for _, task := range tasksResp.Tasks.Items {
			if task.Status == string(consts.TASK_STATUS_DONE) {
				userStats.Stats["overall"].DoneTasks++
			} else if task.Status == string(consts.TASK_STATUS_IN_PROGRESS) {
				userStats.Stats["overall"].InProgressTasks++
			} else if task.Status == string(consts.TASK_STATUS_CREATED) {
				userStats.Stats["overall"].ToDoTasks++
			}

			taskPoints := utils.CalculateTaskPoints(task)
			userStats.Stats["overall"].TotalPoints += taskPoints

			log.Debug("calculated points for task",
				zap.String("task_id", task.Id),
				zap.String("task_type", task.Type),
				zap.Int32("priority", task.Priority),
				zap.String("status", task.Status),
				zap.Int32("points", taskPoints))
		}
	}

	return result, nil
}
