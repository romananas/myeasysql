package myeasysql

import "reflect"

// _ReadNames takes a struct as input and returns a slice of strings containing
// the names of the fields in the struct.
//
// Parameters:
//   - v: an interface{} representing a pointer to a struct.
//
// Returns:
//   - []string: a slice of strings containing the names of the fields in the struct.
func readNames(v any) []string {
	var names []string
	tv := reflect.TypeOf(v).Elem()
	for i := range tv.NumField() {
		field := tv.Field(i)
		names = append(names, field.Name)
	}
	return names
}
