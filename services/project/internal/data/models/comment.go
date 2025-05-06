package models

import "time"

type TaskComment struct {
	ID        string    `db:"id"`
	Body      string    `db:"body"`
	TaskID    string    `db:"task_id"`
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}
