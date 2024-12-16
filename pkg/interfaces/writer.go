package interfaces

import (
	"pms.pkg/types/ctx"
)

type Writer[T any] interface {
	Create(ctx ctx.Context, entity T) (string, error)
	Update(ctx ctx.Context, entity T) error
	Delete(ctx ctx.Context, id string) error
}
