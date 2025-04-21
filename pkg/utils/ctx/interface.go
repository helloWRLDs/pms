package ctxutils

type ContextKeyHolder interface {
	ContextKey() ContextKey
}

type ContextKey string
