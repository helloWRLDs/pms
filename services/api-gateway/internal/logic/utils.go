package logic

import (
	"context"

	"pms.api-gateway/internal/models"
	"pms.pkg/errs"
	ctxutils "pms.pkg/utils/ctx"
)

func (l *Logic) GetSessionInfo(ctx context.Context) (models.Session, error) {
	session, ok := ctxutils.Get(ctx, models.Session{}.ContextKey())
	if !ok {
		return models.Session{}, errs.ErrUnauthorized{
			Reason: "session not found",
		}
	}
	return session.(models.Session), nil
}
