package ctx

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"pms.pkg/utils"
)

type Session struct {
	SessionID string `json:"sessionID"`
}

func (s Session) ContextKey() ContextKey {
	return ContextKey("session")
}

func Test_Store(t *testing.T) {
	session := Session{
		SessionID: uuid.NewString(),
	}
	ctx := Embed(context.Background(), session)
	received, ok := Get(ctx, ContextKey("session"))
	if !ok {
		t.Fatal("failed to get context value")
	}
	t.Log(utils.JSON(received))
}
