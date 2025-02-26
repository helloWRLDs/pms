package userdata

import (
	"context"

	userentity "pms.auth/internal/entity/user"
	"pms.pkg/type/list"
)

type Reader interface {
	GetByID(ctx context.Context, id string) (userentity.User, error)
	List(ctx context.Context, filter list.Filters) (list.List[userentity.User], error)
	Exists(ctx context.Context, email string) bool
	Count(ctx context.Context, filter list.Filters) int
	GetByEmail(ctx context.Context, email string) (userentity.User, error)
}

type Writer interface {
	Create(ctx context.Context, newUser userentity.User) error
}
