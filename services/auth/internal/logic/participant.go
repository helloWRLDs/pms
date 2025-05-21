package logic

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	participantdata "pms.auth/internal/data/participant"
	"pms.pkg/errs"
	"pms.pkg/transport/grpc/dto"
)

func (l *Logic) DeleteParticipant(ctx context.Context, userID, companyID string) error {
	log := l.log.With(
		zap.String("func", "DeleteParticipant"),
		zap.String("user_id", userID),
		zap.String("company_id", companyID),
	)
	log.Debug("DeleteParticipant called")

	if exist := l.Repo.Participant.Exists(ctx, userID, companyID); !exist {
		return errs.ErrNotFound{
			Object: "participant",
			Field:  "id",
			Value:  userID,
		}
	}

	participant, err := l.Repo.Participant.Get(ctx, userID, companyID)
	if err != nil {
		log.Errorw("failed to get participant", "err", err)
		return err
	}

	if err := l.Repo.Participant.Delete(ctx, participant.ID); err != nil {
		log.Errorw("failed to delete participant", "err", err)
		return err
	}

	return nil
}

func (l *Logic) AddParticipant(ctx context.Context, userID, companyID string) (*dto.Participant, error) {
	log := l.log.With(
		zap.String("func", "AddParticipant"),
		zap.String("user_id", userID),
		zap.String("company_id", companyID),
	)
	log.Debug("AddParticipant called")

	if exist := l.Repo.User.Exists(ctx, "id", userID); !exist {
		return nil, errs.ErrNotFound{
			Object: "user",
			Field:  "id",
			Value:  userID,
		}
	}
	if exist := l.Repo.Company.Exists(ctx, "id", companyID); !exist {
		return nil, errs.ErrNotFound{
			Object: "company",
			Field:  "id",
			Value:  companyID,
		}
	}
	if exists := l.Repo.Participant.Exists(ctx, userID, companyID); exists {
		return nil, errs.ErrConflict{
			Reason: "participant already exists",
		}
	}

	participant := participantdata.Participant{
		ID:        uuid.NewString(),
		UserID:    userID,
		CompanyID: companyID,
		Role:      "admin",
	}
	if err := l.Repo.Participant.Create(ctx, participant); err != nil {
		return nil, err
	}

	return participant.DTO(), nil
}
