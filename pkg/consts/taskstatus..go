package consts

import (
	"database/sql/driver"
	"fmt"
)

type TaskStatus string

const (
	TASK_STATUS_CREATED     TaskStatus = "CREATED"
	TASK_STATUS_IN_PROGRESS TaskStatus = "IN_PROGRESS"
	TASK_STATUS_PENDING     TaskStatus = "PENDING"
	TASK_STATUS_DONE        TaskStatus = "DONE"
	TASK_STATUS_ARCHIVED    TaskStatus = "ARCHIVED"
)

func (s *TaskStatus) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot scan %T into TaskStatus", value)
	}
	*s = TaskStatus(str)
	return nil
}

func (s TaskStatus) Value() (driver.Value, error) {
	return string(s), nil
}
