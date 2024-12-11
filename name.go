package myeasysql

import "reflect"

// read all name in the pointer of structure v and return it as a []string
func _ReadNames(v any) []string {
	var names []string
	tv := reflect.TypeOf(v).Elem()
	for i := range tv.NumField() {
		field := tv.Field(i)
		names = append(names, field.Name)
	}
	return names
}
