package myeasysql

import "reflect"

// _IsPtr checks if the given reflect.Value is a pointer.
// It returns true if the value is of kind reflect.Pointer, otherwise false.
func _IsPtr(rv reflect.Value) bool {
	return rv.Kind() == reflect.Pointer
}

// _IsStruct checks if the given reflect.Value is a struct.
// It returns true if the underlying value is of kind struct, otherwise false.
//
// Parameters:
// - rv: reflect.Value to be checked.
//
// Returns:
// - bool: true if the underlying value is a struct, false otherwise.
func _IsStruct(rv reflect.Value) bool {
	return rv.Elem().Kind() == reflect.Struct
}
