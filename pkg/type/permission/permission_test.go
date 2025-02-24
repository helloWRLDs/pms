package permission

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func Test_Valuer(t *testing.T) {
	permSet := PermissionSet{}
	permSet = append(permSet, USER_READ, ORG_UPDATE)
	val, err := permSet.Value()
	assert.NoError(t, err)
	t.Logf("%#v", val)
}
