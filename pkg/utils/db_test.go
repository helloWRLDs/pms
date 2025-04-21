package utils

import (
	"testing"
	"time"
)

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func Test_GetColumns(t *testing.T) {
	user := User{ID: "1", Name: "Bob"}
	t.Log(GetColumns(user))
	t.Log(GetArguments(user)...)
}
