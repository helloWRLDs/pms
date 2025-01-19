package utils

import (
	"reflect"

	"pms.pkg/utils"
)

type EntityMapper struct {
	t      any
	ignore []string
}

func NewEnityMapper(entity any) *EntityMapper {
	return &EntityMapper{
		t:      entity,
		ignore: make([]string, 0),
	}
}

func (f *EntityMapper) Ignore(tags ...string) *EntityMapper {
	f.ignore = append(f.ignore, tags...)
	return f
}

func (f *EntityMapper) Columns() []string {
	if f.t == nil {
		return nil
	}
	var (
		cols = []string{}
		tt   = reflect.TypeOf(f.t)
		vt   = reflect.ValueOf(f.t)
	)
	for i := 0; i < vt.NumField(); i++ {
		tag := tt.Field(i).Tag.Get("db")
		if !utils.ContainsInArray(f.ignore, tag) {
			cols = append(cols, tag)
		}
	}
	return cols
}

func (f *EntityMapper) Values() []interface{} {
	if f.t == nil {
		return nil
	}
	var (
		values = []interface{}{}
		tt     = reflect.TypeOf(f.t)
		vt     = reflect.ValueOf(f.t)
	)
	for i := 0; i < vt.NumField(); i++ {
		if !utils.ContainsInArray(f.ignore, tt.Field(i).Tag.Get("db")) {
			values = append(values, vt.Field(i).Interface())
		}
	}
	return values
}
