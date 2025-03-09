package render

import (
	"reflect"
	"testing"
)

func Test_Render(t *testing.T) {
	greet := GreetContent{
		"Bob", "AITU",
	}
	t.Logf("%#v", mqTable(greet))
}

func mqTable(data interface{}) (table map[string]interface{}) {
	table = make(map[string]interface{}, 0)
	v := reflect.ValueOf(data)

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if !value.CanInterface() {
			continue
		}
		table[field.Name] = value.Interface()
	}
	return
}
