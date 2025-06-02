package ctxutils

import (
	"context"
	"testing"
)

type TestContextKey string

func (k TestContextKey) ContextKey() ContextKey {
	return ContextKey(k)
}

func TestEmbed(t *testing.T) {
	ctx := context.Background()
	key := TestContextKey("test_key")

	newCtx := Embed(ctx, key)
	if newCtx == ctx {
		t.Error("Embed() should return a new context")
	}

	if val, ok := newCtx.Value(key.ContextKey()).(TestContextKey); !ok || val != key {
		t.Errorf("Embed() value = %v, want %v", val, key)
	}
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	key := TestContextKey("test_key")

	if val, ok := Get(ctx, key.ContextKey()); ok {
		t.Errorf("Get() = %v, want nil", val)
	}

	ctx = context.WithValue(ctx, key.ContextKey(), key)
	if val, ok := Get(ctx, key.ContextKey()); !ok || val != key {
		t.Errorf("Get() = %v, want %v", val, key)
	}

	ctx = context.WithValue(ctx, key.ContextKey(), "wrong_type")
	if val, ok := Get(ctx, key.ContextKey()); ok {
		t.Errorf("Get() = %v, want nil", val)
	}
}
