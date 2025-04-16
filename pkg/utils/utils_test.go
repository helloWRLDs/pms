package utils

import "testing"

func Test_MapToArray(t *testing.T) {
	m := make(map[int]string, 0)
	m[0] = "apple"
	m[1] = "banana"

	t.Logf("%#v", MapToArray(m))
}
