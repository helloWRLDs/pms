package logic

import (
	"context"

	"go.uber.org/zap"
	"pms.api-gateway/internal/models"
	"pms.pkg/transport/grpc/dto"
	pb "pms.pkg/transport/grpc/services"
)

func (l *Logic) GetCompanyContext(ctx context.Context, companyID string) (*models.CompanyContext, error) {
	log := l.log.With(
		zap.String("func", "GetCompanyContext"),
	)
	log.Info("GetCompanyContext called")

	cc, err := l.CompanyContext.Get(ctx, companyID)
	if err != nil {
		projectRes, err := l.projectClient.ListProjects(ctx, &pb.ListProjectsRequest{
			Filter: &dto.ProjectFilter{
				CompanyId: companyID,
				Page:      1,
				PerPage:   100,
			},
		})
		if err != nil {
			return nil, err
		}

		participantsRes, err := l.authClient.ListParticipants(ctx, &pb.ListParticipantsRequest{
			Filter: &dto.ParticipantFilter{
				CompanyId: companyID,
				Page:      1,
				PerPage:   100,
			},
		})
		if err != nil {
			return nil, err
		}

		projectIDs := make([]string, 0, len(projectRes.Projects.Items))
		for _, project := range projectRes.Projects.Items {
			projectIDs = append(projectIDs, project.Id)
		}
		sprintIDs := make([]string, 0)
		for _, project := range projectRes.Projects.Items {
			for _, sprint := range project.Sprints {
				sprintIDs = append(sprintIDs, sprint.Id)
			}
		}

		participants := make([]models.Participant, 0, len(participantsRes.Participants.Items))
		for _, p := range participantsRes.Participants.Items {
			userRes, err := l.authClient.GetUser(ctx, &pb.GetUserRequest{
				UserID: p.UserId,
			})
			if err != nil {
				log.Errorw("failed to get user details", "user_id", p.UserId, "err", err)
				continue
			}

			user := userRes.User
			participants = append(participants, models.Participant{
				UserID:    p.UserId,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
			})
		}

		cc = models.CompanyContext{
			CompanyID:    companyID,
			Projects:     projectIDs,
			Sprints:      sprintIDs,
			Participants: participants,
		}
	}
	return &cc, nil
}
