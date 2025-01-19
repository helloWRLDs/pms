package cache

import (
	"context"
	"testing"
	"time"

	cachemodels "pms.auth/internal/modules/cache/models"
	"pms.pkg/utils"
)

var (
	conf = Config{
		Host: "localhost:6379",
	}
)

func Test_Client(t *testing.T) {
	session := cachemodels.NewSession("session_1", "user_3", "someToken", time.Now().Add(24*time.Hour))

	client := New(conf, cachemodels.Session{})
	err := client.Set(context.Background(), session.ID, session, session.Exp.Unix())
	if err != nil {
		t.Fatal(err)
	}
	var retrieved cachemodels.Session
	retrieved, err = client.Get(context.Background(), "session_1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(utils.JSON(retrieved))
}
