package data

import (
	"pms.auth/internal/domain"
	"pms.pkg/types/ctx"
	"pms.pkg/types/list"
)

type Reader interface {
	Get(ctx ctx.Context, id string) (domain.User, error)
	List(ctx ctx.Context, filter list.Filters) (list.List[domain.User], error)
	Exists(ctx ctx.Context, email string) bool
	Count(ctx ctx.Context, filter list.Filters) int
	GetByEmail(ctx ctx.Context, email string) (domain.User, error)
}
