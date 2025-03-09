package userdata

import (
	"context"

	"pms.pkg/type/list"
)

type Reader[T any] interface {
	GetByID(ctx context.Context, id string) (T, error)
	List(ctx context.Context, filter list.Filters) (list.List[T], error)
	Exists(ctx context.Context, email string) bool
	Count(ctx context.Context, filter list.Filters) int
}

type Writer[T any] interface {
	Create(ctx context.Context, newUser T) error
}
