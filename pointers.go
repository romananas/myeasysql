package myeasysql

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

// _AssignPtrType takes a reflect.Value and returns a pointer to the underlying value
// of the specified type. It supports int, bool, string, time.Time, and any types.
// If the type is not recognized, it returns an error.
//
// Parameters:
//   - field: reflect.Value representing the value to be converted to a pointer.
//
// Returns:
//   - any: A pointer to the underlying value of the specified type.
//   - error: An error if the type is not recognized.
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

// _GetPointers takes an input of any type and returns a slice of pointers to the fields of the struct.
// It returns an error if the input is not a pointer or if the input is not a struct.
//
// Parameters:
//   - v: any type, expected to be a pointer to a struct.
//
// Returns:
//   - []any: a slice containing pointers to the fields of the struct.
//   - error: an error if the input is not a pointer or not a struct.
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
