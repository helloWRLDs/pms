package models

type SubTask struct {
	ParentID string `db:"parent_id"`
	ChildID  string `db:"childID"`
}
