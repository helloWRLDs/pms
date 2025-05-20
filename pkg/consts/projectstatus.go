package consts

import (
	"database/sql/driver"
	"fmt"
)

type ProjectStatus string

const (
	PROJECT_STATUS_ACTIVE   ProjectStatus = "ACTIVE"
	PROJECT_STATUS_INACTIVE ProjectStatus = "INACTIVE"
	PROJECT_STATUS_CLOSED   ProjectStatus = "CLOSED"
)

func (s *ProjectStatus) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot scan %T into ProjectStatus", value)
	}
	*s = ProjectStatus(str)
	return nil
}

func (s ProjectStatus) Value() (driver.Value, error) {
	return string(s), nil
}
