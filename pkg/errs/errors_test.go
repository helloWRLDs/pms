package errs

import (
	"testing"

	"pms.pkg/utils"
)

func Test_HttpErr(t *testing.T) {

	err := ErrNotFound{
		Object: "user",
		Field:  "id",
		Value:  "fdsfsdfdsfds",
	}
	t.Log(utils.JSON(WrapHttp(err)))
}
