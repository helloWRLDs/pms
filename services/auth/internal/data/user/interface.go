package userdata

import (
	"context"

	userdomain "pms.auth/internal/domain/user"
	"pms.pkg/type/list"
)

type Reader interface {
	GetByID(ctx context.Context, id string) (userdomain.User, error)
	List(ctx context.Context, filter list.Filters) (list.List[userdomain.User], error)
	Exists(ctx context.Context, email string) bool
	Count(ctx context.Context, filter list.Filters) int
	GetByEmail(ctx context.Context, email string) (userdomain.User, error)
}
