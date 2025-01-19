package utils

import (
	"testing"
	"time"
)

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"fullname"`
	Age       int       `db:"age"`
	CreatedAt time.Time `db:"created_at"`
}

func Test_Cols(t *testing.T) {
	user := User{ID: 1, Name: "Bob", Age: 20, CreatedAt: time.Now()}
	fields := NewEnityMapper(user).Ignore([]string{"id", "created_at"}...)
	t.Log(fields.Columns())
	t.Log(fields.Values())
}
