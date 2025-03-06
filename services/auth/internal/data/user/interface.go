package userdata

import (
	"context"

	"pms.auth/internal/entity"
	"pms.pkg/type/list"
)

type Reader interface {
	GetByID(ctx context.Context, id string) (entity.User, error)
	List(ctx context.Context, filter list.Filters) (list.List[entity.User], error)
	Exists(ctx context.Context, email string) bool
	Count(ctx context.Context, filter list.Filters) int
	GetByEmail(ctx context.Context, email string) (entity.User, error)
}

type Writer interface {
	Create(ctx context.Context, newUser entity.User) error
}
