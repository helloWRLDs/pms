package consts

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Permission string

const (
	USER_READ_PERMISSION   Permission = "user:read"
	USER_WRITE_PERMISSION  Permission = "user:write"
	USER_DELETE_PERMISSION Permission = "user:delete"
	USER_INVITE_PERMISSION Permission = "user:invite"

	COMPANY_READ_PERMISSION   Permission = "company:read"
	COMPANY_WRITE_PERMISSION  Permission = "company:write"
	COMPANY_DELETE_PERMISSION Permission = "company:delete"
	COMPANY_INVITE_PERMISSION Permission = "company:invite"

	PROJECT_READ_PERMISSION   Permission = "project:read"
	PROJECT_WRITE_PERMISSION  Permission = "project:write"
	PROJECT_DELETE_PERMISSION Permission = "project:delete"
	PROJECT_INVITE_PERMISSION Permission = "project:invite"

	TASK_READ_PERMISSION   Permission = "task:read"
	TASK_WRITE_PERMISSION  Permission = "task:write"
	TASK_DELETE_PERMISSION Permission = "task:delete"
	TASK_INVITE_PERMISSION Permission = "task:invite"

	ROLE_READ_PERMISSION   Permission = "role:read"
	ROLE_WRITE_PERMISSION  Permission = "role:write"
	ROLE_DELETE_PERMISSION Permission = "role:delete"
	ROLE_INVITE_PERMISSION Permission = "role:invite"

	SPRINT_READ_PERMISSION   Permission = "sprint:read"
	SPRINT_WRITE_PERMISSION  Permission = "sprint:write"
	SPRINT_DELETE_PERMISSION Permission = "sprint:delete"
	SPRINT_INVITE_PERMISSION Permission = "sprint:invite"
)

type PermissionSet []Permission

func (p PermissionSet) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *PermissionSet) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return json.Unmarshal(src, p)
	case string:
		return json.Unmarshal([]byte(src), p)
	default:
		return fmt.Errorf("unsupported type for PermissionSet: %T", src)
	}
}
