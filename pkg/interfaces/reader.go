package interfaces

import (
	"pms.pkg/types/ctx"
	"pms.pkg/types/list"
)

type Reader[T any] interface {
	Get(ctx ctx.Context, field, value string) (T, error)
	List(ctx ctx.Context, filter list.Filters) (list.List[T], error)
	Count(ctx ctx.Context, filer list.Filters) int
	Exists(ctx ctx.Context, field, value string) bool
}
