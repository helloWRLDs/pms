package permission

import (
	"testing"
)

func Test_Permission(t *testing.T) {
	permSet := PermissionSet{}
	permSet = append(permSet, USER_READ)
	j, err := permSet.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(j))
	t.Log(permSet.StringArray())
}
