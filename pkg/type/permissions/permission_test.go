package permissions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Value(t *testing.T) {
	p := PermissionSet{
		ORG_READ_PERMISSION,
		ORG_WRITE_PERMISSION,
		USER_DELETE_PERMISSION,
		USER_READ_PERMISSION,
		USER_WRITE_PERMISSION,
	}
	val, err := p.Value()
	assert.NoError(t, err)
	t.Logf("%#v", val)
}
