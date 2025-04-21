package ctxutils

import "context"

func Embed(ctx context.Context, t ContextKeyHolder) context.Context {
	ctx = context.WithValue(ctx, t.ContextKey(), t)
	return ctx
}

func Get(ctx context.Context, key ContextKey) (t ContextKeyHolder, ok bool) {
	t, ok = ctx.Value(key).(ContextKeyHolder)
	return t, ok
}
