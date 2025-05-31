package logic

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	documentdata "pms.analytics/internal/data/document"
	htmlmodule "pms.analytics/internal/modules/htmlgen"
	"pms.analytics/internal/modules/htmlgen/models"
	"pms.analytics/internal/utils"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (l *Logic) CreateReportTemplate(ctx context.Context, creation *dto.DocumentCreation) (string, error) {
	log := l.log.Named("CreateReportTemplate").With(
		zap.Any("sprint_id", creation.SprintId),
	)
	log.Info("CreateReportTemplate called")

	newDoc := documentdata.Document{
		ID:        uuid.NewString(),
		ProjectID: creation.ProjectId,
		CreatedAt: time.Now(),
		Title:     creation.Title,
		Body:      new([]byte),
	}

	if creation.SprintId != "" {
		sprint, err := l.getSprint(ctx, creation.SprintId)
		if err == nil {
			// Generate comprehensive sprint summary
			sprintSummary, err := l.generateSprintSummary(ctx, sprint)
			if err != nil {
				log.Errorw("failed to generate sprint summary", "err", err)
				return "", err
			}

			htmlData := htmlmodule.Template[models.SprintSummary]{
				Title:   creation.Title,
				Name:    models.SprintSummary{}.TemplateName(),
				Content: *sprintSummary,
			}

			html, err := htmlmodule.Render(htmlData)
			if err != nil {
				log.Errorw("failed to build report template", "err", err)
				return "", err
			}
			newDoc.Body = &html
		} else {
			log.Error("failed to get sprint", zap.Error(err))
		}
	}

	if err := l.Repo.Document.Create(ctx, newDoc); err != nil {
		log.Errorw("failed to create document", "err", err)
		return "", err
	}
	return newDoc.ID, nil
}

func (l *Logic) generateSprintSummary(ctx context.Context, sprint *dto.Sprint) (*models.SprintSummary, error) {
	log := l.log.Named("generateSprintSummary").With(
		zap.String("sprint_id", sprint.Id),
	)

	summary := &models.SprintSummary{
		Title:           sprint.Title,
		Description:     sprint.Description,
		StartDate:       sprint.StartDate.AsTime().Format("January 2, 2006"),
		EndDate:         sprint.EndDate.AsTime().Format("January 2, 2006"),
		TotalTasks:      len(sprint.Tasks),
		TasksByType:     make(map[string]int),
		TasksByPriority: make(map[string]int),
		UserPerformance: make([]models.UserPerformance, 0),
	}

	var (
		totalPoints   int32 = 0
		doneTasks           = 0
		userTaskMap         = make(map[string]*models.UserPerformance)
		taskInsights        = models.TaskInsights{}
		highestPoints int32 = 0
		typeCount           = make(map[string]int)
		priorityCount       = make(map[string]int)
		pointsSum     int32 = 0
	)

	// Analyze each task
	for _, task := range sprint.Tasks {
		// Count by status
		if task.Status == string(consts.TASK_STATUS_DONE) {
			doneTasks++
		}

		// Count by type and priority
		if task.Type != "" {
			summary.TasksByType[task.Type]++
			typeCount[task.Type]++
		}

		priorityName := l.getPriorityName(task.Priority)
		summary.TasksByPriority[priorityName]++
		priorityCount[priorityName]++

		// Calculate task points
		taskPoints := utils.CalculateTaskPoints(task)
		totalPoints += taskPoints
		pointsSum += taskPoints

		// Track highest value task
		if taskPoints > highestPoints {
			highestPoints = taskPoints
			taskInsights.HighestValueTask = task.Title
		}

		// Track user performance
		if task.AssigneeId != "" {
			if userPerf, exists := userTaskMap[task.AssigneeId]; exists {
				userPerf.TotalTasks++
				userPerf.TotalPoints += taskPoints
				if task.Status == string(consts.TASK_STATUS_DONE) {
					userPerf.DoneTasks++
				}
			} else {
				// Get user details
				userResp, err := l.authClient.GetUser(ctx, &pb.GetUserRequest{UserID: task.AssigneeId})
				if err == nil {
					user := userResp.User
					userTaskMap[task.AssigneeId] = &models.UserPerformance{
						UserID:      user.Id,
						FirstName:   user.FirstName,
						LastName:    user.LastName,
						FullName:    fmt.Sprintf("%s %s", user.FirstName, user.LastName),
						TotalTasks:  1,
						TotalPoints: taskPoints,
						DoneTasks: func() int32 {
							if task.Status == string(consts.TASK_STATUS_DONE) {
								return 1
							}
							return 0
						}(),
					}
				}
			}
		}
	}

	// Calculate metrics
	summary.DoneTasks = doneTasks
	summary.UndoneTasks = summary.TotalTasks - doneTasks
	summary.TotalPoints = totalPoints
	summary.SprintVelocity = totalPoints

	if summary.TotalTasks > 0 {
		summary.CompletionRate = float64(doneTasks) / float64(summary.TotalTasks) * 100
		taskInsights.AveragePointsPerTask = float64(pointsSum) / float64(summary.TotalTasks)
	}

	// Find most common type and priority
	taskInsights.MostCommonType = l.findMostCommon(typeCount)
	taskInsights.MostCommonPriority = l.findMostCommon(priorityCount)

	// Calculate user completion rates and convert to slice
	for _, userPerf := range userTaskMap {
		if userPerf.TotalTasks > 0 {
			userPerf.CompletionRate = float64(userPerf.DoneTasks) / float64(userPerf.TotalTasks) * 100
		}
		summary.UserPerformance = append(summary.UserPerformance, *userPerf)
	}

	// Sort users by points for top performers
	sort.Slice(summary.UserPerformance, func(i, j int) bool {
		return summary.UserPerformance[i].TotalPoints > summary.UserPerformance[j].TotalPoints
	})

	// Get top 3 performers
	topCount := 3
	if len(summary.UserPerformance) < topCount {
		topCount = len(summary.UserPerformance)
	}
	summary.TopPerformers = summary.UserPerformance[:topCount]

	summary.TaskInsights = taskInsights

	log.Info("sprint summary generated",
		zap.Int("total_tasks", summary.TotalTasks),
		zap.Int("done_tasks", summary.DoneTasks),
		zap.Float64("completion_rate", summary.CompletionRate),
		zap.Int32("total_points", summary.TotalPoints))

	return summary, nil
}

func (l *Logic) getPriorityName(priority int32) string {
	switch priority {
	case 1:
		return "Lowest"
	case 2:
		return "Low"
	case 3:
		return "Medium"
	case 4:
		return "High"
	case 5:
		return "Highest"
	default:
		return "Unknown"
	}
}

func (l *Logic) findMostCommon(counts map[string]int) string {
	var maxKey string
	var maxCount int
	for key, count := range counts {
		if count > maxCount {
			maxCount = count
			maxKey = key
		}
	}
	return maxKey
}
