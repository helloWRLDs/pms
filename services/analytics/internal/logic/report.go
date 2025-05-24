package logic

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	documentdata "pms.analytics/internal/data/document"
	htmlmodule "pms.analytics/internal/modules/htmlgen"
	"pms.analytics/internal/modules/htmlgen/models"
	"pms.pkg/consts"
	"pms.pkg/transport/grpc/dto"
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
	}

	if creation.SprintId != "" {
		sprint, err := l.getSprint(ctx, creation.SprintId)
		if err == nil {
			htmlData := htmlmodule.Template[models.SprintSummary]{
				Title: creation.Title,
				Name:  models.SprintSummary{}.TemplateName(),
			}
			htmlData.Content = models.SprintSummary{
				Title:       sprint.Title,
				Description: sprint.Description,
				StartDate:   sprint.StartDate.AsTime().Format(time.RFC3339),
				EndDate:     sprint.EndDate.AsTime().Format(time.RFC3339),
				TotalTasks:  len(sprint.Tasks),
				DoneTasks: func() (count int) {
					for _, task := range sprint.Tasks {
						if task.Status == string(consts.TASK_STATUS_DONE) {
							count++
						}
					}
					return
				}(),
				UndoneTasks: func() (count int) {
					for _, task := range sprint.Tasks {
						if task.Status != string(consts.TASK_STATUS_DONE) {
							count++
						}
					}
					return
				}(),
			}
			html, err := htmlmodule.Render(htmlData)
			if err != nil {
				log.Errorw("failed to build report template", "err", err)
				return "", err
			}
			newDoc.Body = html
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
