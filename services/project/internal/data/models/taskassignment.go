package models

type TaskAssignment struct {
	UserID string `db:"user_id"`
	TaskID string `db:"task_id"`
}
