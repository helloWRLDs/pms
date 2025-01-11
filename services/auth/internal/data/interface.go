package data

import (
	"context"

	"pms.auth/internal/domain"
	"pms.pkg/types/list"
)

type Reader interface {
	GetByID(ctx context.Context, id string) (domain.User, error)
	List(ctx context.Context, filter list.Filters) (list.List[domain.User], error)
	Exists(ctx context.Context, email string) bool
	Count(ctx context.Context, filter list.Filters) int
	GetByEmail(ctx context.Context, email string) (domain.User, error)
}
