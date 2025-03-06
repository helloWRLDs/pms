package ctx

type ContextKeyHolder interface {
	ContextKey() ContextKey
}

type ContextKey string
