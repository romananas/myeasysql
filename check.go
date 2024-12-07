package myeasysql

import "reflect"

func _IsPtr(rv reflect.Value) bool {
	return rv.Kind() == reflect.Pointer
}

func _IsStruct(rv reflect.Value) bool {
	return rv.Elem().Kind() == reflect.Struct
}
