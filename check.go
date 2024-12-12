package myeasysql

import "reflect"

// _IsPtr checks if the given reflect.Value is a pointer.
// It returns true if the value is of kind reflect.Pointer, otherwise false.
func isPtr(rv reflect.Value) bool {
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
func isStruct(rv reflect.Value) bool {
	return rv.Elem().Kind() == reflect.Struct
}

func isSlice(rv reflect.Value) bool {
	// Vérifier que `dest` est un pointeur vers un slice
	return rv.Kind() == reflect.Ptr || rv.Elem().Kind() == reflect.Slice
}

// Obtenir le type de l'élément du slice
// elemType := destVal.Elem().Type().Elem()
// if elemType.Kind() != reflect.Struct {
// 	return fmt.Errorf("slice elements must be structs")
// }
