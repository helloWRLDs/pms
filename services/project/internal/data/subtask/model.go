package subtaskdata

import "pms.pkg/type/list"

type SubTask struct {
	ParentID string `db:"parent_id"`
	ChildID  string `db:"childID"`
}

type SubTaskFilter struct {
	list.Pagination
	list.Order
	list.Date
	ChildID   string `json:"child_id"`
	ParentID  string `json:"parent_id"`
	ProjectID string `json:"project_id"`
	SprintID  string `json:"sprint_id"`
}
