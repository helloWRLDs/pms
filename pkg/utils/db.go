package utils

import "reflect"

func GetColumns(input any) []string {
	var columns []string

	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return columns
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")
		if tag == "-" || tag == "" {
			continue
		}

		if !v.Field(i).IsZero() {
			columns = append(columns, tag)
		}
	}
	return columns
}

func GetArguments(input any) []interface{} {
	var args []interface{}

	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return args
	}

	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i)
		field := t.Field(i)
		tag := field.Tag.Get("db")
		if tag == "-" || tag == "" {
			continue
		}

		if !val.IsZero() {
			args = append(args, val.Interface())
		}
	}
	return args
}
