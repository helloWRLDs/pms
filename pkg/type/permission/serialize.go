package permission

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

func (p *PermissionSet) MarshalJSON() ([]byte, error) {
	perms := make([]string, len(*p))
	for i, perm := range *p {
		perms[i] = perm.String()
	}
	return json.Marshal(perms)
}

func (p *PermissionSet) UnmarshalJSON(data []byte) error {
	var perms []string
	if err := json.Unmarshal(data, &perms); err != nil {
		return err
	}
	var temp PermissionSet
	for _, perm := range perms {
		temp = append(temp, ParsePermission(perm))
	}
	*p = temp
	return nil
}

func (p *PermissionSet) Scan(value interface{}) error {
	if value == nil {
		*p = []Permission{}
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid data type")
	}
	var permString []string
	if err := json.Unmarshal([]byte(str), &permString); err != nil {
		return err
	}
	var permSet []Permission
	for _, perm := range permString {
		permSet = append(permSet, ParsePermission(perm))
	}
	*p = permSet
	return nil
}

func (p *PermissionSet) Value() (driver.Value, error) {
	return json.Marshal(p.StringArray())
}
