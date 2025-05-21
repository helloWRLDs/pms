package consts

import (
	"database/sql/driver"
	"fmt"
)

type TaskPriority int32

const (
	TASK_PRIORITY_HIGHEST TaskPriority = iota + 1
	TASK_PRIORITY_HIGH
	TASK_PRIORITY_MEDIUM
	TASK_PRIORITY_LOW
	TASK_PRIORITY_LOWEST
)

func (s *TaskPriority) Scan(value interface{}) error {
	str, ok := value.(int32)
	if !ok {
		return fmt.Errorf("cannot scan %T into TaskPriority", value)
	}
	*s = TaskPriority(str)
	return nil
}

func (s TaskPriority) Value() (driver.Value, error) {
	return int32(s), nil
}
