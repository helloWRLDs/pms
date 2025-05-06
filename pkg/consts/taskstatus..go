package consts

type TaskStatus string

const (
	TASK_STATUS_CREATED     TaskStatus = "CREATED"
	TASK_STATUS_IN_PROGRESS TaskStatus = "IN_PROGRESS"
	TASK_STATUS_PENDING     TaskStatus = "PENDING"
	TASK_STATUS_DONE        TaskStatus = "DONE"
	TASK_STATUS_ARCHIVED    TaskStatus = "ARCHIVED"
)
