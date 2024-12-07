package myeasysql

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

func _AssignPtrType(field reflect.Value) (any, error) {
	switch field.Type().String() {
	case "int":
		return (*int)(unsafe.Pointer(field.UnsafeAddr())), nil
	case "bool":
		return (*bool)(unsafe.Pointer(field.UnsafeAddr())), nil
	case "string":
		return (*string)(unsafe.Pointer(field.UnsafeAddr())), nil
	case "time.Time":
		return (*time.Time)(unsafe.Pointer(field.UnsafeAddr())), nil
	case "any":
		return (*any)(unsafe.Pointer(field.UnsafeAddr())), nil
	default:
		return nil, fmt.Errorf("type %s unreconized", field.Type().String())
	}
}

func _GetPointers(v any) ([]any, error) {
	var vPointers []any
	var rv = reflect.ValueOf(v)
	if !_IsPtr(rv) {
		return nil, fmt.Errorf("v is not a pointer")
	}
	if !_IsStruct(rv) {
		return nil, fmt.Errorf("v is not a struct")
	}
	rv = reflect.ValueOf(v).Elem()
	for i := range rv.NumField() {
		ptr, err := _AssignPtrType(rv.Field(i))
		if err != nil {
			return nil, err
		}
		vPointers = append(vPointers, ptr)
	}
	return vPointers, nil
}
